package hog

import (
	"bytes"
	"net/http"
	"net/url"
)

// BodyForm handles form urlencoded request bodies.
type BodyForm struct {
	method HMethod
}

// getBuffer encodes the request body as form data and returns it as a buffer.
func (bf *BodyForm) getBuffer() (buf *bytes.Buffer, err error) {
	body, ok := bf.method.getBody().(url.Values)
	if !ok {
		return nil, newError("getBuffer", "body is not url.Values", nil)
	}
	return bytes.NewBuffer([]byte(body.Encode())), nil
}

// getMethod returns the HTTP method associated with this body.
func (bf *BodyForm) getMethod() HMethod {
	return bf.method
}

// fixHeaders ensures appropriate content type headers are set for form requests.
func (*BodyForm) fixHeaders(header *http.Header) {
	if header.Get("Content-Type") == "" {
		header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	if header.Get("Accept") == "" {
		header.Add("Accept", "*/*")
	}
}

// Response executes the request with form body and returns the raw response.
func (bf *BodyForm) Response() (response *http.Response, err error) {
	return getResponse(bf)
}

// AsBytesResponse executes the request and returns the response body as bytes along with the response.
func (bf *BodyForm) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(bf)
}

// AsStringResponse executes the request and returns the response body as a string along with the response.
func (bf *BodyForm) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(bf)
}

// ToStructResponse unmarshals the response body into the provided struct and returns the response.
// The out parameter should be a pointer to the target struct.
func (bf *BodyForm) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(bf, out)
}

// AsBytes executes the request and returns the response body as bytes.
func (bf *BodyForm) AsBytes() (result []byte, err error) {
	result, _, err = bf.AsBytesResponse()
	return
}

// AsString executes the request and returns the response body as a string.
func (bf *BodyForm) AsString() (result string, err error) {
	result, _, err = bf.AsStringResponse()
	return
}

// ToStruct unmarshals the response body into the provided struct.
// The out parameter should be a pointer to the target struct.
func (bf *BodyForm) ToStruct(out interface{}) (err error) {
	_, err = bf.ToStructResponse(out)
	return
}

// AsMap unmarshals the response body into a map[string]interface{}.
func (bf *BodyForm) AsMap() (result map[string]interface{}, err error) {
	err = bf.ToStruct(&result)
	return
}
