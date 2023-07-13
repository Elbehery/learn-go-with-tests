package concurrency

import (
	"reflect"
	"testing"
)

func myMockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestMyCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := MyCheckWebsites(myMockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, but got %v instead", want, got)
	}

}
