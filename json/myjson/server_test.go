package myjson

import (
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

func TestLeague(t *testing.T) {

	store := StubPlayerStore{map[string]int{}}
	srvr := PlayerServer{&store}

	t.Run("it returns 200 on /league", func(t *testing.T) {
		req := newLeagueGetRequest()
		resp := httptest.NewRecorder()

		srvr.ServeHTTP(resp, req)
		assertStatus(t, resp.Code, http.StatusOK)
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
