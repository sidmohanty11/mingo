package minclient

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
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
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
