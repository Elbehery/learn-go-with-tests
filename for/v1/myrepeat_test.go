package iteration

import "testing"

func TestRepeat2(t *testing.T) {
	want := "aaaaa"
	got := Repeat2("a", 5)

	if got != want {
		t.Errorf("expected %q, but got %q instead", want, got)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat2("a", 5)
	}
}
