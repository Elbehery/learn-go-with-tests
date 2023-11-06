package myio

import (
	"os"
	"reflect"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	Database := `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`

	db, removeFunc := createTempFile(t, Database)
	defer removeFunc()
	store, err := NewFileSystemPlayerStore(db)
	assertNoError(t, err)

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
		act := store.GetPlayerWins(player)
		if act != exp {
			t.Errorf("expected %v, but got %v instead", exp, act)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		exp := 34
		player := "Chris"
		store.RecordWins(player)
		act := store.GetPlayerWins(player)

		assertScoreEquals(t, act, exp)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		newPlayer := "Mustafa"
		store.RecordWins(newPlayer)

		exp := 1
		act := store.GetPlayerWins(newPlayer)

		assertScoreEquals(t, act, exp)
	})
}

func assertLeague(t testing.TB, act, exp []Player) {
	t.Helper()

	if !reflect.DeepEqual(act, exp) {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}

func createTempFile(t testing.TB, data string) (*os.File, func()) {
	t.Helper()

	f, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create Database temp file, %v", err)
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

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error, but got '%v' instead", err)
	}
}
