package sync

import "sync"

type Counter struct {
	mux *sync.Mutex
	val int
}

func NewCounter() *Counter {
	return &Counter{
		mux: &sync.Mutex{},
		val: 0,
	}
}

func (c *Counter) Inc() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}
