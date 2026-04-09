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
		assertLeague(t, store.GetLeague(), []httpserver.Player{{"Pepper", 3}})
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
		assertLeague(t, store.GetLeague(), []httpserver.Player{{"Pepper", wantedWins}})
	})

	t.Run("it runs league safely during writes", func(t *testing.T) {
		var (
			store  = httpserver.NewInMemoryPlayerStore()
			player = "Pepper"
		)

		var wg sync.WaitGroup

		for range 1000 {
			wg.Go(func() {
				store.RecordWin(player)
			})
			wg.Go(func() {
				store.GetLeague()
			})
		}
		wg.Wait()
	})
}


