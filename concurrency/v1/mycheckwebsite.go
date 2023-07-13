package concurrency

import "net/http"

func MyCheckWebsite(url string) bool {
	resp, err := http.Head(url)
	if err != nil {
		return false
	}
	return resp.StatusCode == http.StatusOK
}
