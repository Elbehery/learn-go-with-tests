package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	slowSrvr := makeDelayedServer(20 * time.Millisecond)
	fastSrvr := makeDelayedServer(0 * time.Millisecond)

	defer slowSrvr.Close()
	defer fastSrvr.Close()

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

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(d)
		writer.WriteHeader(http.StatusOK)
	}))
}
