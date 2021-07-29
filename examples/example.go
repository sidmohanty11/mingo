package main

// imports the http-client from mingo library
import (
	"fmt"

	mingo "github.com/sidmohanty11/mingo/client"
)

var myclient = getMingoClient() // initializes the client (singleton)

func getMingoClient() mingo.Client {
	client := mingo.
		MakeNewClient().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Make()

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
	postUrls()
}

func getUrls() {
	res, err := myclient.Get("https://api.github.com", nil) // basic get method

	if err != nil {
		panic(err)
	}

	fmt.Println(res.Status())
	fmt.Println(res.StatusCode())
	fmt.Println(res.BodyAsString())
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"-"`
}

func postUrls() {
	u := User{Name: "sid", Password: "hello"}
	res, err := myclient.Post("https://api.github.com", nil, u)

	if err != nil {
		panic(err)
	}

	fmt.Println(res.Status())
	fmt.Println(res.StatusCode())
	fmt.Println(res.BodyAsString())
}
