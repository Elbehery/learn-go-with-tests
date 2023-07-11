package main

import "testing"

func TestHello(t *testing.T) {
	want := "hello mustafa!!"
	got := Hello()

	if got != want {
		t.Errorf("expected '%v', but got '%v' instead", want, got)
	}
}
