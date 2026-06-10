package pokertest

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
	"time"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

func AssertStatus(t testing.TB, got *httptest.ResponseRecorder, want int) {
	t.Helper()
	if got.Code != want {
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

func AssertGameStartedWith(t testing.TB, game *SpyGame, want int) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == want
	})

	if !passed {
		t.Errorf("wanted Start called with %d but got %d", want, game.StartedWith)
	}
}

func AssertFinishCalledWith(t testing.TB, game *SpyGame, want string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinnishedWith == want
	})

	if !passed {
		t.Errorf("expected finish called with '%s' but got %q", want, game.FinnishedWith)
	}
}

func AssertGameNotStarted(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func AssertGameNotFinished(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.FinishCalled {
		t.Errorf("game should not have finished")
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)

	for time.Now().Before(deadline) {
		if f() {
			return true
		}
		time.Sleep(1 * time.Millisecond)
	}
	return false
}
