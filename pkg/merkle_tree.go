package bitstream

import "crypto/sha256"

type Hasher interface {
	GetHash() []byte
}

type Node struct {
	Hash  []byte
	Left  *Node
	Right *Node
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

func NewTree(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return &Node{Hash: make([]byte, 32)}
	}

	for len(nodes) > 1 {
		var newLevel []*Node

		for j := 0; j < len(nodes); j += 2 {
			left := nodes[j]
			right := nodes[j]

			if j+1 < len(nodes) {
				right = nodes[j+1]
			}

			h := sha256.New()
			h.Write(left.Hash)
			h.Write(right.Hash)

			newLevel = append(newLevel, &Node{Hash: h.Sum(nil), Left: left, Right: right})
		}

		nodes = newLevel
	}

	return nodes[0]
}
