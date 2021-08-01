package mingo

import (
	"fmt"
	"net/http"
)

// Mock struct helps to provide a clean way to configure mock HTTP requests
// based on the combination of req method, URL and req body.
type Mock struct {
	// the http method
	Method string
	// the url
	Url string
	// request body
	ReqBody string
	// errors, if you want to attach any
	Error error
	// response body, ex - what it should be?
	ResBody string
	// response status code - what it should be?
	ResStatusCode int
}

// GetResponse returns a response object based on the mock config.
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	// creates a Response object with methods taken from "m"
	res := Response{
		statusCode: m.ResStatusCode,
		body:       []byte(m.ResBody),
		status:     fmt.Sprintf("%d %s", m.ResStatusCode, http.StatusText(m.ResStatusCode)),
	}

	return &res, nil
}
