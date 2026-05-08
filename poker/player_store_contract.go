package poker

import (
	"slices"
	"testing"
)

type PlayerStoreContract struct {
	NewStore func(testing.TB) PlayerStore
}

func (p PlayerStoreContract) Test(t *testing.T) {
	t.Run("records wins for Pepper and Floyd", func(t *testing.T) {
		var (
			sut     = p.NewStore(t)
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
		assertLeague(t, sut.GetLeague(), League{{player1, 3}, {player2, 2}})
	})
}

func assertPlayerScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("wrong score got %d; want %d", got, want)
	}
}

func assertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("returned league table doesn't match, got %v; want %v", got, want)
	}
}
