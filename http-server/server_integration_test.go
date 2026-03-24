package httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpserver "github.com/gregor-pifko/learn-go-with-tests/http-server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	var (
		store  = httpserver.NewInMemoryPlayerStore()
		server = httpserver.PlayerServer{store}
		player = "Pepper"
	)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
