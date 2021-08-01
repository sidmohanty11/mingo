package examples

import (
	"errors"

	mingo "github.com/sidmohanty11/mingo"
)

var m = mingo.MakeNewClient().Make()

type Endpoints struct {
	EventsUrl string `json:"events_url"`
	EmojisUrl string `json:"emojis_url"`
}

func GetEndp() (*Endpoints, error) {
	res, err := m.Get("https://api.github.com")

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
