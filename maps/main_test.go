package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		exp    string
		expErr error
	}{
		{
			name:   "key exist",
			key:    "test",
			exp:    "this is just a test",
			expErr: nil,
		},
		{
			name:   "key not exist",
			key:    "unknown",
			exp:    "",
			expErr: ErrKeyNotFound,
		},
	}

	d := Dictionary{"test": "this is just a test"}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := d.Search(tc.key)
			assertError(t, err, tc.expErr)
			assertStrings(t, act, tc.exp)
		})
	}
}

func assertStrings(t testing.TB, got, exp string) {
	t.Helper()

	if got != exp {
		t.Fatalf("expected %v, but got %v instead", exp, got)
	}
}

func assertError(t testing.TB, got, exp error) {
	t.Helper()
	if got != exp {
		t.Fatalf("expected error %v, but got %v instead", exp, got)
	}
}
