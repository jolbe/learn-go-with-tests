package pokertest

import (
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

type PlayerStoreContract struct {
	NewStore func(testing.TB) poker.PlayerStore
}

func (p PlayerStoreContract) Test(t *testing.T) {
	t.Run("records wins for Pepper and Floyd", func(t *testing.T) {
		var (
			sut     = p.NewStore(t)
			player1 = "Pepper"
			player2 = "Floyd"
		)

		AssertScore(t, sut.GetPlayerScore(player1), 0)
		AssertScore(t, sut.GetPlayerScore(player2), 0)

		sut.RecordWin(player1)
		sut.RecordWin(player1)
		sut.RecordWin(player1)

		sut.RecordWin(player2)
		sut.RecordWin(player2)

		AssertScore(t, sut.GetPlayerScore(player1), 3)
		AssertScore(t, sut.GetPlayerScore(player2), 2)
		AssertLeague(t, sut.GetLeague(), poker.League{{player1, 3}, {player2, 2}})
	})
}
