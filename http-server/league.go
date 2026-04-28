package httpserver

import (
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"slices"
)

type League []Player

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league: %w", err)
	}
	return league, err
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (l League) Sort() League {
	slices.SortFunc(l, func(a, b Player) int {
		return cmp.Compare(b.Wins, a.Wins)
	})
	return l
}
