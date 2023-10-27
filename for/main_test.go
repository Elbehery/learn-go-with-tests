package main

import "testing"

func TestRepeat(t *testing.T) {

	tcs := []struct {
		name        string
		inputString string
		inputRepeat int
		exp         string
	}{
		{
			name:        "empty",
			inputString: "",
			inputRepeat: 10,
			exp:         "",
		},
		{
			name:        "10",
			inputString: "a",
			inputRepeat: 10,
			exp:         "aaaaaaaaaa",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := Repeat(tc.inputString, tc.inputRepeat)
			if act != tc.exp {
				t.Errorf("expected %v, but got %v instead", tc.exp, act)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
