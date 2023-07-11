package iteration

import "testing"

func TestRepeat2(t *testing.T) {
	want := "aaaaa"
	got := Repeat2("a")

	if got != want {
		t.Errorf("expected %q, but got %q instead", want, got)
	}
}
