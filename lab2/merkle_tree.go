package main

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
	return nil
}

// NewMerkleNode creates a new Merkle tree node
// implement
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	return nil
}

func (t *MerkleTree) SPVproof(index int) ([][]byte, error) {

	return nil, nil
}

func (t *MerkleTree) VerifyProof(index int, path [][]byte) (bool, error) {

	return true, nil
}
