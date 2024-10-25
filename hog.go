package hog

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Hog struct {
	client  http.Client
	headers *http.Header
	query   *url.Values
	context context.Context
	url     string
}

func GetF(format string, a ...any) *HGet {
	return Get(fmt.Sprintf(format, a...))
}

func PostF(format string, a ...any) *HPost {
	return Post(fmt.Sprintf(format, a...))
}

func PutF(format string, a ...any) *HPut {
	return Put(fmt.Sprintf(format, a...))
}

func Get(url string) *HGet {
	hog := New()
	return hog.Get(url)
}

func Post(url string) *HPost {
	hog := New()
	return hog.Post(url)
}

func Put(url string) *HPut {
	hog := New()
	return hog.Put(url)
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
	return &Hog{client: client, query: &url.Values{}}
}

func (h *Hog) Context(context context.Context) *Hog {
	h.context = context
	return h
}

func (h *Hog) GetF(format string, a ...any) *HGet {
	return h.Get(fmt.Sprintf(format, a...))
}

func (h *Hog) Get(url string) *HGet {
	h.url = url
	return &HGet{hog: *h}
}

func (h *Hog) Post(url string) *HPost {
	h.url = url
	return &HPost{hog: *h}
}

func (h *Hog) PostF(format string, a ...any) *HPost {
	return h.Post(fmt.Sprintf(format, a...))
}

func (h *Hog) Put(url string) *HPut {
	h.url = url
	return &HPut{HPost{hog: *h}}
}

func (h *Hog) PutF(format string, a ...any) *HPut {
	return h.Put(fmt.Sprintf(format, a...))
}

func (h *Hog) SetHeader(key, value string) *Hog {
	if h.headers == nil {
		h.headers = &http.Header{}
	}
	h.headers.Set(key, value)
	return h
}

func (h *Hog) Headers(headers http.Header) *Hog {
	h.headers = &headers
	return h
}

func getFullUrl(uri string, params *url.Values) string {
	if params == nil {
		return uri
	}
	return fmt.Sprint(uri, "?", params.Encode())
}

func fillHeaders(dest *http.Header, source *http.Header) {
	if dest != nil && source != nil {
		for k, varr := range *source {
			for _, v := range varr {
				dest.Add(k, v)
			}
		}
	}
}
