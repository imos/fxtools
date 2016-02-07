package hirose

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Position struct {
	PositionId      string
	Currency        string
	TransactionTime string
	Side            string
	TransactionRate float64
	Amount          int64
}

func getBaseQuery() map[string]string {
	t := time.Now()
	t.Add(time.Hour * 48)
	return map[string]string{
		"symbol_code": "",
		"f_y":         "2015",
		"f_m":         "01",
		"f_d":         "01",
		"f_h":         "00",
		"f_min":       "00",
		"t_y":         fmt.Sprintf("%d", t.Year()),
		"t_m":         fmt.Sprintf("%d", t.Month()),
		"t_d":         fmt.Sprintf("%d", t.Day()),
		"t_h":         "00",
		"t_min":       "00",
	}
}

func parsePositionListPage(s *goquery.Document) ([]Position, bool, error) {
	c := s.Find(".container").First()
	t := c.Find("div").First().Text()
	if t != ">ﾎﾟｼﾞｼｮﾝ情報(一覧)<" && t != ">ﾎﾟｼﾞｼｮﾝ情報(検索)<" {
		return nil, false, fmt.Errorf("cannot open \"ﾎﾟｼﾞｼｮﾝ情報(一覧)\", but %#v", t)
	}

	results := []Position{}
	c.Find("a").Each(
		func(_ int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if !ok || !strings.HasPrefix(href, "C304.html") {
				return
			}
			u, err := url.Parse(href)
			if err != nil || u.RawQuery == "" {
				return
			}
			v, err := url.ParseQuery(u.RawQuery)
			results = append(results, Position{
				PositionId: v.Get("position_id"),
			})
		})

	return results, c.Find("a[accesskey=\"#\"]").Length() == 1, nil
}

func (c *HiroseClient) GetPositionList() ([]Position, error) {
	positions := []Position{}
	for index := 1; index <= 100; index++ {
		query := getBaseQuery()
		query["page_index"] = fmt.Sprintf("%d", index)
		doc, err := c.FetchWithQuery("GET", "/otc/C302.html", query)
		if err != nil {
			return nil, err
		}
		p, nextPage, err := parsePositionListPage(doc)
		if err != nil {
			return nil, err
		}
		if !nextPage || len(p) == 0 {
			break
		}
		positions = append(positions, p...)
	}
	return positions, nil
}

func parsePositionDetailsPage(s *goquery.Document) (*Position, error) {
	c := s.Find(".container").First()
	t := c.Find("div").First().Text()
	if t != ">ﾎﾟｼﾞｼｮﾝ情報(詳細)<" {
		return nil, fmt.Errorf("cannot open \"ﾎﾟｼﾞｼｮﾝ情報(詳細)\", but %#v", t)
	}

	position := &Position{}

	// タイトル行の削除
	c.Find("hr").First().Next().PrevAll().Remove()

	// 通貨名の設定
	position.Currency = c.Find("div").First().Text()
	// 通貨名行の削除
	c.Find("hr").First().Next().PrevAll().Remove()

	// メニューの削除
	c.Find("hr").First().Prev().NextAll().Remove()

	results := []string{}
	c.Find("div").Each(func(_ int, s *goquery.Selection) {
		results = append(results, s.Text())
	})

	orig := map[string]string{}
	for i := 0; i+1 < len(results); i += 2 {
		k := strings.TrimSpace(strings.Replace(results[i], "：", "", -1))
		v := strings.TrimSpace(results[i+1])
		orig[k] = v
	}

	position.PositionId = orig["ﾎﾟｼﾞｼｮﾝ番号"]
	position.TransactionTime = orig["約定日時"]
	if orig["売買"] == "売" {
		position.Side = "sell"
	} else if orig["売買"] == "買" {
		position.Side = "buy"
	}
	rate, err := strconv.ParseFloat(orig["約定価格"], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid transaction rate: %v", err)
	}
	position.TransactionRate = rate
	lot, err := strconv.ParseInt(orig["残Lot数"], 10, 64)
	position.Amount = lot * 1000
	return position, nil
}

func (c *HiroseClient) GetPosition(position Position) (*Position, error) {
	query := getBaseQuery()
	query["position_id"] = position.PositionId
	doc, err := c.FetchWithQuery("GET", "/otc/C304.html", query)
	if err != nil {
		return nil, err
	}
	return parsePositionDetailsPage(doc)
}

func (c *HiroseClient) GetPositions() ([]Position, error) {
	positions, err := c.GetPositionList()
	if err != nil {
		return nil, err
	}

	for i, position := range positions {
		p, err := c.GetPosition(position)
		if err != nil {
			return nil, err
		}
		positions[i] = *p
	}

	return positions, nil
}
