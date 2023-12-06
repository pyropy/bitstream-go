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
		index    uint64
		preimage []byte
		data     []byte
		expect   []byte
	}{
		{
			index:    0,
			preimage: hexMustDecodeString("f222b61f3508140048c55ad741b819b353629ac56b08437fa1c378b067d52f00"),
			data:     hexMustDecodeString("255044462d312e340a25c3a4c3bcc3b6c39f0a322030206f626a0a3c3c2f4c65"),
			expect:   hexMustDecodeString("d74415cbc1a4ce52f8d1ca29a255289532bb7a18bba30a25ce624a01da55368f"),
		},
	}

	for _, tt := range tests {
		got := ChunkCipher(tt.index, tt.preimage, tt.data)

		if string(got) != string(tt.expect) {
			t.Errorf("got %x, expected %s", got, tt.expect)
		}
	}

}

func TestDecryptChunk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		index    uint64
		preimage []byte
		data     []byte
		expect   []byte
	}{
		{
			index:    0,
			preimage: hexMustDecodeString("f222b61f3508140048c55ad741b819b353629ac56b08437fa1c378b067d52f00"),
			data:     hexMustDecodeString("d74415cbc1a4ce52f8d1ca29a255289532bb7a18bba30a25ce624a01da55368f"),
			expect:   hexMustDecodeString("255044462d312e340a25c3a4c3bcc3b6c39f0a322030206f626a0a3c3c2f4c65"),
		},
	}

	for _, tt := range tests {
		got := ChunkCipher(tt.index, tt.preimage, tt.data)

		if string(got) != string(tt.expect) {
			t.Errorf("got %x, expected %s", got, tt.expect)
		}
	}

}
