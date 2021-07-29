package minclient

import (
	"net/http"
	"time"
)

type client struct {
	theClient          *http.Client
	Headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
}

func New() Client {
	// configure these numbers as your traffic pattern
	httpclient := &client{}

	return httpclient
}

type Client interface {
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	SetHeaders(headers http.Header)
	SetConnectionTimeout(timeout time.Duration)
	SetResponseTimeout(timeout time.Duration)
	SetMaxIdleConnections(i int)
	DisableTimeouts(b bool)
}

func (c *client) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *client) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

func (c *client) SetResponseTimeout(timeout time.Duration) {
	c.responseTimeout = timeout
}

func (c *client) SetMaxIdleConnections(i int) {
	c.maxIdleConnections = i
}

// false by default
func (c *client) DisableTimeouts(b bool) {
	c.disableTimeouts = b
}

func (c *client) Get(url string, headers http.Header) (*http.Response, error) {
	res, err := c.do(http.MethodGet, url, headers, nil)
	return res, err
}

func (c *client) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	res, err := c.do(http.MethodPost, url, headers, body)
	return res, err
}

func (c *client) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	res, err := c.do(http.MethodPut, url, headers, body)
	return res, err
}

func (c *client) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	res, err := c.do(http.MethodPatch, url, headers, body)
	return res, err
}

func (c *client) Delete(url string, headers http.Header) (*http.Response, error) {
	res, err := c.do(http.MethodDelete, url, headers, nil)
	return res, err
}
