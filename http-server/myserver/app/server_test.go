package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/players/Pepper", nil)
		resp := httptest.NewRecorder()

		PlayerServer(resp, req)
		act := resp.Body.String()
		exp := "20"

		assertStrings(t, exp, act)
	})
}

func assertStrings(t testing.TB, act, exp string) {
	t.Helper()

	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
