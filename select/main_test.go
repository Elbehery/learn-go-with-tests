package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	slowSrvr := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(20 * time.Millisecond)
		writer.WriteHeader(http.StatusOK)
	}))

	fastSrvr := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowSrvr.URL
	fastUrl := fastSrvr.URL

	exp := fastUrl
	act := Racer(slowUrl, fastUrl)

	assertStrings(t, exp, act)
}

func assertStrings(t testing.TB, exp, got string) {
	t.Helper()

	if exp != got {
		t.Fatalf("expected %v, but got %v instead", exp, got)
	}
}
