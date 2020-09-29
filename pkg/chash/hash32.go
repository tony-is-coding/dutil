package chash

import (
	"math/bits"
	"unsafe"
)

const (
	C321 uint32 = 0xcc9e2d51
	C322 uint32 = 0x1b873593
	Seed uint32 = 0x1234abcd
)

type murmurHash struct {
}

func (m *murmurHash) Write(p []byte) (int, error) { return 0, nil }
func (m *murmurHash) Sum(b []byte) []byte         { return nil }
func (m *murmurHash) Reset()                      {}
func (m *murmurHash) Size() int                   { return 0 }
func (m *murmurHash) BlockSize() int              { return 0 }
func (m *murmurHash) Sum32() uint32               { return 0 }

func sum32(data []byte) uint32 {
	h1 := Seed
	nblocks := len(data) / 4
	var p uintptr
	if len(data) > 0 {
		p = uintptr(unsafe.Pointer(&data[0]))
	}
	p1 := p + uintptr(4*nblocks)
	for ; p < p1; p += 4 {
		k1 := *(*uint32)(unsafe.Pointer(p))

		k1 *= C321
		k1 = bits.RotateLeft32(k1, 15)
		k1 *= C322

		h1 ^= k1
		h1 = bits.RotateLeft32(h1, 13)
		h1 = h1*4 + h1 + 0xe6546b64
	}

	tail := data[nblocks*4:]

	var k1 uint32
	switch len(tail) & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= C321
		k1 = bits.RotateLeft32(k1, 15)
		k1 *= C322
		h1 ^= k1
	}

	h1 ^= uint32(len(data))

	h1 ^= h1 >> 16
	h1 *= 0x85ebca6b
	h1 ^= h1 >> 13
	h1 *= 0xc2b2ae35
	h1 ^= h1 >> 16

	return h1
}

/*
 	simple and may not useful
 */
func hashFunc(digest []byte, nTime int) uint32 {
	rv := (uint32(digest[3+nTime]&0xFF) << 24) | ((uint32)(digest[2+nTime*4]&0xFF) << 16) |
		((uint32)(digest[1+nTime*4]&0xFF) << 8) | (uint32(digest[0+nTime*4] & 0xFF))
	return rv
}
