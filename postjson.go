package hog

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type PostJson struct {
	post HPost
}



func (h *PostJson) Response() (response *http.Response, err error){
	data, err := json.Marshal(h.post.body)
	if err != nil {
		return
	}

	if h.post.hog.context == nil {
		h.post.hog.context = context.Background()
	}

	req, err := http.NewRequestWithContext(h.post.hog.context, "POST", getFullUrl(h.post.hog.url, h.post.hog.query), bytes.NewBuffer(data))
	if err != nil {
		return
	}

	fillHeaders(&req.Header, h.post.hog.headers)

	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type",  "application/json")
	}

	return h.post.hog.client.Do(req)
}

func (h *PostJson) AsBytesResponse() (result []byte, response *http.Response, err error){
	return AsBytesResponse(h)
}

func (h *PostJson) AsStringResponse() (result string, response *http.Response,  err error) {
	return AsStringResponse(h)
}

func (h *PostJson) ToStructResponse(out interface{}) (response *http.Response, err error) {
	return ToStructResponse(h, out)
}


func (h *PostJson) AsBytes() (result []byte, err error){
	result, _,  err = h.AsBytesResponse()
	return
}

func (h PostJson) AsString() (result string, err error)  {
	result, _, err = h.AsStringResponse()
	return
}

func (h PostJson) ToStruct(out interface{}) (err error)  {
	_, err = h.ToStructResponse(out)
	return
}

func (h PostJson) AsMap() (result map[string]interface{}, err error)  {
	err = h.ToStruct(&result)
	return
}