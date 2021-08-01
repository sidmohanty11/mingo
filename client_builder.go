package mingo

import (
	"net/http"
	"time"
)

type MakeClient interface {
	SetHeaders(headers http.Header) MakeClient
	SetConnectionTimeout(timeout time.Duration) MakeClient
	SetResponseTimeout(timeout time.Duration) MakeClient
	SetMaxIdleConnections(i int) MakeClient
	DisableTimeouts(b bool) MakeClient
	SetHttpClient(c *http.Client) MakeClient
	Make() Client
}

type makeClient struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
	client             *http.Client
	baseURL            string
}

func MakeNewClient() MakeClient {
	makeclient := &makeClient{}

	return makeclient
}

func (c *makeClient) Make() Client {
	myc := client{
		clientMaker: c,
	}

	return &myc
}

func (c *makeClient) SetHeaders(headers http.Header) MakeClient {
	c.headers = headers
	return c
}

func (c *makeClient) SetConnectionTimeout(timeout time.Duration) MakeClient {
	c.connectionTimeout = timeout
	return c
}

func (c *makeClient) SetResponseTimeout(timeout time.Duration) MakeClient {
	c.responseTimeout = timeout
	return c
}

func (c *makeClient) SetMaxIdleConnections(i int) MakeClient {
	c.maxIdleConnections = i
	return c
}

// false by default
func (c *makeClient) DisableTimeouts(b bool) MakeClient {
	c.disableTimeouts = b
	return c
}

func (c *makeClient) SetHttpClient(client *http.Client) MakeClient {
	c.client = client
	return c
}
