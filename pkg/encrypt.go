package bitstream

import (
	"crypto/sha256"
	"io"
)

func Encrypt(preimage []byte, file io.Reader) (io.ReadCloser, error) {
	return nil, nil
}

// EncryptChunk encrypts a byte array (chunk) of data using the preimage and index
func EncryptChunk(index int, preimage []byte, data [32]byte) ([32]byte, error) {
	var result [32]byte

	// add 1 to index
	index += 1
	// convert index to bytes
	indexBytes := toBytes(int(index))

	// hash preimage + index
	hasher := sha256.New()
	hasher.Write(preimage)
	hasher.Write(indexBytes)

	hash := hasher.Sum(nil)

	for i := 0; i < 32; i++ {
		result[i] = hash[i] ^ data[i]
	}

	return result, nil
}

// TODO: Find a better way to do this
func toBytes(n int) []byte {
	if n == 0 {
		return []byte{}
	}

	var result []byte
	neg := n < 0
	absN := abs(n)

	for absN > 0 {
		byteValue := byte(absN & 0xff)
		absN >>= 8
		result = append(result, byteValue)
	}

	if result[len(result)-1]&0x80 != 0 {
		if neg {
			result = append(result, 0x80)
		} else {
			result = append(result, 0)
		}
	} else if neg {
		result[len(result)-1] |= 0x80
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
