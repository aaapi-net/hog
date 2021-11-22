package hog

import (
	"bytes"
	"context"
	"log"
	"net/http"
)

type PostBytes struct {
	post HPost
}

func (h *PostBytes) Response() (response *http.Response, err error) {
	if h.post.hog.context == nil {
		h.post.hog.context = context.Background()
	}

	req, err := http.NewRequestWithContext(h.post.hog.context, "POST", getFullUrl(h.post.hog.url, h.post.hog.query), bytes.NewBuffer(h.post.body.([]byte)))
	if err != nil {
		return
	}

	fillHeaders(&req.Header, h.post.hog.headers)

	if req.Header.Get("Accept") == "" {
		req.Header.Add("Accept", "application/json; charset=utf-8")
	}

	log.Println(req)
	return h.post.hog.client.Do(req)
}

func (h *PostBytes) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(h)
}

func (h *PostBytes) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(h)
}

func (h *PostBytes) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(h, out)
}

func (h *PostBytes) AsBytes() (result []byte, err error) {
	result, _, err = h.AsBytesResponse()
	return
}

func (h PostBytes) AsString() (result string, err error) {
	result, _, err = h.AsStringResponse()
	return
}

func (h PostBytes) ToStruct(out interface{}) (err error) {
	_, err = h.ToStructResponse(out)
	return
}

func (h PostBytes) AsMap() (result map[string]interface{}, err error) {
	err = h.ToStruct(&result)
	return
}
