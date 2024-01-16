package bitstream

import (
	"bytes"
	"io"
)

//func Encrypt(paymentHash []byte, preimage []byte, inFile io.ReadSeeker, outFile io.Writer, chunkSize int64) error {
//	encrypted, tree, err := EncryptFile(preimage, inFile, chunkSize)
//	if err != nil {
//		return err
//	}
//
//	encryptedRoot := tree.GetHash()
//	// TODO: Compute Schnorr signature from message
//	message := append(encryptedRoot, paymentHash...)
//
//	outFile.Write(paymentHash)
//
//}

func EncryptFile(preimage []byte, file io.ReadSeeker, chunkSize int64) (io.Reader, *MerkleTree, error) {
	chunks, err := ChunkFile(file, chunkSize)
	if err != nil {
		return nil, nil, err
	}

	var (
		i     uint64
		b     bytes.Buffer
		nodes []*Node
	)

	for {
		chunk := make([]byte, chunkSize)
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
	return &b, tree, nil
}
