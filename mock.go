package mingo

import (
	"fmt"
	"net/http"
)

type Mock struct {
	Method        string
	Url           string
	ReqBody       string
	Error         error
	ResBody       string
	ResStatusCode int
}

func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	res := Response{
		statusCode: m.ResStatusCode,
		body:       []byte(m.ResBody),
		status:     fmt.Sprintf("%d %s", m.ResStatusCode, http.StatusText(m.ResStatusCode)),
	}

	return &res, nil
}
