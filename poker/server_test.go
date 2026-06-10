package poker_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/gorilla/websocket"
	"github.com/gregor-pifko/learn-go-with-tests/poker"
	"github.com/gregor-pifko/learn-go-with-tests/poker/pokertest"
)

var (
	dummyGame = &pokertest.SpyGame{}
	tenMS     = 10 * time.Millisecond
)

func TestGETPlayers(t *testing.T) {
	store := &pokertest.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  40,
		},
	}
	server := mustMakePlayerServer(t, store, dummyGame)

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

			pokertest.AssertStatus(t, response, tt.expectedHTTPStatus)
			pokertest.AssertResponseBody(t, response.Body.String(), tt.expectedScore)
		})
	}
}

func TestStoreWins(t *testing.T) {
	store := &pokertest.StubPlayerStore{}
	server := mustMakePlayerServer(t, store, dummyGame)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		pokertest.AssertStatus(t, response, http.StatusAccepted)
		pokertest.AssertSingleWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := poker.League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Twest", 14},
		}

		store := &pokertest.StubPlayerStore{League: wantedLeague}
		server := mustMakePlayerServer(t, store, dummyGame)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		pokertest.AssertStatus(t, response, http.StatusOK)
		pokertest.AssertContentType(t, response.Result().Header, "application/json")
		pokertest.AssertLeague(t, got, wantedLeague)
	})
}

func TestGame(t *testing.T) {
	approvals.UseFolder("testdata")

	t.Run("GET /game renders correct HTML", func(t *testing.T) {
		server := mustMakePlayerServer(t, &pokertest.StubPlayerStore{}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		pokertest.AssertStatus(t, response, http.StatusOK)
		approvals.VerifyString(t, response.Body.String())
	})

	t.Run("start a game with 3 players and declare Ruth the winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 100"
		winner := "Ruth"

		game := &pokertest.SpyGame{BlindAlert: []byte(wantedBlindAlert)}
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		pokertest.AssertGameStartedWith(t, game, 3)
		pokertest.AssertFinishCalledWith(t, game, winner)
		within(t, tenMS, func() { assertWebsocketGotMsg(t, ws, wantedBlindAlert) })
	})
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game poker.Game) *poker.PlayerServer {
	t.Helper()
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	t.Helper()
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s: %v", url, err)
	}
	return ws
}

func writeWSMessage(t *testing.T, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection: %v", err)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{})

	go func() {
		assert()
		close(done)
	}()

	select {
	case <-time.After(d):
		t.Fatal("timed out")
	case <-done:
	}
}

func assertWebsocketGotMsg(t testing.TB, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()

	if string(msg) != want {
		t.Errorf("got blind alert %q; want %q", string(msg), want)
	}
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

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/game", nil)
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
