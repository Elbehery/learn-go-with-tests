package myapp

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	input := strings.NewReader("Mustafa\n")
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore: playerStore, in: input}
	cli.PlayPoker()

	act := playerStore.WinCalls[0]
	exp := "Mustafa"

	assertStrings(t, act, exp)
}

func assertStrings(t testing.TB, act, exp string) {
	t.Helper()
	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
