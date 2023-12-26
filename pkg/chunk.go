package bitstream

import (
	"bytes"
	"io"
)

func ChunkFile(file io.ReadSeeker, chunkSize int64) (io.ReadCloser, error) {
	var (
		b      bytes.Buffer
		offset int64
	)

	for {
		_, err := file.Seek(offset, 0)
		if err != nil {
			return nil, err
		}

		chunk := make([]byte, chunkSize)
		n, err := file.Read(chunk)
		if n == 0 {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		b.Write(chunk)
		offset += chunkSize
	}

	return io.NopCloser(&b), nil
}
