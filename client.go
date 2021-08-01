package mingo

import (
	"net/http"
	"sync"
)

// client struct is the overall carrier of all essential functions that constitute the http-client.
type client struct {
	// this is the http client that interacts with all the methods
	theClient *http.Client
	// this is the client builder which contains all the necessary config info that constitutes
	// the client struct
	clientMaker *makeClient
	// it ensures that client is generated only once, performing a concurrency check
	clientOnce sync.Once
}

// Client interface stores the reference to how the function actually looks like,
// contains all the basic methods => GET, POST, PUT, PACTH, DELETE, OPTIONS
type Client interface {
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Options(url string, headers ...http.Header) (*Response, error)
}

// the GET request calls the do function and returns the response object or any error
// basically, performs a GET request
func (c *client) Get(url string, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodGet, url, nil, attachHeaders(headers...))
	return res, err
}

// the POST request calls the do function and returns the response object or any error
// basically, performs a POST request
func (c *client) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodPost, url, body, attachHeaders(headers...))
	return res, err
}

// the PUT request calls the do function and returns the response object or any error
// basically, performs a PUT request
func (c *client) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodPut, url, body, attachHeaders(headers...))
	return res, err
}

// the PATCH request calls the do function and returns the response object or any error
// basically, performs a PATCH request
func (c *client) Patch(url string, body interface{}, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodPatch, url, body, attachHeaders(headers...))
	return res, err
}

// the DELETE request calls the do function and returns the response object or any error
// basically, performs a DELETE request
func (c *client) Delete(url string, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodDelete, url, nil, attachHeaders(headers...))
	return res, err
}

// the OPTIONS request calls the do function and returns the response object or any error
// basically, performs a OPTIONS request
func (c *client) Options(url string, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodOptions, url, nil, attachHeaders(headers...))
	return res, err
}
