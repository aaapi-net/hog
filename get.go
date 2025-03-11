package hog

import (
	"context"
	"net/http"
	"net/url"
)

type HGet struct {
	hog Hog
}

func (h *HGet) Headers(headers http.Header) *HGet {
	h.hog.headers = headers
	return h
}

func (h *HGet) SetHeader(key, value string) *HGet {
	h.hog.SetHeader(key, value)
	return h
}

func (h *HGet) Query(query url.Values) *HGet {
	h.hog.query = query
	return h
}

func (h *HGet) SetValue(key, value string) *HGet {
	if h.hog.query == nil {
		query := url.Values{}
		h.hog.query = query
	}
	h.hog.query.Set(key, value)
	return h
}

func (h *HGet) Response() (response *http.Response, err error) {
	if h.hog.context == nil {
		h.hog.context = context.Background()
	}

	req, err := http.NewRequestWithContext(h.hog.context, "GET", getFullUrl(h.hog.url, h.hog.query), nil)
	if err != nil {
		return nil, newError("NewRequest", "failed to create request", err)
	}

	fillHeaders(req.Header, h.hog.headers)

	h.hog.logger.Debug("Executing GET request:", req.URL)
	return h.hog.client.Do(req)
}

func (h *HGet) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(h)
}

func (h *HGet) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(h)
}

func (h *HGet) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(h, out)
}

func (h *HGet) AsMapResponse() (result map[string]interface{}, response *http.Response, err error) {
	response, err = h.ToStructResponse(&result)
	return
}

func (h *HGet) AsBytes() (result []byte, err error) {
	result, _, err = h.AsBytesResponse()
	return
}

func (h *HGet) AsString() (result string, err error) {
	result, _, err = h.AsStringResponse()
	return
}

func (h *HGet) ToStruct(out interface{}) (err error) {
	_, err = h.ToStructResponse(out)
	return
}

func (h *HGet) AsMap() (result map[string]interface{}, err error) {
	err = h.ToStruct(&result)
	return
}
