package main

// imports the http-client from mingo library
import (
	"fmt"
	"github.com/sidmohanty11/mingo/client"
	"io/ioutil"
	"net/http"
)

var myclient = getGithubClient() // initializes the client

func getGithubClient() minclient.Client {
	client := minclient.New()

	cH := make(http.Header)
	// cH.Set("Authorization", "Bearer 123ABC")

	client.SetHeaders(cH)

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
	u := User{"sid", "hello"}
	postUrls(u)
}

func getUrls() {
	res, err := myclient.Get("https://api.github.com", nil) // basic get method

	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode)

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes)) // res.Body
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"-"`
}

func postUrls(user User) {
	res, err := myclient.Post("https://api.github.com", nil, user) // basic get method

	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode)

	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes)) // res.Body
}
