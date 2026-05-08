package poker_test

import (
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
	"github.com/gregor-pifko/learn-go-with-tests/poker/pokertest"
)

func TestInMemoryPlayerStore(t *testing.T) {
	pokertest.PlayerStoreContract{NewStore: func(t testing.TB) poker.PlayerStore {
		return poker.NewInMemoryPlayerStore()
	}}.Test(t)
}

func TestFileSystemPlayerStore(t *testing.T) {
	pokertest.PlayerStoreContract{NewStore: func(t testing.TB) poker.PlayerStore {
		database := createTempFile(t, "[]")
		store, err := poker.NewFileSystemPlayerStore(database)
		pokertest.AssertNoError(t, err)

		return store
	}}.Test(t)
}
