package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
	"github.com/gregor-pifko/learn-go-with-tests/poker/pokertest"
)

var (
	dummyBlindAlerter = &pokertest.SpyBlindAlerter{}
	dummyPlayerStore  = &pokertest.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

type SpyGame struct {
	StartCalled bool
	StartedWith int

	FinishCalled  bool
	FinnishedWith string
}

func (g *SpyGame) Start(numberOfPlayers int) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *SpyGame) Finish(winner string) {
	g.FinishCalled = true
	g.FinnishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		in := userSends("3", "Chris wins")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &SpyGame{}

		in := userSends("8", "Cleo wins")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		in := userSends("Pies")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		game := &SpyGame{}

		in := userSends("5", "Lloyd is a killer")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})
}

func userSends(inputs ...string) io.Reader {
	return strings.NewReader(strings.Join(inputs, "\n"))
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	got := stdout.String()
	want := strings.Join(messages, "")

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertGameStartedWith(t testing.TB, game *SpyGame, want int) {
	t.Helper()
	if game.StartedWith != want {
		t.Errorf("wanted Start called with %d but got %d", want, game.StartedWith)
	}
}

func assertFinishCalledWith(t testing.TB, game *SpyGame, want string) {
	t.Helper()
	if game.FinnishedWith != want {
		t.Errorf("expected finish called with '%s' but got %q", want, game.FinnishedWith)
	}
}

func assertGameNotStarted(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameNotFinished(t testing.TB, game *SpyGame) {
	t.Helper()
	if game.FinishCalled {
		t.Errorf("game should not have finished")
	}
}
