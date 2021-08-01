package examples

import (
	"fmt"
	"testing"
)

type Repository struct {
	Name string `json:"name"`
}

func TestPost(t *testing.T) {
	repo := Repository{
		Name: "hello",
	}
	res, err := httpClient.Post("", nil, repo)
	fmt.Println(res, err)
}
