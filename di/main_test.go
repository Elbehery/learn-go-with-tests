package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := &bytes.Buffer{}
	Greet(buf, "Mustafa")
	want := "Hello, Mustafa"
	act := buf.String()

	if want != act {
		t.Fatalf("expected %v, but got %v instead", want, act)
	}
}
