package bitstream

import "crypto/sha256"

type Hasher interface {
	Hash() []byte
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{
		Root: &Node{},
	}
}

func (t *Tree) Insert(n *Node) {
	t.Root.Leaves = append(t.Root.Leaves, n)
}

func (t *Tree) Hash() []byte {
	return t.Root.Hash()
}

type Node struct {
	Data   []byte
	Leaves []*Node
}

func (n *Node) Hash() []byte {
	h := sha256.New()
	if n.IsLeaf() {
		h.Write(n.Data)
	} else {
		for _, l := range n.Leaves {
			h.Write(l.Hash())
		}
	}

	return h.Sum(nil)
}

func (n *Node) IsLeaf() bool {
	return len(n.Leaves) == 0
}
