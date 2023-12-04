package bitstream

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	t.Parallel()

	left := &Node{
		Data: []byte("hello"),
	}

	right := &Node{
		Data: []byte("world"),
	}

	tree := Tree{
		Root: &Node{
			Leaves: []*Node{left, right},
		},
	}

	nodes := []struct {
		node   Hasher
		expect string
	}{
		{
			node:   left,
			expect: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			node:   right,
			expect: "486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7",
		},
		{
			node:   tree.Root,
			expect: "7305db9b2abccd706c256db3d97e5ff48d677cfe4d3a5904afb7da0e3950e1e2",
		},
	}

	for _, tt := range nodes {
		got := tt.node.Hash()
		if fmt.Sprintf("%x", got) != tt.expect {
			t.Errorf("got %x, expected %s", got, tt.expect)
		}
	}
}
