package hog

import (
	"bytes"
	"net/http"
)

// BodyBytes handles binary request bodies.
type BodyBytes struct {
	method HMethod
}

// fixHeaders ensures appropriate content type headers are set for binary requests.
func (bb *BodyBytes) fixHeaders(header *http.Header) {

}

// getBuffer wraps the request body as a buffer.
func (bb *BodyBytes) getBuffer() (buf *bytes.Buffer, err error) {
	return bytes.NewBuffer(bb.method.getBody().([]byte)), nil
}

// getMethod returns the HTTP method associated with this body.
func (bb *BodyBytes) getMethod() HMethod {
	return bb.method
}

// Response executes the request with binary body and returns the raw response.
func (bb *BodyBytes) Response() (response *http.Response, err error) {
	return getResponse(bb)
}

// AsBytesResponse executes the request and returns the response body as bytes along with the response.
func (bb *BodyBytes) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(bb)
}

// AsStringResponse executes the request and returns the response body as a string along with the response.
func (bb *BodyBytes) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(bb)
}

// ToStructResponse unmarshals the response body into the provided struct and returns the response.
// The out parameter should be a pointer to the target struct.
func (bb *BodyBytes) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(bb, out)
}

// AsBytes executes the request and returns the response body as bytes.
func (bb *BodyBytes) AsBytes() (result []byte, err error) {
	result, _, err = bb.AsBytesResponse()
	return
}

// AsString executes the request and returns the response body as a string.
func (bb *BodyBytes) AsString() (result string, err error) {
	result, _, err = bb.AsStringResponse()
	return
}

// ToStruct unmarshals the response body into the provided struct.
// The out parameter should be a pointer to the target struct.
func (bb *BodyBytes) ToStruct(out interface{}) (err error) {
	_, err = bb.ToStructResponse(out)
	return
}

// AsMap unmarshals the response body into a map[string]interface{}.
func (bb *BodyBytes) AsMap() (result map[string]interface{}, err error) {
	err = bb.ToStruct(&result)
	return
}
