package httpserver_test

import (
	"testing"

	httpserver "github.com/gregor-pifko/learn-go-with-tests/http-server"
)

func TestInMemoryPlayerStore(t *testing.T) {
	httpserver.PlayerStoreContract{NewStore: func() httpserver.PlayerStore {
		return httpserver.NewInMemoryPlayerStore()
	}}.Test(t)
}
