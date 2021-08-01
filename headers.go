package mingo

import "net/http"

// a helper function to attach all necessary headers, if present
func attachHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
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
