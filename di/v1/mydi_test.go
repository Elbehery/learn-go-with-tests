package main

import (
	"bytes"
	"testing"
)

func TestMyGreet(t *testing.T) {
	w := bytes.Buffer{}

	want := "Hello, Mustafa"
	MyGreet(&w, "Mustafa")
	got := w.String()
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %v, but got %v instead", want, got)
	}
}
