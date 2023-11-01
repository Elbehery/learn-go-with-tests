package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		act, _ := Racer(slowUrl, fastUrl)
		exp := fastUrl

		assertStrings(t, act, exp)
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		srvr := makeDelayedServer(25 * time.Millisecond)
		url := srvr.URL

		_, err := ConfigurableRacer(url, url, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %v but got %v instead", want, got)
	}
}

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(d)
		writer.WriteHeader(http.StatusOK)
	}))
}
