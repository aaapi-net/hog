package hog

import (
	"net/http"
	"net/url"
)

type HPost struct {
	hog Hog
	body interface{}
}

func (h *HPost) getHog() *Hog {
	return &h.hog
}

func (h *HPost) getBody() interface{} {
	return h.body
}

func (h *HPost) getName() string {
	return http.MethodPost
}

func (h *HPost) Json(body interface{}) *BodyJson {
	h.body = body
	return &BodyJson{method: h}
}

func (h *HPost) Bytes(body []byte) *BodyBytes {
	h.body = body
	return &BodyBytes{method: h}
}

func (h *HPost) Form(body url.Values) *BodyForm {
	h.body = body
	return &BodyForm{method: h}
}

func (h *HPost) Xml(body interface{}) *BodyXml {
	h.body = body
	return &BodyXml{method: h}
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

