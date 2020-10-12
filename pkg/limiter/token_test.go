package limiter

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	tb := NewTokenBucket(&sync.Mutex{}, 1000, time.Nanosecond*100)
	for i := 0; i < 1000; i++ {
		if err := tb.Try(); err != nil {
			fmt.Println("try failed")
			continue
		}
		fmt.Println("try success")
	}
}
