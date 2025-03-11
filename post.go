package hog

import (
	"net/http"
	"net/url"
)

// HPost handles HTTP POST requests with chainable options.
type HPost struct {
	hog  Hog
	body interface{}
}

// getHog returns the Hog instance associated with this request.
func (h *HPost) getHog() *Hog {
	return &h.hog
}

// getBody returns the request body.
func (h *HPost) getBody() interface{} {
	return h.body
}

// getName returns the HTTP method name.
func (h *HPost) getName() string {
	return http.MethodPost
}

// Json sets a JSON request body.
func (h *HPost) Json(body interface{}) *BodyJson {
	h.body = body
	return &BodyJson{method: h}
}

// Bytes sets a binary request body.
func (h *HPost) Bytes(body []byte) *BodyBytes {
	h.body = body
	return &BodyBytes{method: h}
}

// Form sets a form urlencoded request body.
func (h *HPost) Form(body url.Values) *BodyForm {
	h.body = body
	return &BodyForm{method: h}
}

// Xml sets an XML request body.
func (h *HPost) Xml(body interface{}) *BodyXml {
	h.body = body
	return &BodyXml{method: h}
}

// Headers sets all headers for this POST request.
func (h *HPost) Headers(headers http.Header) *HPost {
	h.hog.headers = headers
	return h
}

// SetHeader adds or replaces a single header for this POST request.
func (h *HPost) SetHeader(key, value string) *HPost {
	h.hog.SetHeader(key, value)
	return h
}

// Query sets query parameters for this POST request.
func (h *HPost) Query(query url.Values) *HPost {
	h.hog.query = query
	return h
}

// SetValue adds or updates a single query parameter.
func (h *HPost) SetValue(key, value string) *HPost {
	if h.hog.query == nil {
		h.hog.query = url.Values{}
	}

	h.hog.query.Set(key, value)
	return h
}
