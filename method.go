package hog

type HMethod interface {
	getHog() *Hog
	getBody() interface{}
	getName() string
}
