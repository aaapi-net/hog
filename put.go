package hog

import (
	"net/http"
	"net/url"
)

// HPut handles HTTP PUT requests with chainable options.
type HPut struct {
	HPost
}

// getName returns the HTTP method name.
func (h *HPut) getName() string {
	return http.MethodPut
}

// Json sets a JSON request body.
func (h *HPut) Json(body interface{}) *BodyJson {
	h.body = body
	return &BodyJson{method: h}
}

// Bytes sets a binary request body.
func (h *HPut) Bytes(body []byte) *BodyBytes {
	h.body = body
	return &BodyBytes{method: h}
}

// Form sets a form urlencoded request body.
func (h *HPut) Form(body url.Values) *BodyForm {
	h.body = body
	return &BodyForm{method: h}
}

// Xml sets an XML request body.
func (h *HPut) Xml(body interface{}) *BodyXml {
	h.body = body
	return &BodyXml{method: h}
}

// Headers sets all headers for this PUT request.
func (h *HPut) Headers(headers http.Header) *HPut {
	h.hog.headers = headers
	return h
}

// SetHeader adds or replaces a single header for this PUT request.
func (h *HPut) SetHeader(key, value string) *HPut {
	h.hog.SetHeader(key, value)
	return h
}

// Query sets query parameters for this PUT request.
func (h *HPut) Query(query url.Values) *HPut {
	h.hog.query = query
	return h
}

// SetValue adds or updates a single query parameter.
func (h *HPut) SetValue(key, value string) *HPut {
	if h.hog.query == nil {
		h.hog.query = url.Values{}
	}

	h.hog.query.Set(key, value)
	return h
}
