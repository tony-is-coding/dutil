package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	l       sync.Locker
	lastTry int64
	limit   int64
	size    int64
	rate    time.Duration
}

func (t *TokenBucket) Try() error {
	now := int64(time.Now().Nanosecond())
	fillPieces := (now - t.lastTry) / int64(t.rate)
	if (t.size + fillPieces - 1) < 0 {
		// 令牌桶中令牌不够一次取
		return PoolFullErr
	}
	// 先定令牌桶上限
	t.l.Lock()
	t.size = t.size + fillPieces - 1
	t.lastTry = now
	t.l.Unlock()
	return nil
}

func (t *TokenBucket) Serve() {}

func NewTokenBucket(lc sync.Locker, lm int64, r time.Duration) *TokenBucket {
	return &TokenBucket{
		l:       lc,
		lastTry: int64(time.Now().Nanosecond()),
		limit:   lm,
		size:    0,
		rate:    r,
	}
}
