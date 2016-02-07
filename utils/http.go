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
	"os"
	"strings"
)

type HttpClient struct {
	initialized bool
	cookie      *cookiejar.Jar
	UseSjis     bool
}

func (h *HttpClient) getHttpClient() (*http.Client, error) {
	if !h.initialized {
		h.initialized = true
		if h.cookie == nil {
			cookie, err := cookiejar.New(nil)
			if err != nil {
				return nil, err
			}
			h.cookie = cookie
		}
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
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
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

func (h *HttpClient) Save() error {
	return h.cookie.Save(os.Getenv("HOME") + "/.fxcookie")
}

func (h *HttpClient) Load() error {
	if h.cookie == nil {
		cookie, err := cookiejar.New(nil)
		if err != nil {
			return err
		}
		h.cookie = cookie
	}
	return h.cookie.Load(os.Getenv("HOME") + "/.fxcookie")
}

func (h *HttpClient) Do(method, urlStr string, query map[string]string) ([]byte, error) {
	values := url.Values{}
	for k, v := range query {
		buf, _ := h.encode([]byte(v))
		values.Add(k, string(buf))
	}
	q := values.Encode()
	var resp []byte
	var err error
	if len(query) == 0 {
		resp, err = h.sendRequest(method, urlStr, nil)
	} else if method == "GET" {
		resp, err = h.sendRequest(method, urlStr+"?"+q, nil)
	} else {
		resp, err = h.sendRequest(method, urlStr, strings.NewReader(q))
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
