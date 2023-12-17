package bitstream

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		leaves [][]byte
		expect string
	}{
		{
			leaves: [][]byte{},
			expect: "0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			leaves: [][]byte{
				[]byte("hello"),
			},
			expect: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			leaves: [][]byte{
				[]byte("world"),
			},
			expect: "486ea46224d1bb4fb680f34f7c9ad96a8f24ec88be73ea8e5a6c65260e9cb8a7",
		},
		{
			leaves: [][]byte{
				[]byte("hello"),
				[]byte("world"),
			},
			expect: "7305db9b2abccd706c256db3d97e5ff48d677cfe4d3a5904afb7da0e3950e1e2",
		},
		{
			leaves: [][]byte{
				[]byte("hello"),
				[]byte("world"),
				[]byte("world"),
			},
			expect: "060f35066ef7e1db64e560fcdcf9f55955fb942671e3ecc79d3fb8adefb79bcb",
		},
		{
			leaves: [][]byte{
				[]byte("hello"),
				[]byte("world"),
				[]byte("hello"),
				[]byte("world"),
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
