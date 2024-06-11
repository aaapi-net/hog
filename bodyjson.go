package hog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type BodyJson struct {
	method HMethod
}

func (bj *BodyJson) getBuffer() (buf *bytes.Buffer, err error) {
	data, err := json.Marshal(bj.method.getBody())
	if err != nil {
		return
	}

	return bytes.NewBuffer(data), nil
}

func (bj *BodyJson) getMethod() HMethod {
	return bj.method
}

func (*BodyJson) fixHeaders(header *http.Header) {
	if header.Get("Content-Type") == "" {
		header.Add("Content-Type", "application/json")
	}

	if header.Get("Accept") == "" {
		header.Add("Accept", "application/json; charset=utf-8")
	}
}

func (bj *BodyJson) Response() (response *http.Response, err error) {
	return getResponse(bj)
}

func (bj *BodyJson) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(bj)
}

func (bj *BodyJson) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(bj)
}

func (bj *BodyJson) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(bj, out)
}

func (bj *BodyJson) AsBytes() (result []byte, err error) {
	result, _, err = bj.AsBytesResponse()
	return
}

func (bj *BodyJson) AsString() (result string, err error) {
	result, _, err = bj.AsStringResponse()
	return
}

func (bj *BodyJson) ToStruct(out interface{}) (err error) {
	_, err = bj.ToStructResponse(out)
	return
}

func (bj *BodyJson) AsMap() (result map[string]interface{}, err error) {
	err = bj.ToStruct(&result)
	return
}
