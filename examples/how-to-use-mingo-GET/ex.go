package examples

import (
	"errors"

	mingo "github.com/sidmohanty11/mingo"
)

var httpClient = mingo.MakeNewClient().Make()

type Endpoints struct {
	EmailsUrl string `json:"emails_url"`
	EmojisUrl string `json:"emojis_url"`
}

func GetEndp() (*Endpoints, error) {
	res, err := httpClient.Get("https://api.github.com", nil)

	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		return nil, errors.New("error")
	}

	var endp Endpoints

	if err := res.UnmarshalJSON(&endp); err != nil {
		return nil, err
	}

	return &endp, nil
}
