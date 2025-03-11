package hog

// HMethod defines the interface for HTTP methods with common functionality.
type HMethod interface {
	getHog() *Hog
	getBody() interface{}
	getName() string
}
