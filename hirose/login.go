package hirose

import (
	"../utils"
	"github.com/PuerkitoBio/goquery"
	"log"
)

const HOST string = "https://lionfx-mob.hirose-fx.co.jp"

type HiroseClient struct {
	client *utils.HttpClient
}

func (c *HiroseClient) FetchWithQuery(method, urlStr string, query map[string]string) (*goquery.Document, error) {
	if c.client == nil {
		c.client = &utils.HttpClient{
			UseSjis: true,
		}
	}
	err := c.client.Load()
	if err != nil {
		log.Printf("Failed to load Cookie: %s", err)
	}
	doc, err := c.client.FetchDocument(method, HOST+urlStr, query)
	if err != nil {
		return nil, err
	}
	c.client.Save()
	return doc, err
}

func (c *HiroseClient) Fetch(method, urlStr string) (*goquery.Document, error) {
	return c.FetchWithQuery(method, urlStr, map[string]string{})
}

func (c *HiroseClient) SignIn(userId, password string) (bool, error) {
	// NOTE: HiroseFX site requires Cookie before signing in.
	doc, err := c.Fetch("GET", "/L001.html")
	if err != nil {
		return false, err
	}

	doc, err = c.FetchWithQuery("POST", "/L002.html", map[string]string{"user_id": userId, "password": password})
	if err != nil {
		return false, err
	}
	return doc.Find(".container>div").First().Text() == ">>ﾒｲﾝﾒﾆｭｰ<<", nil
}

func (c *HiroseClient) IsSignedIn() (bool, error) {
	doc, err := c.Fetch("GET", "/M001.html")
	if err != nil {
		return false, err
	}
	return doc.Find(".container>div").First().Text() == ">>ﾒｲﾝﾒﾆｭｰ<<", nil
}
