package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	act := Sum(input)
	exp := 15

	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}

func TestSumAll(t *testing.T) {
	act := SumAll([]int{1, 2}, []int{0, 9})
	exp := []int{3, 9}

	if !reflect.DeepEqual(act, exp) {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}

func TestSumAllTails(t *testing.T) {
	act := SumAllTails([]int{1, 2}, []int{0, 9})
	exp := []int{2, 9}

	if !reflect.DeepEqual(act, exp) {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
