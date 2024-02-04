package bitstream

import (
	"crypto/sha256"
)

type Hasher interface {
	GetHash() []byte
}

type MerkleTree struct {
	Root *Node
}

func (t *MerkleTree) GetHash() []byte {
	return t.Root.GetHash()
}

type Node struct {
	Hash  []byte
	Left  *Node
	Right *Node
}

type MerkleProof struct {
	Root []byte
	Path [][]byte
}

func (n *Node) GetHash() []byte {
	if len(n.Hash) > 0 {
		return n.Hash
	}

	h := sha256.New()
	h.Write(n.Left.GetHash())

	if n.Right != nil {
		h.Write(n.Right.GetHash())
	} else {
		h.Write(n.Left.GetHash())
	}

	return h.Sum(nil)
}

func NewNode(data []byte) *Node {
	h := sha256.New()
	h.Write(data)

	return &Node{Hash: h.Sum(nil)}
}

func NewNodeFromHash(hash []byte) *Node {
	return &Node{Hash: hash}
}

func NewTree(leaves []*Node) *MerkleTree {
	if len(leaves) == 0 {
		return &MerkleTree{
			Root: &Node{
				Hash: make([]byte, 32),
			},
		}
	}

	for len(leaves) > 1 {
		var nextLevel []*Node

		if len(leaves)%2 != 0 {
			leaf := leaves[len(leaves)-1]
			leaves = append(leaves, leaf)
		}

		// go through all leaves on current level in pairs
		for i := 0; i < len(leaves); i += 2 {
			left := leaves[i]
			right := leaves[i+1]

			// create parent node
			parent := newParentNode(left, right)

			// append parent to next level
			nextLevel = append(nextLevel, parent)
		}

		leaves = nextLevel
	}

	return &MerkleTree{Root: leaves[0]}
}

func GenerateMerkleProof(leaves []*Node, index int) *MerkleProof {
	var path [][]byte

	for len(leaves) > 1 {
		var nextLevel []*Node

		if len(leaves)%2 != 0 {
			leaf := leaves[len(leaves)-1]
			leaves = append(leaves, leaf)
		}

		// go through all leaves on current level in pairs
		for i := 0; i < len(leaves); i += 2 {
			left := leaves[i]
			right := leaves[i+1]

			// create parent node
			parent := newParentNode(left, right)

			// append parent to next level
			nextLevel = append(nextLevel, parent)

			if i <= index && index < i+2 {
				if index%2 == 0 {
					path = append(path, right.GetHash())
				} else {
					path = append(path, left.GetHash())
				}

				index = i / 2
			}
		}

		leaves = nextLevel
	}

	return &MerkleProof{Path: path, Root: leaves[0].GetHash()}
}

func newParentNode(left, right *Node) *Node {
	// generate parent hash
	h := sha256.New()
	h.Write(left.Hash)
	h.Write(right.Hash)

	return &Node{Hash: h.Sum(nil), Left: left, Right: right}
}
