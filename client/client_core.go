package minclient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultMaxIdleConnections = 10
	defaultResponseTimeout    = 10 * time.Second
	defaultConnectionTimeout  = 10 * time.Second
)

func (c *client) getReqBody(body interface{}, contentType string) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	// by default we assume body is json
	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}

func (c *client) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {
	allHeaders := c.getReqHeaders(headers)

	reqBody, err := c.getReqBody(body, allHeaders.Get("Content-Type"))

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	req.Header = allHeaders

	myclient := c.getHttpClient()

	res, err := myclient.Do(req)

	return res, err
}

func (c *client) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		c.theClient = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext:           (&net.Dialer{Timeout: c.getConnectionTimeout()}).DialContext,
			},
		}
	})
	return c.theClient
}

func (c *client) getMaxIdleConnections() int {
	if c.clientMaker.maxIdleConnections > 0 {
		return c.clientMaker.maxIdleConnections
	}

	if c.clientMaker.disableTimeouts {
		return 0
	}
	// no one configured
	return defaultMaxIdleConnections
}

func (c *client) getResponseTimeout() time.Duration {
	if c.clientMaker.responseTimeout > 0 {
		return c.clientMaker.responseTimeout
	}

	if c.clientMaker.disableTimeouts {
		return 0
	}
	// no one configured
	return defaultResponseTimeout
}

func (c *client) getConnectionTimeout() time.Duration {
	if c.clientMaker.connectionTimeout > 0 {
		return c.clientMaker.connectionTimeout
	}

	if c.clientMaker.disableTimeouts {
		return 0
	}
	// no one configured
	return defaultConnectionTimeout
}

func (c *client) getReqHeaders(reqHeaders http.Header) http.Header {
	res := make(http.Header)

	// setting common headers to the request. <<DEFAULT_HEADERS>>
	for header, value := range c.clientMaker.headers {
		if len(value) > 0 {
			res.Set(header, value[0])
		}
	}

	// setting custom headers to the request.
	for key, value := range reqHeaders {
		if len(value) > 0 {
			res.Set(key, value[0])
		}
	}

	return res
}
