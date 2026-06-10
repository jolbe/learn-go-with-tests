package pokertest

import (
	"fmt"
	"io"
	"time"

	"github.com/gregor-pifko/learn-go-with-tests/poker"
)

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   poker.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() poker.League {
	return s.League
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, alertsDestination io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

type SpyGame struct {
	StartCalled bool
	StartedWith int
	BlindAlert  []byte

	FinishCalled  bool
	FinnishedWith string
}

func (g *SpyGame) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *SpyGame) Finish(winner string) {
	g.FinishCalled = true
	g.FinnishedWith = winner
}
