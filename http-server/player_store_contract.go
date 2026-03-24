package httpserver

import (
	"testing"
)

type PlayerStoreContract struct {
	NewStore func() PlayerStore
}

func (p PlayerStoreContract) Test(t *testing.T) {
	t.Run("records wins for Pepper and Floyd", func(t *testing.T) {
		var (
			sut     = p.NewStore()
			player1 = "Pepper"
			player2 = "Floyd"
		)

		assertPlayerScore(t, sut.GetPlayerScore(player1), 0)
		assertPlayerScore(t, sut.GetPlayerScore(player2), 0)

		sut.RecordWin(player1)
		sut.RecordWin(player1)
		sut.RecordWin(player1)

		sut.RecordWin(player2)
		sut.RecordWin(player2)

		assertPlayerScore(t, sut.GetPlayerScore(player1), 3)
		assertPlayerScore(t, sut.GetPlayerScore(player2), 2)
	})
}

func assertPlayerScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("wrong score got %d; want %d", got, want)
	}
}
