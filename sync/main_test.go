package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter.val, 3)
	})

	t.Run("test concurrently", func(t *testing.T) {
		wanted := 10000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wanted)

		for i := 0; i < wanted; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()
		assertCounter(t, counter.val, wanted)
	})
}

func assertCounter(t testing.TB, exp, act int) {
	t.Helper()
	if exp != act {
		t.Fatalf("expected %v, but got %v instead", exp, act)
	}
}
