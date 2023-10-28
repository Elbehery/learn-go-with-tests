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

	tcs := []struct {
		name  string
		input [][]int
		exp   []int
	}{
		{
			"all slices",
			[][]int{{1, 2}, {0, 9}},
			[]int{2, 9},
		},
		{
			"empty slice",
			[][]int{{}, []int{3, 4, 5}},
			[]int{0, 9},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := SumAllTails(tc.input...)
			if !reflect.DeepEqual(act, tc.exp) {
				t.Errorf("expected %v, but got %v instead", tc.exp, act)
			}
		})
	}
}
