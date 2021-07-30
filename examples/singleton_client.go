package examples

import (
	"time"

	mingo "github.com/sidmohanty11/mingo/client"
)

var httpClient = getClient()

func getClient() mingo.Client {
	return mingo.
		MakeNewClient().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		SetConnectionTimeout(2 * time.Second).
		Make()
}
