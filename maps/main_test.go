package maps

import (
	"os"
	"testing"
)

var d Dictionary

func TestMain(m *testing.M) {
	d = Dictionary{}
	code := m.Run()
	d = nil
	os.Exit(code)
}

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
	d["test"] = "this is just a test"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act, err := d.Search(tc.key)
			assertError(t, err, tc.expErr)
			assertStrings(t, act, tc.exp)
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		exp    string
		expErr error
	}{
		{
			name:   "new key",
			key:    "newTest",
			exp:    "this is just a test",
			expErr: nil,
		},
		{
			name:   "existing key",
			key:    "test",
			exp:    "this is just a test",
			expErr: ErrKeyAlreadyExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := d.Add(tc.key, tc.exp)
			assertError(t, err, tc.expErr)
			act, err := d.Search(tc.key)
			assertStrings(t, act, tc.exp)
		})
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		exp    string
		expErr error
	}{
		{
			name:   "key exist",
			key:    "test",
			exp:    "this is an updated test",
			expErr: nil,
		},
		{
			name:   "key not exist",
			key:    "unknown",
			exp:    "",
			expErr: ErrUpdateKeyNotFound,
		},
	}

	d.Add("key exist", "this is just a test")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := d.Update(tc.key, tc.exp)
			assertError(t, err, tc.expErr)
			act, err := d.Search(tc.key)
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
