package myio

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(db *os.File) (*FileSystemPlayerStore, error) {
	db.Seek(0, 0)
	league, err := NewLeague(db)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, '%v'", db.Name(), err)
	}
	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&tape{db}),
		league:   league,
	}, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerWins(name string) int {
	player := f.league.Find(name)
	if player == nil {
		return -1
	}
	return player.Wins
}

func (f *FileSystemPlayerStore) RecordWins(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{
			Name: name,
			Wins: 1,
		})
	}

	f.Database.Encode(f.league)
}
