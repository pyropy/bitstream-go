package bitstream

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		leaves []*Node
		expect string
	}{
		{
			leaves: []*Node{},
			expect: "0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			leaves: []*Node{
				NewNode([]byte("hello")),
			},
			expect: "1d25c19a1a3fb65c78d018561057362916c14bfd36b75aa8cb0f4d696293b183",
		},
		{
			leaves: []*Node{
				NewNode([]byte("world")),
			},
			expect: "a9e049081e9eaeef2d30029031314c888955d7181a032c29b11633ac11d6076c",
		},
		{
			leaves: []*Node{
				NewNode([]byte("hello")),
				NewNode([]byte("world")),
			},
			expect: "7305db9b2abccd706c256db3d97e5ff48d677cfe4d3a5904afb7da0e3950e1e2",
		},
		{
			leaves: []*Node{
				NewNode([]byte("hello")),
				NewNode([]byte("world")),
				NewNode([]byte("world")),
			},
			expect: "060f35066ef7e1db64e560fcdcf9f55955fb942671e3ecc79d3fb8adefb79bcb",
		},
		{
			leaves: []*Node{
				NewNode([]byte("hello")),
				NewNode([]byte("world")),
				NewNode([]byte("hello")),
				NewNode([]byte("world")),
			},
			expect: "f1159918525d2c95e89b7bcd62bc3da6297a49402e85d26e21bf08e5eace81c0",
		},
	}

	for _, tt := range tests {
		tree := NewTree(tt.leaves)
		got := tree.GetHash()
		if fmt.Sprintf("%x", got) != tt.expect {
			t.Errorf("got %x, expected %s", got, tt.expect)
		}
	}
}

func TestGenerateMerkleProof(t *testing.T) {
	t.Parallel()

	tests := []struct {
		index      int
		leaves     []*Node
		expectRoot []byte
		expectPath [][]byte
	}{
		{
			index: 2,
			leaves: []*Node{
				NewNode([]byte("hello")),
				NewNode([]byte("world")),
				NewNode([]byte("hello")),
				NewNode([]byte("world")),
			},
			expectRoot: hexMustDecodeString("f1159918525d2c95e89b7bcd62bc3da6297a49402e85d26e21bf08e5eace81c0"),
			expectPath: [][]byte{
				hexMustDecodeString("486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7"),
				hexMustDecodeString("7305db9b2abccd706c256db3d97e5ff48d677cfe4d3a5904afb7da0e3950e1e2"),
			},
		},
	}

	for _, tt := range tests {
		got := GenerateMerkleProof(tt.leaves, tt.index)

		if !bytes.Equal(tt.expectRoot, got.Root) {
			t.Errorf("Expected root %x, got %x", tt.expectRoot, got.Root)
		}

		if len(tt.expectPath) != len(got.Path) {
			t.Errorf("expected len %d got %d", len(tt.expectPath), len(got.Path))
			continue
		}

		for i := 0; i < len(tt.expectPath); i++ {
			x := tt.expectPath[i]
			y := got.Path[i]

			if !bytes.Equal(x, y) {
				t.Errorf("Expected %x, got %x", x, y)
			}
		}
	}
}
