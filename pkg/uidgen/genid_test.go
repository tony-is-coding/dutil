package uidgen

import (
	"testing"
)

func TestSnowFlake(t *testing.T) {
	SnowFlake()
}

func BenchmarkSnowFlake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SnowFlake()
	}
}
