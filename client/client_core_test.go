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

func TestGetRequestBody(t *testing.T) {
	// initialization
	client := client{}

	// execution
	t.Run("nobodynilresponse", func(t *testing.T) {
		body, err := client.getReqBody(nil, "")

		// validation
		if err != nil {
			t.Error("no error expected here (nil body)")
		}

		if body != nil {
			t.Error("no body expected here (nil body)")
		}
	})

	t.Run("bodywithjson", func(t *testing.T) {
		reqBody := []string{"one", "two"}
		body, err := client.getReqBody(reqBody, "application/json")

		// validation
		if err != nil {
			t.Error("no error expected marshaling (json body)")
		}

		if string(body) != `["one","two"]` {
			t.Error("no body expected here (marhsal body)")
		}
	})

	t.Run("bodywithxml", func(t *testing.T) {
		reqBody := []string{"one", "two"}
		body, err := client.getReqBody(reqBody, "application/xml")

		// validation
		if err != nil {
			t.Error("no error expected marshaling (xml body)")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("no body expected here (xml body)")
		}
	})

	t.Run("defaultbodythatisjson", func(t *testing.T) {
		reqBody := []string{"one", "two"}
		body, err := client.getReqBody(reqBody, "")

		// validation
		if err != nil {
			t.Error("no error expected here (default json body)")
		}

		if string(body) != `["one","two"]` {
			t.Error("no body expected marshaling (default json body)")
		}
	})
}
