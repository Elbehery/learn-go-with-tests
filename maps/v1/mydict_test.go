package main

import "testing"

func TestDictionary_Search(t *testing.T) {
	scenarios := []struct {
		name string
		key  string
		want string
		err  error
	}{
		{"know word",
			"test",
			"this is a test",
			nil},
		{"no match",
			"notest",
			"",
			ErrNoMatch},
	}

	dict := NewMyDict()
	dict["test"] = "this is a test"

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			got, err := dict.Search(scenario.key)
			if scenario.err != nil {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if scenario.err != err {
					assertResult(t, err.Error(), scenario.err.Error())
				}
			}
			assertResult(t, scenario.want, got)
		})
	}
}

func assertResult(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, but expected %v instead", got, want)
	}
}
