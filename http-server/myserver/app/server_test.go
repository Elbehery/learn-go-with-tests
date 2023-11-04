package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)
		act := resp.Body.String()
		exp := "20"

		assertStrings(t, exp, act)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)
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
