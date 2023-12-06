package bitstream

import (
	"crypto/sha256"
	"encoding/binary"
	"io"
)

func Decrypt(preimage []byte, file io.Reader) (io.ReadCloser, error) {

	return nil, nil
}

func DecryptChunk(index uint64, preimage []byte, expectedHash []byte, data []byte) ([]byte, error) {
	result := make([]byte, len(data))

	// add 1 to index
	index += 1
	// convert index to bytes
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

	// hasher.Reset()
	// hasher.Write(result)
	// resultHash := hasher.Sum(nil)
	// if string(resultHash) != string(expectedHash) {
	// 	return nil, fmt.Errorf("hash mismatch: %x != %x", resultHash, expectedHash)
	// }

	return result, nil
}
