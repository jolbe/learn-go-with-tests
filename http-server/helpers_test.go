package httpserver_test

import (
	"bytes"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"

	httpserver "github.com/gregor-pifko/learn-go-with-tests/http-server"
)

func getLeagueFromResponse(t testing.TB, buf *bytes.Buffer) []httpserver.Player {
	t.Helper()

	body := buf.String()
	league, err := httpserver.NewLeague(strings.NewReader(body))
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d; want %d", got, want)
	}
}

func assertScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q; want %q", got, want)
	}
}

func assertLeague(t testing.TB, got, want []httpserver.Player) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("returned league table doesn't match, got %v; want %v", got, want)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
