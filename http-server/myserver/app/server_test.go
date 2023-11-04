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

	svr := &PlayerServer{Store: store}

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

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("Apollo")
		resp := httptest.NewRecorder()

		svr.ServerHttp(resp, req)
		act := resp.Code
		exp := http.StatusNotFound

		assertStatusCode(t, exp, act)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{map[string]int{}}
	svr := &PlayerServer{store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		req := newPostScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		svr.ServerHttp(resp, req)
		act := resp.Code
		exp := http.StatusAccepted
		assertStatusCode(t, exp, act)
	})
}

func newPostScoreRequest(name string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", name), nil)
	return req
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

func assertStatusCode(t testing.TB, exp, act int) {
	t.Helper()

	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
