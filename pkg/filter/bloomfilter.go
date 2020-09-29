package filter

import "dutil/pkg/murmur3"

const BitSize = 1 << 10 << 10

type BloomFilter interface {
	Contains([]byte) bool
	Put([]byte)
	Delete([]byte)
}

type bloomFilter64 struct {
	BitArr [BitSize]byte
	Seeds  []uint32
}

func NewBloomFilter64(seeds []uint32) *bloomFilter64 {
	// TODO: seed 去重
	return &bloomFilter64{
		BitArr: [BitSize]byte{},
		Seeds:  seeds,
	}
}

func (b *bloomFilter64) Contains(data []byte) bool {
	for _, s := range b.Seeds {
		idx := murmur3.Sum64WithSeed(data, s)
		if b.BitArr[idx] == 0 {
			return false
		}
	}
	return true
}

func (b *bloomFilter64) Put(data []byte) {
	for _, s := range b.Seeds {
		idx := murmur3.Sum64WithSeed(data, s)
		if b.BitArr[idx] == 0 {
			b.BitArr[idx] = 1
		}
	}
}

func (b *bloomFilter64) Delete(data []byte) {
	for _, s := range b.Seeds {
		idx := murmur3.Sum64WithSeed(data, s)
		if b.BitArr[idx] == 0 {
			return
		}
		b.BitArr[idx] = 0
	}
}
