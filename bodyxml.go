package hog

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

type BodyXml struct {
	method HMethod
	prefix string
	indent string
}

func (bx *BodyXml) getBuffer() (buf *bytes.Buffer, err error) {
	data, err := xml.MarshalIndent(bx.method.getBody(), bx.prefix, bx.indent)
	if err != nil {
		return
	}

	return bytes.NewBuffer(data), nil
}

func (bx *BodyXml) getMethod() HMethod {
	return bx.method
}

func (*BodyXml) fixHeaders(header *http.Header) {
	if header.Get("Content-Type") == "" {
		header.Add("Content-Type", "application/xml")
	}

	if header.Get("Accept") == "" {
		header.Add("Accept", "application/xml; charset=utf-8")
	}
}

func (bx *BodyXml) MarshalSettings(prefix, indent string) *BodyXml {
	bx.prefix = prefix
	bx.indent = indent
	return bx
}

func (bx *BodyXml) Response() (response *http.Response, err error) {
	return getResponse(bx)
}

func (bx *BodyXml) AsBytesResponse() (result []byte, response *http.Response, err error) {
	return asBytesResponse(bx)
}

func (bx *BodyXml) AsStringResponse() (result string, response *http.Response, err error) {
	return asStringResponse(bx)
}

func (bx *BodyXml) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return toStructResponse(bx, out)
}

func (bx *BodyXml) AsBytes() (result []byte, err error) {
	result, _, err = bx.AsBytesResponse()
	return
}

func (bx *BodyXml) AsString() (result string, err error) {
	result, _, err = bx.AsStringResponse()
	return
}

func (bx *BodyXml) ToStruct(out interface{}) (err error) {
	_, err = bx.ToStructResponse(out)
	return
}

func (bx *BodyXml) AsMap() (result map[string]interface{}, err error) {
	err = bx.ToStruct(&result)
	return
}
