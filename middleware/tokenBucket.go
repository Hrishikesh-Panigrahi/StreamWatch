package middleware

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity     int       // Max tokens in the bucket
	tokens       int       // Current tokens
	refillRate   time.Duration // Interval to refill 1 token
	lastRefill   time.Time // Last refill timestamp
	mu           sync.Mutex
}

func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Refill bucket over time
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)
	tb.lastRefill = now
	tb.tokens += int(elapsed / tb.refillRate)

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
}

// TryConsume checks if we can consume a token
func (tb *TokenBucket) TryConsume() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}
