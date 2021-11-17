package hog

import (
	"net/http"
	"net/url"
)

type HPost struct {
	hog Hog
	body interface{}
}


func (h *HPost) Json(body interface{}) *PostJson {
	h.body = body
	return &PostJson{post: *h}
}

func (h *HPost) Bytes(body []byte) *PostBytes {
	h.body = body
	return &PostBytes{post: *h}
}

func (h *HPost) Form(body url.Values) *PostForm {
	h.body = body
	return &PostForm{post: *h}
}

func (h *HPost) Xml(body interface{}) *PostXml {
	h.body = body
	return &PostXml{post: *h}
}

func (h *HPost) Headers(headers http.Header) *HPost {
	h.hog.headers = &headers
	return h
}

func (h *HPost) SetHeader(key, value string) *HPost {
	h.hog.headers.Set(key, value)
	return h
}

func (h *HPost) Query(query url.Values) *HPost {
	h.hog.query = &query
	return h
}

func (h *HPost) SetValue(key, value string) *HPost {
	h.hog.query.Set(key, value)
	return h
}

