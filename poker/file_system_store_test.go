package poker_test

import (
	"os"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
	"github.com/gregor-pifko/learn-go-with-tests/poker/pokertest"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league sorted", func(t *testing.T) {
		database := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store, err := poker.NewFileSystemPlayerStore(database)
		pokertest.AssertNoError(t, err)

		got := store.GetLeague()

		want := poker.League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		pokertest.AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		pokertest.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store, err := poker.NewFileSystemPlayerStore(database)
		pokertest.AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		pokertest.AssertScore(t, got, 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store, err := poker.NewFileSystemPlayerStore(database)
		pokertest.AssertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		pokertest.AssertScore(t, got, 34)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store, err := poker.NewFileSystemPlayerStore(database)
		pokertest.AssertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		pokertest.AssertScore(t, got, 1)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database := createTempFile(t, "")

		_, err := poker.NewFileSystemPlayerStore(database)

		pokertest.AssertNoError(t, err)
	})
}

func createTempFile(t testing.TB, initialData string) *os.File {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	t.Cleanup(func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	})

	return tmpfile
}
