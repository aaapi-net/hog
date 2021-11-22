package hog

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type response interface {
	Response() (response *http.Response, err error)
}

func asBytesResponse(r response) (result []byte, response *http.Response, err error){
	response, err = r.Response()
	defer response.Body.Close()

	if err == nil {
		result, err = ioutil.ReadAll(response.Body)
		return
	}

	return
}

func asStringResponse(r response) (result string, response *http.Response,  err error) {
	data, response, err := asBytesResponse(r)

	if err == nil{
		result = string(data)
	}

	return
}

func toStructResponse(r response, out interface{}) (response *http.Response, err error) {
	data, response, err := asBytesResponse(r)

	if err == nil {
		err = json.Unmarshal(data, out)
		log.Println(string(data))
	}

	return
}
