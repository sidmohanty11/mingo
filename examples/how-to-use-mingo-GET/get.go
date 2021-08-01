package examples

import "fmt"

type url struct {
	CurrentUserUrl string `json:"current_user_url"`
	RepoUrl        string `json:"repository_url"`
}

func Get() (*url, error) {
	res, err := httpClient.Get("https://api.github.com", nil)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Status Code: %d", res.StatusCode())
	fmt.Printf("Status: %s", res.Status())
	fmt.Printf("Body: %s\n", res.BodyAsString())

	var endp url

	if err := res.UnmarshalJSON(&endp); err != nil {
		return nil, err
	}

	fmt.Printf("Repo URL: %s", endp.RepoUrl)
	fmt.Printf("User URL: %s", endp.CurrentUserUrl)

	return &endp, nil
}
