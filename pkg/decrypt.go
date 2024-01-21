package bitstream

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
)

func Decrypt(preimage []byte, file io.ReadSeeker, chunkSize int64) (io.Reader, error) {
	var out bytes.Buffer

	sig := make([]byte, 64, 64)
	paymentHash := make([]byte, 32, 32)

	_, err := file.Read(sig)
	if err != nil {
		return nil, err
	}

	_, err = file.Read(paymentHash)
	if err != nil {
		return nil, err
	}

	var (
		i      uint64
		offset int64 = 96 // offset at sig + paymentHash bytes len
	)

	for {
		// read expected hash first
		_, err := file.Seek(offset, 0)
		if err != nil {
			return nil, err
		}

		expectedHash := make([]byte, chunkSize)
		n, err := file.Read(expectedHash)
		if n == 0 {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		offset += chunkSize // set offset to read chunk

		_, err = file.Seek(offset, 0)
		if err != nil {
			return nil, err
		}

		encryptedChunk := make([]byte, chunkSize)
		n, err = file.Read(encryptedChunk)
		if n == 0 {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		offset += chunkSize // TODO: use constant hash size to adjust next offset

		decryptedChunk := ChunkCipher(i, preimage, encryptedChunk)
		h := sha256.New()
		h.Write(decryptedChunk)
		hash := h.Sum(nil)

		if bytes.Compare(expectedHash, hash) != 0 {
			// TODO: Compute merkle proof
			return nil, fmt.Errorf("failed to decrypt file")
		}

		out.Write(decryptedChunk)
		i++
	}

	return nil, nil
}
