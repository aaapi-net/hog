package hog

import (
	"context"
	"net/http"
	"net/url"
)

type HGet struct {
	hog Hog
}

// Headers sets all headers for this GET request.
func (h *HGet) Headers(headers http.Header) *HGet {
	h.hog.headers = headers
	return h
}

// SetHeader adds or replaces a single header for this GET request.
func (h *HGet) SetHeader(key, value string) *HGet {
	h.hog.SetHeader(key, value)
	return h
}

// Query sets query parameters for this GET request.
func (h *HGet) Query(query url.Values) *HGet {
	h.hog.query = query
	return h
}

// SetValue adds or updates a single query parameter.
func (h *HGet) SetValue(key, value string) *HGet {
	if h.hog.query == nil {
		query := url.Values{}
		h.hog.query = query
	}
	h.hog.query.Set(key, value)
	return h
}

// Response executes the GET request and returns the raw http.Response.
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

// AsBytesResponse executes the GET request and returns the response body as bytes along with the response.
func (h *HGet) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(h)
}

// AsStringResponse executes the GET request and returns the response body as a string along with the response.
func (h *HGet) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(h)
}

// ToStructResponse unmarshals the GET response body into the provided struct and returns the response.
// The out parameter should be a pointer to the target struct.
func (h *HGet) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(h, out)
}

// AsMapResponse unmarshals the GET response body into a map and returns it along with the response.
func (h *HGet) AsMapResponse() (result map[string]interface{}, response *http.Response, err error) {
	response, err = h.ToStructResponse(&result)
	return
}

// AsBytes executes the GET request and returns the response body as bytes.
func (h *HGet) AsBytes() (result []byte, err error) {
	result, _, err = h.AsBytesResponse()
	return
}

// AsString executes the GET request and returns the response body as a string.
func (h *HGet) AsString() (result string, err error) {
	result, _, err = h.AsStringResponse()
	return
}

// ToStruct unmarshals the GET response body into the provided struct.
// The out parameter should be a pointer to the target struct.
func (h *HGet) ToStruct(out interface{}) (err error) {
	_, err = h.ToStructResponse(out)
	return
}

// AsMap unmarshals the GET response body into a map[string]interface{}.
func (h *HGet) AsMap() (result map[string]interface{}, err error) {
	err = h.ToStruct(&result)
	return
}
