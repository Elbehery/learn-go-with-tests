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

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting info for file %s: '%v'", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte(`[]`))
		file.Seek(0, 0)
	}

	return nil
}
func NewFileSystemPlayerStore(db *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(db)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}
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
