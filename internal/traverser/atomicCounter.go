package traverser

import "sync/atomic"

type AtomicLimiter struct {
	counter int64
}

// Increment the counter atomically.
func (c *AtomicLimiter) IsExceed() bool {
	return atomic.LoadInt64(&c.counter) <= 0

}

// Decrement the counter atomically.
func (c *AtomicLimiter) Consume() {
	atomic.AddInt64(&c.counter, -1)

}

func NewLimiter() *AtomicLimiter {
	return &AtomicLimiter{counter: GRAPH_LIMIT}
}
