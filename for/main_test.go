package main

import "testing"

func TestRepeat(t *testing.T) {
	act := Repeat("a")
	exp := "aaaaa"

	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
