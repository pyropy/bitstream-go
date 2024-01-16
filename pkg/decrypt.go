package bitstream

import (
	"fmt"
	"io"
)

func Decrypt(preimage []byte, file io.ReadSeeker, chunkSize int64) (io.ReadCloser, error) {
	chunks, err := ChunkFile(file, chunkSize)
	if err != nil {
		return nil, err
	}

	fmt.Println(chunks)

	return nil, nil
}
