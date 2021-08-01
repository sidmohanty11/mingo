package examples

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	mingo "github.com/sidmohanty11/mingo"
)

// func TestMain(m *testing.M) {
// 	fmt.Println("Let's start testing!!!")
// 	mingo.StartMockServer()
// 	os.Exit(m.Run())
// }

func TestGet(t *testing.T) {
	t.Run("testerrfetching", func(t *testing.T) {
		mingo.FlushMocks() // clean all of the prev mocks or flushes it.
		mock := mingo.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting req"),
		}
		mingo.AddMock(mock)

		endpoints, err := Get()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting req" {
			t.Error("invalid error message received")
		}
	})

	t.Run("testerrunmarshalingjson", func(t *testing.T) {
		mock := mingo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResStatusCode: http.StatusOK,
			ResBody:       `{"current_user_url": 123}`,
		}

		mingo.AddMock(mock)
		endpoints, err := Get()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message received")
		}
	})

	t.Run("testnoerr", func(t *testing.T) {
		mock := mingo.Mock{
			Method:        http.MethodGet,
			Url:           "https://api.github.com",
			ResStatusCode: http.StatusOK,
			ResBody:       `{"current_user_url": "https://api.github.com/user"}`,
		}

		mingo.AddMock(mock)

		endpoints, err := Get()

		if err != nil {
			t.Error(fmt.Sprintf("no error was expected and we got '%s'", err.Error()))
		}

		if endpoints == nil {
			t.Error("endpoints were expected and we got nil")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})
}
