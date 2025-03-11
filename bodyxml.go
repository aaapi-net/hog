package hog

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

// BodyXml handles XML request bodies with appropriate Content-Type headers.
type BodyXml struct {
	method HMethod
	prefix string
	indent string
}

// getBuffer serializes the request body as XML and returns it as a buffer.
func (bx *BodyXml) getBuffer() (buf *bytes.Buffer, err error) {
	data, err := xml.MarshalIndent(bx.method.getBody(), bx.prefix, bx.indent)
	if err != nil {
		return
	}

	return bytes.NewBuffer(data), nil
}

// getMethod returns the HTTP method associated with this body.
func (bx *BodyXml) getMethod() HMethod {
	return bx.method
}

// fixHeaders ensures appropriate content type headers are set for XML requests.
func (*BodyXml) fixHeaders(header *http.Header) {
	if header.Get("Content-Type") == "" {
		header.Add("Content-Type", "application/xml")
	}

	if header.Get("Accept") == "" {
		header.Add("Accept", "application/xml; charset=utf-8")
	}
}

// MarshalSettings configures XML marshaling with the specified prefix and indent.
func (bx *BodyXml) MarshalSettings(prefix, indent string) *BodyXml {
	bx.prefix = prefix
	bx.indent = indent
	return bx
}

// Response executes the request with XML body and returns the raw response.
func (bx *BodyXml) Response() (response *http.Response, err error) {
	return getResponse(bx)
}

// AsBytesResponse executes the request and returns the response body as bytes along with the response.
func (bx *BodyXml) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(bx)
}

// AsStringResponse executes the request and returns the response body as a string along with the response.
func (bx *BodyXml) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(bx)
}

// ToStructResponse unmarshals the response body into the provided struct and returns the response.
// The out parameter should be a pointer to the target struct.
func (bx *BodyXml) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(bx, out)
}

// AsBytes executes the request and returns the response body as bytes.
func (bx *BodyXml) AsBytes() (result []byte, err error) {
	result, _, err = bx.AsBytesResponse()
	return
}

// AsString executes the request and returns the response body as a string.
func (bx *BodyXml) AsString() (result string, err error) {
	result, _, err = bx.AsStringResponse()
	return
}

// ToStruct unmarshals the response body into the provided struct.
// The out parameter should be a pointer to the target struct.
func (bx *BodyXml) ToStruct(out interface{}) (err error) {
	_, err = bx.ToStructResponse(out)
	return
}

// AsMap unmarshals the response body into a map[string]interface{}.
func (bx *BodyXml) AsMap() (result map[string]interface{}, err error) {
	err = bx.ToStruct(&result)
	return
}
