package poker_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
	"github.com/gregor-pifko/learn-go-with-tests/poker/pokertest"
)

func TestGETPlayers(t *testing.T) {
	store := pokertest.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  40,
		},
	}
	server := poker.NewPlayerServer(&store)

	tests := []struct {
		name               string
		player             string
		expectedHTTPStatus int
		expectedScore      string
	}{
		{
			name:               "Returns Pepper's score",
			player:             "Pepper",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "20",
		},
		{
			name:               "Returns Floyd's score",
			player:             "Floyd",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "40",
		},
		{
			name:               "Returns 404 on missing players",
			player:             "Apollo",
			expectedHTTPStatus: http.StatusNotFound,
			expectedScore:      "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newGetScoreRequest(tt.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			pokertest.AssertStatus(t, response.Code, tt.expectedHTTPStatus)
			pokertest.AssertResponseBody(t, response.Body.String(), tt.expectedScore)
		})
	}
}

func TestStoreWins(t *testing.T) {
	store := pokertest.StubPlayerStore{}
	server := poker.NewPlayerServer(&store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		pokertest.AssertStatus(t, response.Code, http.StatusAccepted)
		pokertest.AssertSingleWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := poker.League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Twest", 14},
		}

		store := pokertest.StubPlayerStore{League: wantedLeague}
		server := poker.NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		pokertest.AssertStatus(t, response.Code, http.StatusOK)
		pokertest.AssertContentType(t, response.Result().Header, "application/json")
		pokertest.AssertLeague(t, got, wantedLeague)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return req
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t testing.TB, buf *bytes.Buffer) poker.League {
	t.Helper()

	body := buf.String()
	league, err := poker.NewLeague(strings.NewReader(body))
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}
