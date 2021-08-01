package examples3

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	mingo "github.com/sidmohanty11/mingo"
)

func TestCreateRepo(t *testing.T) {
	t.Run("timeoutfromgithub", func(t *testing.T) {
		mingo.FlushMocks()
		mingo.AddMock(mingo.Mock{
			Method:  http.MethodPost,
			Url:     "https://api.github.com/user/repos",
			ReqBody: `{"name":"test", "private":true, "description":"test"}`,
			Error:   errors.New("timeout from github"),
		})

		r := Repository{
			Name:        "test",
			Description: "test",
			Private:     true,
		}

		repo, err := CreateRepo(r)

		if repo != nil {
			t.Error("no repo expected at timeout from github")
		}

		if err == nil {
			t.Error("an error is expected when we get a timeout")
		}

		if err.Error() != "timeout from github" {
			fmt.Println(err.Error())
			t.Error("invalid error message")
		}
	})

	t.Run("noerr", func(t *testing.T) {
		mingo.FlushMocks()
		mingo.AddMock(mingo.Mock{
			Method:        http.MethodPost,
			Url:           "https://api.github.com/user/repos",
			ReqBody:       `{"name":"test", "private":true, "description":"test"}`,
			ResStatusCode: http.StatusCreated,
			ResBody:       `{"id":1, "name":"test"}`,
		})

		r := Repository{
			Name:        "test",
			Description: "test",
			Private:     true,
		}

		repo, err := CreateRepo(r)

		if repo == nil {
			t.Error("a valid repo was expected")
		}

		if err != nil {
			t.Error("no error expected")
		}

		if repo.Name != r.Name {
			t.Error("invalid repo name")
		}
	})

}
