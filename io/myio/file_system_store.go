package myio

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(player string) int {
	league := f.GetLeague()

	for _, l := range league {
		if l.Name == player {
			return l.Wins
		}
	}
	return -1
}

func (f *FileSystemPlayerStore) RecordWin(player string) {
	league := f.GetLeague()
	for i, l := range league {
		if l.Name == player {
			league[i].Wins++
		}
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(&league)
}
