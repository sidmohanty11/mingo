package minclient

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// initialization
	client := client{}

	cH := make(http.Header)
	cH.Set("Content-Type", "application/json")
	cH.Set("User-Agent", "mingo")
	client.Headers = cH

	// execution
	reqHeaders := make(http.Header)

	reqHeaders.Set("X-Request-Id", "123456")

	finalHeaders := client.getReqHeaders(reqHeaders)

	// validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received.")
	}
}
