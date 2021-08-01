package mingo

import (
	"net/http"
	"sync"
)

type client struct {
	theClient   *http.Client
	clientMaker *makeClient
	clientOnce  sync.Once
}

type Client interface {
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Options(url string, headers ...http.Header) (*Response, error)
}

func (c *client) Get(url string, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodGet, url, nil, attachHeaders(headers...))
	return res, err
}

func (c *client) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodPost, url, body, attachHeaders(headers...))
	return res, err
}

func (c *client) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodPut, url, body, attachHeaders(headers...))
	return res, err
}

func (c *client) Patch(url string, body interface{}, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodPatch, url, body, attachHeaders(headers...))
	return res, err
}

func (c *client) Delete(url string, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodDelete, url, nil, attachHeaders(headers...))
	return res, err
}

func (c *client) Options(url string, headers ...http.Header) (*Response, error) {
	res, err := c.do(http.MethodOptions, url, nil, attachHeaders(headers...))
	return res, err
}
