package pokertest

import (
	"net/http"
	"slices"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d; want %d", got, want)
	}
}

func AssertScore(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wrong score, got %d; want %d", got, want)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q; want %q", got, want)
	}
}

func AssertLeague(t testing.TB, got, want poker.League) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("returned league table doesn't match, got %v; want %v", got, want)
	}
}

func AssertContentType(t testing.TB, header http.Header, want string) {
	t.Helper()
	if header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, header)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertSingleWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	if store.WinCalls[0] != winner {
		t.Errorf("did not store correct winner got %q; want %q", store.WinCalls[0], winner)
	}
}
