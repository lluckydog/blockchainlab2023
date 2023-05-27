package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// MerkleTree represent a Merkle tree
type MerkleTree struct {
	RootNode *MerkleNode
	Leaf     [][]byte
}

// MerkleNode represent a Merkle tree node
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
// implement
func NewMerkleTree(data [][]byte) *MerkleTree {
	// var node = MerkleNode{nil, nil, data[0]}
	// var mTree = MerkleTree{&node}
	MerkleNodes := make([]*MerkleNode, len(data))

	for i := 0; i < len(data); i++ {
		MerkleNodes[i] = NewMerkleNode(nil, nil, data[i])
	}

	for len(MerkleNodes) > 1 {
		var MerkleNodesNext []*MerkleNode
		for i := 0; i < len(MerkleNodes); i += 2 {
			left := MerkleNodes[i]
			right := left
			if len(MerkleNodes) != i+1 {
				right = MerkleNodes[i+1]
			}
			node := NewMerkleNode(left, right, nil)
			fmt.Printf("Merkle root %v, val %v\n", i, hex.EncodeToString(node.Data))
			MerkleNodesNext = append(MerkleNodesNext, node)
		}
		MerkleNodes = MerkleNodesNext
	}
	mTree := MerkleTree{
		RootNode: MerkleNodes[0],
		Leaf:     data,
	}

	return &mTree
}

// NewMerkleNode creates a new Merkle tree node
// implement
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	var node MerkleNode
	if left == nil && right == nil {
		tmp := sha256.Sum256(data)
		node = MerkleNode{
			Left:  left,
			Right: right,
			Data:  tmp[:],
		}
	} else {
		if right == nil {
			right = left
		}

		tmp := append(left.Data, right.Data...)

		d := sha256.Sum256(tmp)
		node = MerkleNode{
			Left:  left,
			Right: right,
			Data:  d[:],
		}
	}

	return &node
}

func (t *MerkleTree) SPVproof(index int) ([][]byte, error) {
	leafCount := len(t.Leaf)
	if index > leafCount {
		return nil, fmt.Errorf("No Such Leaf!")
	}
	h := 0
	cnt := leafCount
	for cnt > 1 {
		cnt = cnt/2 + cnt%2
		h++
	}
	var path [][]byte
	node := t.RootNode
	for i := h; i > 0; i-- {
		signal := 1 << (i - 1)

		if index&signal == 0 {
			path = append(path, node.Right.Data)
			node = node.Left
		} else {
			path = append(path, node.Left.Data)
			node = node.Right
		}
	}
	return path, nil
}

func (t *MerkleTree) VerifyProof(index int, path [][]byte) (bool, error) {
	if index >= len(t.Leaf) {
		return false, fmt.Errorf("No Such Leaf!")
	}
	data := sha256.Sum256(t.Leaf[index])
	signal := 1

	for i := len(path) - 1; i >= 0; i-- {

		if index&signal != 0 {
			tmp := append(path[i], data[:]...)
			data = sha256.Sum256(tmp)
		} else {
			tmp := append(data[:], path[i]...)
			data = sha256.Sum256(tmp)
		}
		signal = signal << 1
	}

	return bytes.Equal(data[:], t.RootNode.Data), nil
}
