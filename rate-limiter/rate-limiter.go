package ratelimiter

import (
	"math"
	"sync"
	"time"
)

type RateLimiter struct {
	tokens         int
	maxTokens      int
	refillRate     int
	lastRefillTime time.Time
	mu             sync.Mutex
}

func NewRateLimiter(rate int, burst int) *RateLimiter {
	return &RateLimiter{
		tokens:         burst,
		maxTokens:      burst,
		refillRate:     rate,
		lastRefillTime: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	elapsed := time.Since(rl.lastRefillTime).Seconds()
	if elapsed > 0 {
		newTokens := int64(elapsed) * int64(rl.refillRate)
		rl.tokens = int(math.Min(float64(int64(rl.tokens)+newTokens), float64(rl.maxTokens)))
		rl.lastRefillTime = now
	}

	if rl.tokens > 0 {
		rl.tokens -= 1
		return true
	}

	return false
}
