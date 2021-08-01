package examples3

import (
	"errors"
	"net/http"
)

type Error struct {
	Message    string `json:"message"`
	DocURL     string `json:"documentation_url"`
	StatusCode int    `json:"-"`
}

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

func CreateRepo(req Repository) (*Repository, error) {
	res, err := httpClient.Post("https://api.github.com/user/repos", req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusCreated {
		var githubErr Error
		if err := res.UnmarshalJSON(&githubErr); err != nil {
			return nil, errors.New("error processing github error")
		}

		return nil, errors.New(githubErr.Message)
	}

	var r Repository
	if err := res.UnmarshalJSON(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
