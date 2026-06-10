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

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 100"
		game := &pokertest.SpyGame{BlindAlert: []byte(wantedBlindAlert)}

		out := &bytes.Buffer{}
		in := userSends("3", "Chris wins")

		poker.NewCLI(in, out, game).PlayPoker()

		assertMessagesSentToUser(t, out, poker.PlayerPrompt, wantedBlindAlert)
		pokertest.AssertGameStartedWith(t, game, 3)
		pokertest.AssertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &pokertest.SpyGame{}

		in := userSends("8", "Cleo wins")

		poker.NewCLI(in, dummyStdOut, game).PlayPoker()

		pokertest.AssertGameStartedWith(t, game, 8)
		pokertest.AssertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &pokertest.SpyGame{}

		out := &bytes.Buffer{}
		in := userSends("Pies")

		poker.NewCLI(in, out, game).PlayPoker()

		pokertest.AssertGameNotStarted(t, game)
		assertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		game := &pokertest.SpyGame{}

		out := &bytes.Buffer{}
		in := userSends("5", "Lloyd is a killer")

		poker.NewCLI(in, out, game).PlayPoker()

		pokertest.AssertGameNotFinished(t, game)
		assertMessagesSentToUser(t, out, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
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
