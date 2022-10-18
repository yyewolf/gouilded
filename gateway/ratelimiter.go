package gateway

import (
	"context"
	"time"

	"github.com/sasha-s/go-csync"
)

type Ratelimiter struct {
	config *RatelimiterConfig

	reset   time.Time
	remains int

	mu csync.Mutex
}

func NewRatelimiter(config *RatelimiterConfig) *Ratelimiter {
	return &Ratelimiter{
		config: config,
		reset:  time.Now(),
	}
}

func (r *Ratelimiter) Close(ctx context.Context) {
	_ = r.mu.CLock(ctx)
}

func (r *Ratelimiter) Wait(ctx context.Context) error {
	// Locks the rate limiter for 60/config.maxRequestsPerSecond seconds
	if err := r.mu.CLock(ctx); err != nil {
		return err
	}

	now := time.Now()

	var until time.Time

	// We can reset the rate limiter
	if r.remains == 0 && r.reset.After(now) {
		until = r.reset
	}

	// We can't reset the rate limiter
	if until.After(now) {
		if deadline, ok := ctx.Deadline(); ok && until.After(deadline) {
			return context.DeadlineExceeded
		}

		select {
		case <-ctx.Done():
			r.Unlock()
			return ctx.Err()
		case <-time.After(until.Sub(now)):
		}
	}
	return nil

}

func (r *Ratelimiter) Unlock() {
	// Unlocks the rate limiter
	now := time.Now()
	if r.reset.Before(now) {
		r.reset = now.Add(time.Minute)
		r.remains = r.config.maxRequestsPerMinute
	}
	r.mu.Unlock()
}
