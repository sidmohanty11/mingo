package minclient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
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
	client := http.Client{}
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

	res, err := client.Do(req)

	return res, err
}

func (c *client) getReqHeaders(reqHeaders http.Header) http.Header {
	res := make(http.Header)

	// setting common headers to the request. <<DEFAULT_HEADERS>>
	for header, value := range c.Headers {
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
