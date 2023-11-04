package myio

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

func (s *StubPlayerStore) GetLeague() []Player {
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
}

func newGetScoreRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
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
