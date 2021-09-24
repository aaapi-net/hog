package hog

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)


type Hog struct {
	client http.Client
	headers *http.Header
	query *url.Values
	context context.Context
	url string
}

func Get(url string) *HGet {
	hog := New()
	return hog.Get(url)
}

func Post(url string) *HPost {
	hog := New()
	return hog.Post(url)
}


func New() *Hog {
	return NewConfig(true, 30)
}

func NewConfig(secure bool, timeout int) *Hog {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !secure},
	}
	client := http.Client{Transport: tr, Timeout: time.Duration(timeout) * time.Second}
	return NewClient(client)
}

func NewClient(client http.Client) *Hog {
	return &Hog{client: client}
}


func (h *Hog) Context(context context.Context) *Hog {
	h.context = context
	return h
}


func (h *Hog) Post(url string) *HPost {
	h.url = url
	return  &HPost{hog: *h}
}

func getFullUrl(uri string, params *url.Values) string {
	url, _ := url.ParseRequestURI(uri)

	if params != nil {
		for k, varr := range *params {
			for _, v := range varr {
				url.Query().Add(k, v)
			}
		}

	}

	return url.String()
}

func fillHeaders(dest *http.Header, source *http.Header,)  {
	if dest != nil && source != nil{
		for k, varr := range *source {
			for _, v := range varr {
				dest.Add(k, v)
			}
		}
	}
}