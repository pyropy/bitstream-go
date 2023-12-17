package bitstream

import (
	"bytes"
	"io"
	"os"
)

func EncryptFile(preimage []byte, file *os.File, chunkSize int64) (io.ReadCloser, *Node, error) {
	chunks, err := ChunkFile(file, chunkSize)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		err = chunks.Close()
		if err != nil {
			panic(err)
		}
	}()

	var (
		i     uint64
		b     bytes.Buffer
		nodes []*Node
	)

	for {
		chunk := make([]byte, 32)
		n, err := chunks.Read(chunk)
		if n == 0 {
			if err == io.EOF {
				break
			}

			return nil, nil, err
		}

		encryptedChunk := ChunkCipher(i, preimage, chunk)
		b.Write(encryptedChunk)

		encNode := NewNode(encryptedChunk)
		node := NewNode(chunk)
		nodes = append(nodes, encNode, node)

		i++
	}

	tree := NewTree(nodes)
	return io.NopCloser(&b), tree, nil
}
