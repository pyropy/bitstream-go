package bitstream

import (
	"crypto/sha256"
	"encoding/binary"
)

// ChunkCipher encrypts or decrypts a byte array (chunk) of data using the preimage and index
func ChunkCipher(index uint64, preimage []byte, data []byte) []byte {
	result := make([]byte, len(data))

	// add 1 to index
	index += 1

	// convert index to bytes (original implementation uses btc specific encoding)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// hash preimage + index
	hasher := sha256.New()
	hasher.Write(preimage)
	hasher.Write(indexBytes)

	hash := hasher.Sum(nil)

	// NOTE: Original implementation uses 32 byte chunks, here we use arbitrary chunk length
	// xor hash with data
	for i := 0; i < len(data); i++ {
		j := i % 32
		result[i] = hash[j] ^ data[i]
	}

	return result
}
