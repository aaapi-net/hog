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

func (*BodyJson) fixHeaders(header *http.Header ) {
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

func (h *BodyJson) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(h)
}

func (h *BodyJson) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(h)
}

func (h *BodyJson) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(h, out)
}

func (h *BodyJson) AsBytes() (result []byte, err error) {
	result, _, err = h.AsBytesResponse()
	return
}

func (h BodyJson) AsString() (result string, err error) {
	result, _, err = h.AsStringResponse()
	return
}

func (h BodyJson) ToStruct(out interface{}) (err error) {
	_, err = h.ToStructResponse(out)
	return
}

func (h BodyJson) AsMap() (result map[string]interface{}, err error) {
	err = h.ToStruct(&result)
	return
}
