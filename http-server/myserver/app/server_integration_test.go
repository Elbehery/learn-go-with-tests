package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type InMemoryPlayerStore struct {
	store map[string]int
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.store[name]++
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	srvr := PlayerServer{Store: store}
	player := "Mustafa"

	srvr.ServerHttp(httptest.NewRecorder(), newPostWinRequest(player))
	srvr.ServerHttp(httptest.NewRecorder(), newPostWinRequest(player))
	srvr.ServerHttp(httptest.NewRecorder(), newPostWinRequest(player))

	resp := httptest.NewRecorder()
	srvr.ServerHttp(resp, newGetWinRequest(player))

	assertHttpStatus(t, resp.Code, http.StatusOK)
	assertString(t, resp.Body.String(), "3")
}

func newGetWinRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", name), nil)
}

func newPostWinRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", name), nil)
}

func assertHttpStatus(t testing.TB, act, exp int) {
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
