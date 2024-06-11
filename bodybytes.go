package hog

import (
	"bytes"
	"net/http"
)

type BodyBytes struct {
	method HMethod
}

func (bb *BodyBytes) fixHeaders(header *http.Header) {

}

func (bb *BodyBytes) getBuffer() (buf *bytes.Buffer, err error) {
	return bytes.NewBuffer(bb.method.getBody().([]byte)), nil
}

func (bb *BodyBytes) getMethod() HMethod {
	return bb.method
}

func (bb *BodyBytes) Response() (response *http.Response, err error) {
	return getResponse(bb)
}

func (bb *BodyBytes) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(bb)
}

func (bb *BodyBytes) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(bb)
}

func (bb *BodyBytes) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(bb, out)
}

func (bb *BodyBytes) AsBytes() (result []byte, err error) {
	result, _, err = bb.AsBytesResponse()
	return
}

func (bb *BodyBytes) AsString() (result string, err error) {
	result, _, err = bb.AsStringResponse()
	return
}

func (bb *BodyBytes) ToStruct(out interface{}) (err error) {
	_, err = bb.ToStructResponse(out)
	return
}

func (bb *BodyBytes) AsMap() (result map[string]interface{}, err error) {
	err = bb.ToStruct(&result)
	return
}
