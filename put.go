package hog

import (
	"net/http"
	"net/url"
)

type HPut struct {
	HPost
}

func (h *HPut) getName() string {
	return http.MethodPut
}

func (h *HPut) Json(body interface{}) *BodyJson {
	h.body = body
	return &BodyJson{method: h}
}

func (h *HPut) Bytes(body []byte) *BodyBytes {
	h.body = body
	return &BodyBytes{method: h}
}

func (h *HPut) Form(body url.Values) *BodyForm {
	h.body = body
	return &BodyForm{method: h}
}

func (h *HPut) Xml(body interface{}) *BodyXml {
	h.body = body
	return &BodyXml{method: h}
}

func (h *HPut) Headers(headers http.Header) *HPut {
	h.hog.headers = &headers
	return h
}

func (h *HPut) SetHeader(key, value string) *HPut {
	h.hog.headers.Set(key, value)
	return h
}

func (h *HPut) Query(query url.Values) *HPut {
	h.hog.query = &query
	return h
}

func (h *HPut) SetValue(key, value string) *HPut {
	h.hog.query.Set(key, value)
	return h
}
