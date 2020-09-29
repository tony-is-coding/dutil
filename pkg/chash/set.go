package chash

import (
	"errors"
	"fmt"
)

type Set interface {
	Get(i uint32) string
	Add(node string) error
	Remove(node string) error
	Len() uint32
}

type NodeSet struct {
	nodeList []string
	nodeMap  map[string]uint32
	nodeNum  uint32
}

func (n *NodeSet) Get(i uint32) string {
	return n.nodeList[i]
}

func (n *NodeSet) Add(node string) error {
	if _, ok := n.nodeMap[node]; ok {
		return errors.New("duplicate register node")
	}
	n.nodeList = append(n.nodeList, node)
	n.nodeMap[node] = n.nodeNum
	n.nodeNum += 1
	return nil
}

func (n *NodeSet) Remove(node string) error {
	index, ok := n.nodeMap[node]
	if !ok {
		return errors.New(fmt.Sprintf("node: < %s > does not exist", node))
	}
	n.nodeList = append(n.nodeList[:index], n.nodeList[index+1:]...)
	return nil
}

func (n *NodeSet) Len() uint32 {
	return n.nodeNum
}
