package examples

// imports the http-client from mingo library
import (
	"github.com/sidmohanty11/mingo/client"
)

func basicEx() {
	client := client.New() // initializes the client
	client.Get()           // basic get method
}
