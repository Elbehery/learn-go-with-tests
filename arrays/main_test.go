package main

import "testing"

func TestSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	act := Sum(input)
	exp := 15

	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
