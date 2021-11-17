package hog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response interface {
	Response() (response *http.Response, err error)
}

func AsBytesResponse(r Response) (result []byte, response *http.Response, err error){
	response, err = r.Response()
	defer response.Body.Close()

	if err == nil {
		result, err = ioutil.ReadAll(response.Body)
		return
	}

	return
}

func AsStringResponse(r Response) (result string, response *http.Response,  err error) {
	data, response, err := AsBytesResponse(r)

	if err == nil{
		result = string(data)
	}

	return
}

func ToStructResponse(r Response, out interface{}) (response *http.Response, err error) {
	data, response, err := AsBytesResponse(r)

	if err == nil {
		err = json.Unmarshal(data, out)
	}

	return
}
