package myapp

import (
	"encoding/json"
	"fmt"
	"io"
)

// Player stores player name with number of wins.
type Player struct {
	Name string
	Wins int
}

// League stores a collection of players.
type League []Player

// Find searches League for a player by name.
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// NewLeague creates a League from JSON.
func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("problem parsing json file: '%v'", err)
	}
	return league, nil
}
