package myio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores  map[string]int
	winCall []string
	league  []Player
}

func (s *StubPlayerStore) GetPlayerWins(player string) int {
	return s.scores[player]
}

func (s *StubPlayerStore) RecordWins(player string) {
	s.winCall = append(s.winCall, player)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func TestGetPlayer(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	srvr := NewPlayerServer(store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		srvr.ServeHTTP(resp, req)
		assertStatus(t, resp.Code, http.StatusOK)
		assertString(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		srvr.ServeHTTP(resp, req)
		assertStatus(t, resp.Code, http.StatusOK)
		assertString(t, resp.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("mustii")
		resp := httptest.NewRecorder()

		srvr.ServeHTTP(resp, req)
		assertStatus(t, resp.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{scores: map[string]int{}}
	srvr := NewPlayerServer(store)

	t.Run("it records wins on POST", func(t *testing.T) {
		req := newPostWinRequest("Mustafa")
		resp := httptest.NewRecorder()

		srvr.ServeHTTP(resp, req)
		assertStatus(t, resp.Code, http.StatusAccepted)

		if len(store.winCall) != 1 {
			t.Fatal()
		}

		if store.winCall[0] != "Mustafa" {
			t.Fatal()
		}
	})
}

func TestLeague(t *testing.T) {
	store := &StubPlayerStore{league: []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14}}}

	srvr := NewPlayerServer(store)

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		req := newLeagueRequest()
		resp := httptest.NewRecorder()

		srvr.ServeHTTP(resp, req)
		players := getLeagueFromResponse(t, resp.Body)
		assertStatus(t, resp.Code, http.StatusOK)
		if !reflect.DeepEqual(players, store.league) {
			t.Errorf("expected %v, but got %v instead", store.league, players)
		}
		if resp.Header().Get("content-type") != jsonContentType {
			t.Fatal()
		}
	})
}

func newGetScoreRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
}

func newPostWinRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
}

func newLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}

func getLeagueFromResponse(t testing.TB, data io.Reader) []Player {
	t.Helper()
	var players []Player
	err := json.NewDecoder(data).Decode(&players)
	if err != nil {
		t.Fatalf("can not decode json data")
	}
	return players
}

func assertStatus(t testing.TB, act, exp int) {
	t.Helper()
	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}

func assertString(t testing.TB, act, exp string) {
	t.Helper()
	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
