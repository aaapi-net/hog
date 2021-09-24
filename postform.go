package hog

import (
	"net/http"
)

type PostForm struct {
	post HPost
}

func (h *PostForm) Response() (response *http.Response, err error){
	return
}

func (h *PostForm) AsBytesResponse() (result []byte, response *http.Response, err error){
	return AsBytesResponse(h)
}

func (h *PostForm) AsStringResponse() (result string, response *http.Response,  err error) {
	return AsStringResponse(h)
}

func (h *PostForm) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return ToStructResponse(h, out)
}


func (h *PostForm) AsBytes() (result []byte, err error){
	result, _,  err = h.AsBytesResponse()
	return
}

func (h PostForm) AsString() (result string, err error)  {
	result, _, err = h.AsStringResponse()
	return
}

func (h PostForm) ToStruct(out interface{}) (err error)  {
	_, err = h.ToStructResponse(out)
	return
}

func (h PostForm) AsMap() (result map[string]interface{}, err error)  {
	err = h.ToStruct(&result)
	return
}