package myjson

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	store map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	return s.store[player]
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.store[player]++
}

func (s *StubPlayerStore) GetLeague() []Player {
	var players []Player
	for k, v := range s.store {
		players = append(players, Player{
			Name: k,
			Wins: v,
		})
	}
	return players
}

func TestLeague(t *testing.T) {

	store := StubPlayerStore{map[string]int{}}
	srvr := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		req := newLeagueGetRequest()
		resp := httptest.NewRecorder()

		srvr.ServeHTTP(resp, req)
		var players []Player
		err := json.NewDecoder(resp.Body).Decode(&players)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", resp.Body, err)
		}

		assertStatus(t, resp.Code, http.StatusOK)
		assertContentType(t, resp, "application/json")
	})
}

func newLeagueGetRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}

func assertStatus(t testing.TB, act, exp int) {
	t.Helper()
	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}

func assertContentType(t testing.TB, resp *httptest.ResponseRecorder, exp string) {
	t.Helper()
	if resp.Result().Header.Get("content-type") != exp {
		t.Errorf("response did not have content-type of application/json, got %v", resp.Result().Header)
	}
}
