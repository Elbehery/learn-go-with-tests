package myio

import (
	"reflect"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		exp := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		store := FileSystemPlayerStore{database: database}
		// read first time
		act := store.GetLeague()
		assertLeague(t, act, exp)
		// read second time
		act = store.GetLeague()
		assertLeague(t, act, exp)
	})
}

func assertLeague(t testing.TB, act, exp []Player) {
	t.Helper()

	if !reflect.DeepEqual(act, exp) {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
