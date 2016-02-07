package utils

import (
	"./cookiejar"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpClient struct {
	initialized bool
	cookie      *cookiejar.Jar
	UseSjis     bool
}

func (h *HttpClient) getHttpClient() (*http.Client, error) {
	if !h.initialized {
		h.initialized = true
		cookie, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
		h.cookie = cookie
	}
	return &http.Client{
		Jar: h.cookie,
	}, nil
}

func (h *HttpClient) sendRequest(method, urlStr string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 8_4_1 like Mac OS X) "+
			"AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 "+
			"Mobile/12H321 Safari/600.1.4")
	client, err := h.getHttpClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (h *HttpClient) encode(buf []byte) ([]byte, error) {
	if !h.UseSjis {
		return buf, nil
	}
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader(buf), japanese.ShiftJIS.NewEncoder()))
}

func (h *HttpClient) decode(buf []byte) ([]byte, error) {
	if !h.UseSjis {
		return buf, nil
	}
	return ioutil.ReadAll(transform.NewReader(bytes.NewReader(buf), japanese.ShiftJIS.NewDecoder()))
}

func (h *HttpClient) Save(filename string) error {
	return h.cookie.Save(filename)
}

func (h *HttpClient) Load(filename string) error {
	return h.cookie.Load(filename)
}

func (h *HttpClient) Do(method, urlStr string, query map[string]string) ([]byte, error) {
	values := url.Values{}
	for k, v := range query {
		values.Add(k, v)
	}
	q, err := h.encode([]byte(values.Encode()))
	if err != nil {
		return nil, err
	}
	var resp []byte
	if method == "GET" {
		if len(query) == 0 {
			resp, err = h.sendRequest(method, urlStr, nil)
		} else {
			resp, err = h.sendRequest(method, urlStr+"?"+string(q), nil)
		}
	} else {
		resp, err = h.sendRequest(method, urlStr, bytes.NewReader(q))
	}
	if err != nil {
		return nil, err
	}
	return h.decode(resp)
}

func (h *HttpClient) FetchDocument(method, urlStr string, query map[string]string) (*goquery.Document, error) {
	buf, err := h.Do(method, urlStr, query)
	if err != nil {
		return nil, err
	}
	return goquery.NewDocumentFromReader(bytes.NewReader(buf))
}
