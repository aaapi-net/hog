package hog

import (
	"bytes"
	"context"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"time"
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

	if hog.logger == nil {
		hog.logger = &defaultLogger{level: LogLevelError}
	}

	buf, err := r.getBuffer()
	if err != nil {
		return nil, newError("getBuffer", "failed to prepare request body", err)
	}

	if hog.context == nil {
		hog.context = context.Background()
	}

	req, err := http.NewRequestWithContext(hog.context, method.getName(), getFullUrl(hog.url, hog.query), buf)
	if err != nil {
		return nil, newError("NewRequest", "failed to create request", err)
	}

	fillHeaders(&req.Header, hog.headers)
	r.fixHeaders(&req.Header)

	hog.logger.Debug("Sending request:", req.Method, req.URL)

	var lastErr error
	for retry := 0; retry <= hog.retryCount; retry++ {
		if retry > 0 {
			hog.logger.Info("Retrying request, attempt:", retry)
			select {
			case <-hog.context.Done():
				return nil, newError("Retry", "context cancelled during retry", hog.context.Err())
			case <-time.After(time.Duration(retry) * time.Second):
				// Exponential delay between retries
			}
		}

		response, err = hog.client.Do(req)
		if err == nil {
			hog.logger.Debug("Request successful:", response.Status)
			return response, nil
		}

		lastErr = err
		hog.logger.Error("Request failed:", err)
	}

	return nil, newError("Request", "all retry attempts failed", lastErr)
}

func asBytesResponse(r response) (result []byte, response *http.Response, err error) {
	response, err = r.Response()

	if err == nil {
		defer response.Body.Close()
		result, err = io.ReadAll(response.Body)
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
		err = sonic.ConfigFastest.Unmarshal(data, out)
	}

	return
}
