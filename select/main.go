package _select

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aStart := time.Now()
	http.Get(a)
	aDuration := time.Since(aStart)

	bStart := time.Now()
	http.Get(b)
	bDuration := time.Since(bStart)

	if aDuration < bDuration {
		return a
	}
	return b
}
