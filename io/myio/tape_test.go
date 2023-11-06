package myio

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestTape_Write(t *testing.T) {
	f, cleanFunc := createTmpFile(t, "123456")
	defer cleanFunc()

	exp := "abcd"

	tape := &tape{f}
	tape.Write([]byte(exp))
	f.Seek(0, 0)
	fileContent, _ := io.ReadAll(f)

	act := string(fileContent)
	assertStrings(t, act, exp)
}

func createTmpFile(t testing.TB, data string) (*os.File, func()) {
	t.Helper()

	f, err := os.CreateTemp("", "tst.data")
	if err != nil {
		log.Fatalf("error creating temp file: '%v'", err)
	}

	f.Write([]byte(data))
	clean := func() {
		f.Close()
		os.Remove(f.Name())
	}

	return f, clean
}

func assertStrings(t testing.TB, act, exp string) {
	t.Helper()
	if act != exp {
		t.Errorf("expected %v, but got %v instead", exp, act)
	}
}
