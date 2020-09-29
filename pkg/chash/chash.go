package chash

import (
	"hash"
)

// this package implement the consistent hash

type CHash interface {
	AddNode(node string) error
	RemoveNode(node string) error
	GetNode([]byte) string
	GetNodeNum() uint32
}

type murmur3CHash struct {
	h hash.Hash32
	s Set
}

func (m *murmur3CHash) AddNode(node string) error { return m.s.Add(node) }

func (m *murmur3CHash) RemoveNode(node string) error { return m.s.Remove(node) }

func (m *murmur3CHash) GetNode(p []byte) string {
	// todo: hash32 compute
	index := m.h.Sum32() % m.s.Len()
	return m.s.Get(index)

}

func (m *murmur3CHash) GetNodeNum() uint32 { return m.s.Len() }

func NewMurmur3CHash() *murmur3CHash {
	return nil
}
