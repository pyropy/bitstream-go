package bitstream

import (
	"encoding/hex"
	"testing"
)

func hexMustDecodeString(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

func TestEncryptChunk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		index    int
		preimage []byte
		data     []byte
		expect   []byte
	}{
		{
			index:    0,
			preimage: hexMustDecodeString("f222b61f3508140048c55ad741b819b353629ac56b08437fa1c378b067d52f00"),
			data:     hexMustDecodeString("255044462d312e340a25c3a4c3bcc3b6c39f0a322030206f626a0a3c3c2f4c65"),
			expect:   hexMustDecodeString("a8a3e951d73f4b3b4d8d30eb538eed1e805b8f91082916c80a3419edaa5d85b1"),
		},
	}

	for _, tt := range tests {
		data32 := [32]byte{}
		copy(data32[:], tt.data)
		got, err := EncryptChunk(tt.index, tt.preimage, data32)
		if err != nil {
			t.Error(err)
		}

		if string(got[:]) != string(tt.expect) {
			t.Errorf("got %x, expected %s", got, tt.expect)
		}
	}

}
