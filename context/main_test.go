package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response  string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "hello world !!!"
	store := StubStore{response: data}
	svr := Server(&store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	exp := data
	act := response.Body.String()

	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Fatal()
		}
	})
}
