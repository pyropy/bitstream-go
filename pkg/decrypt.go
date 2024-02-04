package bitstream

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
)

// Decrypt tries to decrypt file with given preimage and returns io.Reader of decrypted file.
//
// In case of an error, decrypt returns chunk index alongisde the error.
//
// If chunk index is not -1, then error ocurred while decrypting the chunk
// hence proof for given chunk should be generated.
func Decrypt(preimage []byte, file io.ReadSeeker, chunkSize int64) (io.Reader, int, error) {
	var out bytes.Buffer

	expectedHash := make([]byte, HashSize)
	encryptedChunk := make([]byte, chunkSize)

	var (
		i      uint64
		offset int64 = 96 // offset at sig len (64) + paymentHash len (32)
	)

	for {
		_, err := file.Seek(offset, 0)
		if err != nil {
			return nil, -1, err
		}

		// read expected hash from file
		n, err := file.Read(expectedHash)
		if n == 0 {
			if err == io.EOF {
				break
			}

			return nil, -1, err
		}

		// update offset for chunk read
		offset += chunkSize
		_, err = file.Seek(offset, 0)
		if err != nil {
			return nil, -1, err
		}

		// read chunk
		n, err = file.Read(encryptedChunk)
		if n == 0 {
			if err == io.EOF {
				break
			}

			return nil, -1, err
		}

		// update offset for next expected hash read
		offset += chunkSize

		// decrypt chunk, compute hash and if hashes match
		decryptedChunk := ChunkCipher(i, preimage, encryptedChunk)
		h := sha256.New()
		h.Write(decryptedChunk)
		hash := h.Sum(nil)

		if !bytes.Equal(expectedHash, hash) {
			return nil, int(i), fmt.Errorf("failed to decrypt file")
		}

		out.Write(decryptedChunk)
		clear(expectedHash)
		clear(encryptedChunk)

		i++
	}

	return nil, 0, nil
}
