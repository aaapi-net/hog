package hog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// BodyJson handles JSON request bodies with appropriate Content-Type headers.
type BodyJson struct {
	method HMethod
}

// getBuffer serializes the request body as JSON and returns it as a buffer.
func (bj *BodyJson) getBuffer() (buf *bytes.Buffer, err error) {
	data, err := json.Marshal(bj.method.getBody())
	if err != nil {
		return
	}

	return bytes.NewBuffer(data), nil
}

// getMethod returns the HTTP method associated with this body.
func (bj *BodyJson) getMethod() HMethod {
	return bj.method
}

// fixHeaders ensures appropriate content type headers are set for JSON requests.
func (*BodyJson) fixHeaders(header *http.Header) {
	if header.Get("Content-Type") == "" {
		header.Add("Content-Type", "application/json")
	}

	if header.Get("Accept") == "" {
		header.Add("Accept", "application/json; charset=utf-8")
	}
}

// Response executes the request with JSON body and returns the raw response.
func (bj *BodyJson) Response() (response *http.Response, err error) {
	return getResponse(bj)
}

// AsBytesResponse executes the request and returns the response body as bytes along with the response.
func (bj *BodyJson) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(bj)
}

// AsStringResponse executes the request and returns the response body as a string along with the response.
func (bj *BodyJson) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(bj)
}

// ToStructResponse unmarshals the response body into the provided struct and returns the response.
// The out parameter should be a pointer to the target struct.
func (bj *BodyJson) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(bj, out)
}

// AsBytes executes the request and returns the response body as bytes.
func (bj *BodyJson) AsBytes() (result []byte, err error) {
	result, _, err = bj.AsBytesResponse()
	return
}

// AsString executes the request and returns the response body as a string.
func (bj *BodyJson) AsString() (result string, err error) {
	result, _, err = bj.AsStringResponse()
	return
}

// ToStruct unmarshals the response body into the provided struct.
// The out parameter should be a pointer to the target struct.
func (bj *BodyJson) ToStruct(out interface{}) (err error) {
	_, err = bj.ToStructResponse(out)
	return
}

// AsMap unmarshals the response body into a map[string]interface{}.
func (bj *BodyJson) AsMap() (result map[string]interface{}, err error) {
	err = bj.ToStruct(&result)
	return
}
