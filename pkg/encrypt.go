package bitstream

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"io"
)

type EncryptedFile interface {
	io.WriterTo
	io.Reader
}

func Encrypt(pk *btcec.PrivateKey, paymentHash []byte, preimage []byte, inFile io.ReadSeeker, outFile io.Writer, chunkSize int64) error {
	encrypted, tree, err := EncryptFile(preimage, inFile, chunkSize)
	if err != nil {
		return err
	}

	encryptedRoot := tree.GetHash()
	message := append(encryptedRoot, paymentHash...)

	sig, err := schnorr.Sign(pk, message)
	if err != nil {
		return err
	}

	_, err = outFile.Write(sig.Serialize())
	if err != nil {
		return fmt.Errorf("failed to write schnorr sig: %w", err)
	}

	_, err = outFile.Write(paymentHash)
	if err != nil {
		return fmt.Errorf("failed to write payment hash: %w", err)
	}

	_, err = encrypted.WriteTo(outFile)
	if err != nil {
		return fmt.Errorf("failed to write encrypted file: %w", err)
	}

	return nil
}

func EncryptFile(preimage []byte, file io.ReadSeeker, chunkSize int64) (EncryptedFile, *MerkleTree, error) {
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
