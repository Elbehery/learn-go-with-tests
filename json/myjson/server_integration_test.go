package myjson

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	srvr := NewPlayerServer(store)
	player := "Mustafa"

	srvr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	srvr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	srvr.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		resp := httptest.NewRecorder()
		srvr.ServeHTTP(resp, newGetRequest(player))

		assertStatusCode(t, resp.Code, http.StatusOK)
		assertStrings(t, resp.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		exp := []Player{
			{"Mustafa", 3},
		}

		resp := httptest.NewRecorder()
		srvr.ServeHTTP(resp, newLeagueRequest())
		var res []Player
		err := json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			t.Fatalf("error decoding")
		}
		if !reflect.DeepEqual(res, exp) {
			t.Fatalf("error result")
		}
	})
}

func newPostWinRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodPost, fmt.Sprintf("/player/%s", player), nil)
}

func newGetRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/player/%s", player), nil)
}

func newLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}

func assertStatusCode(t testing.TB, act, exp int) {
	t.Helper()
	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}

func assertStrings(t testing.TB, act, exp string) {
	t.Helper()
	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
