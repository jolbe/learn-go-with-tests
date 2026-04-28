package httpserver_test

import (
	"testing"

	httpserver "github.com/gregor-pifko/learn-go-with-tests/http-server"
)

func TestInMemoryPlayerStore(t *testing.T) {
	httpserver.PlayerStoreContract{NewStore: func(t testing.TB) httpserver.PlayerStore {
		return httpserver.NewInMemoryPlayerStore()
	}}.Test(t)
}

func TestFileSystemPlayerStore(t *testing.T) {
	httpserver.PlayerStoreContract{NewStore: func(t testing.TB) httpserver.PlayerStore {
		database := createTempFile(t, "[]")
		store, err := httpserver.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		return store
	}}.Test(t)
}
