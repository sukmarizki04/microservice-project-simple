package messageque

import (
	"time"
)

type TokenBucket struct {
	capacity     int
	tokens       int
	refillRate   int
	lastRefilled time.Time
}

func (tb *TokenBucket) Allow() bool {
	tb.refillRateTokens()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func (tb *TokenBucket) refillTokens() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefilled)
	tokenToAdd := int(elapsed.Seconds() * float64(tb.refillRate))
	tbtokens := min(tb.capacity, tb.tokens+tokenToAdd)
	tb.lastRefilled = norm
}
