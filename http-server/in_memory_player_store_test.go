package httpserver_test

import (
	"sync"
	"testing"

	httpserver "github.com/gregor-pifko/learn-go-with-tests/http-server"
)

func TestInMemoryPlayerStoreConcurrently(t *testing.T) {
	t.Run("recording a win 3 times leaves it at 3", func(t *testing.T) {
		var (
			store  = httpserver.NewInMemoryPlayerStore()
			player = "Pepper"
		)

		store.RecordWin(player)
		store.RecordWin(player)
		store.RecordWin(player)

		assertScore(t, store.GetPlayerScore(player), 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		var (
			wantedWins = 1000
			store      = httpserver.NewInMemoryPlayerStore()
			player     = "Pepper"
		)

		var wg sync.WaitGroup

		for range wantedWins {
			wg.Go(func() {
				store.RecordWin(player)
			})
		}
		wg.Wait()

		assertScore(t, store.GetPlayerScore(player), wantedWins)
	})
}

func assertScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
