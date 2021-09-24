package hog

import "net/http"

type PostXml struct {
	post HPost
}

func (h *PostXml) Response() (response *http.Response, err error){
	return
}

func (h *PostXml) AsBytesResponse() (result []byte, response *http.Response, err error){
	return AsBytesResponse(h)
}

func (h *PostXml) AsStringResponse() (result string, response *http.Response,  err error) {
	return AsStringResponse(h)
}

func (h *PostXml) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return ToStructResponse(h, out)
}


func (h *PostXml) AsBytes() (result []byte, err error){
	result, _,  err = h.AsBytesResponse()
	return
}

func (h PostXml) AsString() (result string, err error)  {
	result, _, err = h.AsStringResponse()
	return
}

func (h PostXml) ToStruct(out interface{}) (err error)  {
	_, err = h.ToStructResponse(out)
	return
}

func (h PostXml) AsMap() (result map[string]interface{}, err error)  {
	err = h.ToStruct(&result)
	return
}