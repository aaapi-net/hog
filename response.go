package hog

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type response interface {
	Response() (response *http.Response, err error)
}

type responseBody interface {
	Response() (response *http.Response, err error)
	getBuffer() (buf *bytes.Buffer, err error)
	getMethod() HMethod
	fixHeaders(header *http.Header)
}

func getResponse(r responseBody) (response *http.Response, err error) {
	method := r.getMethod()
	hog := method.getHog()
	buf, err := r.getBuffer()

	if err != nil {
		return
	}

	if hog.context == nil {
		hog.context = context.Background()
	}

	req, err := http.NewRequestWithContext(hog.context, method.getName(), getFullUrl(hog.url, hog.query), buf)
	if err != nil {
		return
	}

	fillHeaders(&req.Header, hog.headers)
	r.fixHeaders(&req.Header)

	log.Println(req)
	return hog.client.Do(req)
}

func asBytesResponse(r response) (result []byte, response *http.Response, err error) {
	response, err = r.Response()
	defer response.Body.Close()

	if err == nil {
		result, err = ioutil.ReadAll(response.Body)
		return
	}

	return
}

func asStringResponse(r response) (result string, response *http.Response, err error) {
	data, response, err := asBytesResponse(r)

	if err == nil {
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
