package pokertest

import (
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
