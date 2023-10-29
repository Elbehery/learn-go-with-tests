package maps

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	got := d.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, exp string) {
	t.Helper()

	if got != exp {
		t.Fatalf("expected %v, but got %v instead", exp, got)
	}
}
