package examples

import (
	"errors"
	"net/http"
	"os"
	"testing"

	mingo "github.com/sidmohanty11/mingo"
)

func TestMain(m *testing.M) {
	mingo.StartMockServer()
	os.Exit(m.Run())
}

func TestGetEndPoints(t *testing.T) {
	t.Run("errorgettingfromapi", func(t *testing.T) {
		mingo.FlushMocks()

		mingo.AddMock(mingo.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting the endp"),
		})

		endp, err := GetEndp()

		if err == nil {
			t.Error("error was expected")
		}

		if endp != nil {
			t.Error("endpoint nil was expected")
		}

		if err.Error() != "timeout getting the endp" {
			t.Error("got different endp")
		}
	})

	t.Run("noerr", func(t *testing.T) {
		mingo.FlushMocks()

		mingo.AddMock(mingo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResStatusCode: http.StatusOK,
			ResBody:       `{"events_url":"https://api.github.com/events","emojis_url":"https://api.github.com/emojis"}`,
		})

		endp, err := GetEndp()

		if err != nil {
			t.Error("error was not expected")
		}

		if endp == nil {
			t.Error("endpoint nil was not expected")
		}

		if endp.EventsUrl != "https://api.github.com/events" {
			t.Error("got different values")
		}
	})

	t.Run("jsonerr", func(t *testing.T) {
		mingo.FlushMocks()

		mingo.AddMock(mingo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResStatusCode: http.StatusOK,
			ResBody:       `{"events_url": `,
		})

		endp, err := GetEndp()

		if err == nil {
			t.Error("error was expected")
		}

		if endp != nil {
			t.Error("endpoint nil was expected")
		}

		if err.Error() != "unexpected end of JSON input" {
			t.Error("got err")
		}
	})

}
