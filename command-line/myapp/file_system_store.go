package myapp

import (
	"encoding/json"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)
	if err != nil {
		return nil, err
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, err
	}

	store := &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}
	return store, nil
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, err
	}

	closeFn := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		return nil, nil, err
	}
	return store, closeFn, nil
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	stat, err := os.Stat(file.Name())
	if err != nil {
		return err
	}

	if stat.Size() == 0 {
		file.Write([]byte(`[]`))
		file.Seek(0, 0)
	}
	return nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	p := f.league.Find(name)

	if p != nil {
		return p.Wins
	}
	return -1
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	p := f.league.Find(name)

	if p != nil {
		p.Wins++
	} else {
		f.league = append(f.league, Player{
			Name: name,
			Wins: 1,
		})
	}

	f.database.Encode(f.league)
}
