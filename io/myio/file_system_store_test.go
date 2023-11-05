package myio

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database := `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`

	db, removeFunc := createTempFile(t, database)
	defer removeFunc()
	store := FileSystemPlayerStore{database: db}

	t.Run("league from a reader", func(t *testing.T) {

		exp := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		// read first time
		act := store.GetLeague()
		assertLeague(t, act, exp)
		// read second time
		act = store.GetLeague()
		assertLeague(t, act, exp)
	})

	t.Run("get player score", func(t *testing.T) {
		exp := 33
		player := "Chris"
		act := store.GetPlayerScore(player)
		if act != exp {
			t.Errorf("expected %v, but got %v instead", exp, act)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		exp := 34
		player := "Chris"
		store.RecordWin(player)
		act := store.GetPlayerScore(player)

		assertScoreEquals(t, act, exp)
	})
}

func assertLeague(t testing.TB, act, exp []Player) {
	t.Helper()

	if !reflect.DeepEqual(act, exp) {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}

func createTempFile(t testing.TB, data string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	f, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create database temp file, %v", err)
	}

	f.Write([]byte(data))

	removeFunc := func() {
		f.Close()
		os.Remove(f.Name())
	}

	return f, removeFunc
}

func assertScoreEquals(t testing.TB, act, exp int) {
	t.Helper()
	if act != exp {
		t.Errorf("got %d want %d", act, exp)
	}
}
