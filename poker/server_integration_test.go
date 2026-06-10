package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
	"github.com/gregor-pifko/learn-go-with-tests/poker/pokertest"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	var (
		// store  = poker.NewInMemoryPlayerStore() // in case you need to switch out the file database for the in-memory one
		database   = createTempFile(t, "[]")
		store, err = poker.NewFileSystemPlayerStore(database)
		server     = mustMakePlayerServer(t, store, dummyGame)
		player     = "Pepper"
	)
	pokertest.AssertNoError(t, err)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		pokertest.AssertStatus(t, response, http.StatusOK)

		pokertest.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		pokertest.AssertStatus(t, response, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := poker.League{
			{"Pepper", 3},
		}
		pokertest.AssertLeague(t, got, want)
	})
}
