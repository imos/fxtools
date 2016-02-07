package hirose

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	// "strconv"
	// "github.com/imos/go/var_dump"
	"strconv"
	"strings"
	"time"
)

type Order struct {
	OrderId      string
	OrderMethod  string
	Currency     string
	PositionId   string
	IsSettlement bool
	Side         string
	IsStop       bool
	Price        float64
	Amount       int64
}

func getBaseQueryForOrder() map[string]string {
	t := time.Now()
	t.Add(time.Hour * 48)
	return map[string]string{
		"symbol_code":     "",
		"open_close_type": "",
		"f_y":             "2015",
		"f_m":             "01",
		"f_d":             "01",
		"f_h":             "00",
		"f_min":           "00",
		"t_y":             fmt.Sprintf("%d", t.Year()),
		"t_m":             fmt.Sprintf("%d", t.Month()),
		"t_d":             fmt.Sprintf("%d", t.Day()),
		"t_h":             "00",
		"t_min":           "00",
	}
}

func parseOrderListPage(s *goquery.Document) ([]Order, bool, error) {
	c := s.Find(".container").First()
	t := c.Find("div").First().Text()
	if t != ">注文情報(一覧)<" && t != ">注文情報(検索)<" {
		return nil, false, fmt.Errorf("cannot open \"注文情報(一覧)\", but %#v", t)
	}
	// タイトル行の削除
	c.Find("hr").First().Next().PrevAll().Remove()

	results := []Order{}
	c.Find("a").Each(
		func(_ int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if !ok || !strings.HasPrefix(href, "../otc/C003.html?") {
				return
			}
			u, err := url.Parse(href)
			if err != nil || u.RawQuery == "" {
				return
			}
			v, err := url.ParseQuery(u.RawQuery)
			results = append(results, Order{
				OrderId:     v.Get("order_id"),
				OrderMethod: v.Get("order_method"),
			})
		})

	return results, c.Find("a[accesskey=\"#\"]").Length() == 1, nil
}

func (c *HiroseClient) GetOrderList() ([]Order, error) {
	orders := []Order{}
	for index := 1; index <= 100; index++ {
		query := getBaseQueryForOrder()
		query["page_index"] = fmt.Sprintf("%d", index)
		doc, err := c.FetchWithQuery("GET", "/otc/C402.html", query)
		if err != nil {
			return nil, err
		}
		o, nextPage, err := parseOrderListPage(doc)
		if err != nil {
			return nil, err
		}
		if !nextPage || len(o) == 0 {
			break
		}
		orders = append(orders, o...)
	}
	return orders, nil
}

func parseOrderDetailsPage(s *goquery.Document) (*Order, error) {
	c := s.Find(".container").First()
	t := c.Find("div").First().Text()
	if t != ">注文情報(詳細)<" {
		return nil, fmt.Errorf("cannot open \"注文情報(詳細)\", but %#v", t)
	}

	o := &Order{}
	// タイトル行の削除
	c.Find("hr").First().Next().PrevAll().Remove()

	c.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		// 取消リンク
		if strings.HasPrefix(href, "../otc/CT01.html?") {
			u, err := url.Parse(href)
			if err != nil || u.RawQuery == "" {
				return
			}
			v, err := url.ParseQuery(u.RawQuery)
			if err != nil {
				return
			}
			o.OrderId = v.Get("order_id")
			o.OrderMethod = v.Get("order_method")
		}
		// 指定決済リンク
		if strings.HasPrefix(href, "../common/I124.html?") {
			u, err := url.Parse(href)
			if err != nil || u.RawQuery == "" {
				return
			}
			v, err := url.ParseQuery(u.RawQuery)
			if err != nil {
				return
			}
			o.PositionId = v.Get("position_id")
		}
	})

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

	o.Currency = orig["通貨ﾍﾟｱ"]
	lot, err := strconv.ParseInt(orig["注文Lot数"], 10, 64)
	if err != nil {
		return nil, err
	}
	o.Amount = lot * 1000
	if orig["売買"] == "売" {
		o.Side = "sell"
	} else if orig["売買"] == "買" {
		o.Side = "buy"
	} else {
		return nil, fmt.Errorf("invalid value for 売買: %#v", orig["売買"])
	}
	if orig["執行条件"] == "指値" {
		o.IsStop = false
	} else if orig["執行条件"] == "逆指" {
		o.IsStop = true
	} else {
		return nil, fmt.Errorf("invalid value for 執行条件: %#v", orig["執行条件"])
	}
	o.Price, err = strconv.ParseFloat(orig["指定ﾚｰﾄ"], 64)
	if err != nil {
		return nil, err
	}
	if orig["注文区分"] == "指定決済" {
		o.IsSettlement = true
	} else if orig["注文区分"] == "売買" {
		o.IsSettlement = false
	} else {
		return nil, fmt.Errorf("invalid value for 注文区分: %#v", orig["注文区分"])
	}

	return o, nil
}

func (c *HiroseClient) GetOrder(order Order) (*Order, error) {
	query := getBaseQueryForOrder()
	query["order_id"] = order.OrderId
	query["order_method"] = order.OrderMethod
	// NOTE: ヒロセ側が戻るページを生成するために必要
	query["page_index"] = "1"
	doc, err := c.FetchWithQuery("GET", "/otc/C403.html", query)
	if err != nil {
		return nil, err
	}
	return parseOrderDetailsPage(doc)
}

func (c *HiroseClient) GetOrders() ([]Order, error) {
	orders, err := c.GetOrderList()
	if err != nil {
		return nil, err
	}

	for i, order := range orders {
		o, err := c.GetOrder(order)
		if err != nil {
			return nil, err
		}
		orders[i] = *o
	}

	return orders, nil
}
