package client

import (
	"errors"
	"net/http"
)

func (c *client) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, errors.New("Unable to create a new request.")
	}

	res, err := client.Do(req)

	return res, err
}
