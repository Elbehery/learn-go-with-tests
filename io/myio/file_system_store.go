package myio

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	Database io.Writer
	league   League
}

func NewFileSystemPlayerStore(db *os.File) *FileSystemPlayerStore {
	db.Seek(0, 0)
	league, _ := NewLeague(db)
	return &FileSystemPlayerStore{
		Database: &tape{db},
		league:   league,
	}
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
	league := f.league
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{
			Name: name,
			Wins: 1,
		})
	}

	json.NewEncoder(f.Database).Encode(league)
}
