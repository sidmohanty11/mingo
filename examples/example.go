package main

// imports the http-client from mingo library
import (
	"fmt"

	"github.com/sidmohanty11/mingo/client"
)

func main() {
	client := client.New()                                // initializes the client
	res, err := client.Get("https://api.github.com", nil) // basic get method

	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode)
}
