package hog

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Hog is a flexible HTTP client that provides a chainable API for HTTP requests.
type Hog struct {
	client     http.Client
	headers    http.Header
	query      url.Values
	context    context.Context
	url        string
	logger     Logger
	retryCount int
}

func GetF(format string, a ...any) *HGet {
	return Get(fmt.Sprintf(format, a...))
}

func PostF(url string, a ...any) *HPost {
	return Post(fmt.Sprintf(url, a...))
}

func PutF(url string, a ...any) *HPut {
	return Put(fmt.Sprintf(url, a...))
}

func Get(url string) *HGet {
	hog := New()
	return hog.Get(url)
}

func Post(url string) *HPost {
	hog := New()
	return hog.Post(url)
}

func Put(url string) *HPut {
	hog := New()
	return hog.Put(url)
}

// New creates a new Hog instance with default secure configuration.
func New() *Hog {
	h := NewConfig(true, 30)
	h.logger = newDefaultLogger(LogLevelError)
	return h
}

// NewConfig creates a new Hog instance with specified TLS security and timeout.
// The secure parameter determines if TLS verification is enabled, and timeout is in seconds.
func NewConfig(secure bool, timeout int) *Hog {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !secure},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   time.Duration(timeout) * time.Second,
	}

	return NewClient(client)
}

// NewClient creates a new Hog instance using an existing http.Client.
func NewClient(client http.Client) *Hog {
	return &Hog{client: client, query: url.Values{}}
}

// Context sets the request context.
func (h *Hog) Context(context context.Context) *Hog {
	h.context = context
	return h
}

func (h *Hog) GetF(format string, a ...any) *HGet {
	return h.Get(fmt.Sprintf(format, a...))
}

// Get initiates a GET request on this Hog instance.
func (h *Hog) Get(url string) *HGet {
	h.url = url
	return &HGet{hog: *h}
}

// Post initiates a POST request on this Hog instance.
func (h *Hog) Post(url string) *HPost {
	h.url = url
	return &HPost{hog: *h}
}

// PostF initiates a POST request with a formatted URL string.
func (h *Hog) PostF(format string, a ...any) *HPost {
	return h.Post(fmt.Sprintf(format, a...))
}

// Put initiates a PUT request on this Hog instance.
func (h *Hog) Put(url string) *HPut {
	h.url = url
	return &HPut{HPost{hog: *h}}
}

// PutF initiates a PUT request with a formatted URL string.
func (h *Hog) PutF(format string, a ...any) *HPut {
	return h.Put(fmt.Sprintf(format, a...))
}

// SetHeader adds or replaces a single header for the request.
func (h *Hog) SetHeader(key, value string) *Hog {
	h.headers.Set(key, value)
	return h
}

// Headers sets all headers for the request.
func (h *Hog) Headers(headers http.Header) *Hog {
	h.headers = headers
	return h
}

// Logger sets a custom logger for the request.
func (h *Hog) Logger(logger Logger) *Hog {
	h.logger = logger
	return h
}

// RetryCount sets the number of retry attempts for failed requests.
func (h *Hog) RetryCount(count int) *Hog {
	h.retryCount = count
	return h
}

// LogLevel sets the logging verbosity when using the default logger.
func (h *Hog) LogLevel(level LogLevel) *Hog {
	if logger, ok := h.logger.(*defaultLogger); ok {
		logger.SetLevel(level)
	}
	return h
}

// getFullUrl constructs the complete URL with query parameters.
func getFullUrl(uri string, params url.Values) string {
	if params == nil {
		return uri
	}
	return fmt.Sprint(uri, "?", params.Encode())
}

// fillHeaders copies headers from source to destination.
func fillHeaders(dest, source http.Header) {
	for k, varr := range source {
		for _, v := range varr {
			dest.Add(k, v)
		}
	}
}
