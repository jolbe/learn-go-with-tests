package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	var (
		// store  = poker.NewInMemoryPlayerStore() // in case you need to switch out the file database for the in-memory one
		database   = createTempFile(t, "[]")
		store, err = poker.NewFileSystemPlayerStore(database)
		server     = poker.NewPlayerServer(store)
		player     = "Pepper"
	)
	assertNoError(t, err)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
