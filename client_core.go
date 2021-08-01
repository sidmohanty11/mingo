package mingo

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

// default values of some constants
const (
	defaultMaxIdleConnections = 10
	defaultResponseTimeout    = 10 * time.Second
	defaultConnectionTimeout  = 10 * time.Second
)

// returns the body as json/xml as referred in content-type header
// the default is json format
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

// the do function does all the hard work.
// takes in the method, url, body and headers and performs accordingly
func (c *client) do(method, url string, body interface{}, headers http.Header) (*Response, error) {
	// gets all the headers
	allHeaders := c.getReqHeaders(headers)

	// gets the request body or any error (json/xml)
	reqBody, err := c.getReqBody(body, allHeaders.Get("Content-Type"))

	if err != nil {
		return nil, err
	}

	// when testing, this part gets executed and ensures no API call is being made
	// gets a fake response body which is provided by you, usually
	if mock := mockSrv.getMock(method, url, string(reqBody)); mock != nil {
		return mock.GetResponse()
	}

	// else the api call is being made,
	// gets the request and the error
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))

	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	// sets the request headers
	req.Header = allHeaders

	// the httpclient
	myclient := c.getHttpClient()

	// performs the http.Client Do function which returns response or any error
	res, err := myclient.Do(req)

	if err != nil {
		return nil, err
	}

	// as the res.Body is a ReadAll function it must be closed
	defer res.Body.Close()

	// the respose body is being read
	rbody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// this is the final response object,
	// which is almost human readable and 10x nicer
	final := Response{
		status:     res.Status,
		statusCode: res.StatusCode,
		headers:    res.Header,
		body:       rbody,
	}

	// no error upto this, soooo
	// we got our final response
	return &final, nil
}

// returns the http client with all config set
func (c *client) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		// if there is a client already present then return
		if c.clientMaker.client != nil {
			c.theClient = c.clientMaker.client
			return
		}
		// else, set the client with required params
		c.theClient = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext:           (&net.Dialer{Timeout: c.getConnectionTimeout()}).DialContext,
			},
		}
	})
	// return the client
	return c.theClient
}

// gets the max idle connections to the client that is being set to the http.Client
func (c *client) getMaxIdleConnections() int {
	if c.clientMaker.maxIdleConnections > 0 {
		return c.clientMaker.maxIdleConnections
	}

	// by disabling timeouts
	if c.clientMaker.disableTimeouts {
		return 0
	}
	// no one configured
	return defaultMaxIdleConnections
}

// gets the max response timeout as time.Duration
func (c *client) getResponseTimeout() time.Duration {
	if c.clientMaker.responseTimeout > 0 {
		return c.clientMaker.responseTimeout
	}

	// by disabling timeouts
	if c.clientMaker.disableTimeouts {
		return 0
	}
	// no one configured
	return defaultResponseTimeout
}

// gets the max connection timeout as time.Duration
func (c *client) getConnectionTimeout() time.Duration {
	if c.clientMaker.connectionTimeout > 0 {
		return c.clientMaker.connectionTimeout
	}

	// by disabling timeouts
	if c.clientMaker.disableTimeouts {
		return 0
	}
	// no one configured
	return defaultConnectionTimeout
}
