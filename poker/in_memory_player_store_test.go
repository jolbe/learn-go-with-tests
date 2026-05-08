package poker_test

import (
	"sync"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
	"github.com/gregor-pifko/learn-go-with-tests/poker/pokertest"
)

func TestInMemoryPlayerStoreConcurrently(t *testing.T) {
	t.Run("recording a win 3 times leaves it at 3", func(t *testing.T) {
		var (
			store  = poker.NewInMemoryPlayerStore()
			player = "Pepper"
		)

		store.RecordWin(player)
		store.RecordWin(player)
		store.RecordWin(player)

		pokertest.AssertScore(t, store.GetPlayerScore(player), 3)
		pokertest.AssertLeague(t, store.GetLeague(), poker.League{{"Pepper", 3}})
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		var (
			wantedWins = 1000
			store      = poker.NewInMemoryPlayerStore()
			player     = "Pepper"
		)

		var wg sync.WaitGroup

		for range wantedWins {
			wg.Go(func() {
				store.RecordWin(player)
			})
		}
		wg.Wait()

		pokertest.AssertScore(t, store.GetPlayerScore(player), wantedWins)
		pokertest.AssertLeague(t, store.GetLeague(), poker.League{{"Pepper", wantedWins}})
	})

	t.Run("it runs league safely during writes", func(t *testing.T) {
		var (
			store  = poker.NewInMemoryPlayerStore()
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
