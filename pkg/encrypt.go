package bitstream

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

type EncryptedFile interface {
	io.WriterTo
	io.Reader
}

func Encrypt(pk *btcec.PrivateKey, paymentHash, preimage []byte, inFile io.ReadSeeker, outFile io.Writer, chunkSize int64) error {
	encrypted, tree, err := EncryptFile(preimage, inFile, chunkSize)
	if err != nil {
		return err
	}

	encryptedRoot := tree.GetHash()
	// original implementation signs encrypted root + payment hash
	sig, err := schnorr.Sign(pk, encryptedRoot)
	if err != nil {
		return fmt.Errorf("failed to generate schorr sig: %w", err)
	}

	_, err = outFile.Write(sig.Serialize())
	if err != nil {
		return fmt.Errorf("failed to write schnorr sig: %w", err)
	}

	var paymentHashFixed [32]byte
	copy(paymentHashFixed[:], paymentHash)
	_, err = outFile.Write(paymentHashFixed[:])
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

		h := sha256.New()
		h.Write(chunk)
		hash := h.Sum(nil)

		encryptedChunk := ChunkCipher(i, preimage, chunk)

		b.Write(hash)
		b.Write(encryptedChunk)

		encNode := NewNode(encryptedChunk)
		node := NewNode(chunk)
		nodes = append(nodes, encNode, node)

		i++
	}

	tree := NewTree(nodes)
	return &b, tree, nil
}
