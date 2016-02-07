package hirose

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type Status struct {
	ActualDeposit    *int64 `json:"actual_deposit"`
	NecessaryDeposit *int64 `json:"necessary_deposit"`
}

func parsePrice(str string) *int64 {
	val, err := strconv.ParseInt(strings.Replace(str, ",", "", -1), 10, 64)
	if err != nil {
		return nil
	}
	return &val
}

func parseStatusPage(s *goquery.Document) (*Status, error) {
	c := s.Find(".container").First()
	t := c.Find("div").First().Text()
	if t != ">証拠金状況照会<" {
		return nil, fmt.Errorf("cannot open \"証拠金状況照会\", but %#v", t)
	}

	results := []string{}
	c.Find("hr").First().NextUntil("hr").Each(
		func(_ int, s *goquery.Selection) {
			results = append(results, s.Text())
		})

	orig := map[string]string{}
	for i := 0; i+1 < len(results); i += 2 {
		k := strings.TrimSpace(strings.Replace(results[i], "：", "", -1))
		v := strings.TrimSpace(results[i+1])
		orig[k] = v
	}

	return &Status{
		ActualDeposit:    parsePrice(orig["有効証拠金額"]),
		NecessaryDeposit: parsePrice(orig["必要証拠金額"]),
	}, nil
}

func (c *HiroseClient) GetStatus() (*Status, error) {
	doc, err := c.Fetch("GET", "/common/I103.html")
	if err != nil {
		return nil, err
	}
	return parseStatusPage(doc)
}
