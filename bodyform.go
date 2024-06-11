package hog

import (
	"bytes"
	"net/http"
	"net/url"
)

type BodyForm struct {
	method HMethod
}

func (bf *BodyForm) getBuffer() (buf *bytes.Buffer, err error) {
	body := bf.method.getBody().(url.Values).Encode()
	return bytes.NewBuffer([]byte(body)), nil
}

func (bf *BodyForm) getMethod() HMethod {
	return bf.method
}

func (*BodyForm) fixHeaders(header *http.Header) {
	if header.Get("Content-Type") == "" {
		header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	if header.Get("Accept") == "" {
		header.Add("Accept", "*/*")
	}
}

func (bf *BodyForm) Response() (response *http.Response, err error) {
	return getResponse(bf)
}

func (bf *BodyForm) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(bf)
}

func (bf *BodyForm) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(bf)
}

func (bf *BodyForm) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(bf, out)
}

func (bf *BodyForm) AsBytes() (result []byte, err error) {
	result, _, err = bf.AsBytesResponse()
	return
}

func (bf *BodyForm) AsString() (result string, err error) {
	result, _, err = bf.AsStringResponse()
	return
}

func (bf *BodyForm) ToStruct(out interface{}) (err error) {
	_, err = bf.ToStructResponse(out)
	return
}

func (bf *BodyForm) AsMap() (result map[string]interface{}, err error) {
	err = bf.ToStruct(&result)
	return
}
