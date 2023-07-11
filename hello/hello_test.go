package main

import "testing"

func TestHello(t *testing.T) {
	asserCorrectMsg := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("expected %q, but got %q instead", want, got)
		}
	}

	t.Run("happy path", func(t *testing.T) {
		want := "hello mustafa!!"
		got := Hello()

		asserCorrectMsg(t, got, want)
	})
}
