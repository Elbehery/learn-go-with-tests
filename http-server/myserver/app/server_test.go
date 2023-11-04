package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	store map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func TestGetPlayers(t *testing.T) {
	store := &StubPlayerStore{store: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}

	svr := &PlayerServer{store: store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		svr.ServerHttp(resp, req)
		act := resp.Body.String()
		exp := "20"

		assertStrings(t, exp, act)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		svr.ServerHttp(resp, req)
		act := resp.Body.String()
		exp := "10"

		assertStrings(t, exp, act)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
	return req
}

func assertStrings(t testing.TB, exp, act string) {
	t.Helper()

	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
