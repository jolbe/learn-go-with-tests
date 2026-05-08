package poker_test

import (
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

func TestInMemoryPlayerStore(t *testing.T) {
	poker.PlayerStoreContract{NewStore: func(t testing.TB) poker.PlayerStore {
		return poker.NewInMemoryPlayerStore()
	}}.Test(t)
}

func TestFileSystemPlayerStore(t *testing.T) {
	poker.PlayerStoreContract{NewStore: func(t testing.TB) poker.PlayerStore {
		database := createTempFile(t, "[]")
		store, err := poker.NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		return store
	}}.Test(t)
}
