package mingo

import (
	"net/http"
	"time"
)

// MakeClient contains all the config info
// and it also makes the client.
type MakeClient interface {
	SetHeaders(headers http.Header) MakeClient
	SetConnectionTimeout(timeout time.Duration) MakeClient
	SetResponseTimeout(timeout time.Duration) MakeClient
	SetMaxIdleConnections(i int) MakeClient
	DisableTimeouts(b bool) MakeClient
	SetHttpClient(c *http.Client) MakeClient
	SetUserAgent(userAgent string) MakeClient
	Make() Client
}

// the struct which is responsible for storing all the necessary configurations
// for making the request
type makeClient struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
	client             *http.Client
	userAgent          string
}

// MakeNewClient initializes the makeClient struct which now has all the methods
// which can be set on it
func MakeNewClient() MakeClient {
	makeclient := &makeClient{}

	return makeclient
}

// this is the final call where the actual Client is being made
// which performs all the requests (GET,POST,...)
func (c *makeClient) Make() Client {
	myc := client{
		clientMaker: c,
	}

	return &myc
}

// Sets the Headers into the makeClient struct, Make function must be
// called to get the actual client
func (c *makeClient) SetHeaders(headers http.Header) MakeClient {
	c.headers = headers
	return c
}

// Sets the Connection Timeout into the makeClient struct,
// Make function must be called to get the actual client
func (c *makeClient) SetConnectionTimeout(timeout time.Duration) MakeClient {
	c.connectionTimeout = timeout
	return c
}

// Sets the Response Timeout into the makeClient struct,
// Make function must be called to get the actual client
func (c *makeClient) SetResponseTimeout(timeout time.Duration) MakeClient {
	c.responseTimeout = timeout
	return c
}

// Sets the Max Idle Connections into the makeClient struct,
// Make function must be called to get the actual client
func (c *makeClient) SetMaxIdleConnections(i int) MakeClient {
	c.maxIdleConnections = i
	return c
}

// Disables all the Timeout features into the makeClient struct,
// Make function must be called to get the actual client
// false by default
func (c *makeClient) DisableTimeouts(b bool) MakeClient {
	c.disableTimeouts = b
	return c
}

// when you pass a client already configured, then no need to set all the configs
// so this sets client to the client you pass in
func (c *makeClient) SetHttpClient(client *http.Client) MakeClient {
	c.client = client
	return c
}

// sets the user agent, which can be used in headers when making requests
func (c *makeClient) SetUserAgent(userAgent string) MakeClient {
	c.userAgent = userAgent
	return c
}
