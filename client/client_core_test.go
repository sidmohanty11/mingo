package mingo

import (
	"testing"
)

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
