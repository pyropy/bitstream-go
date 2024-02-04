package bitstream

import "io"

const (
	SignatureSize = 64
	HashSize      = 32
)

type Proof struct {
	Sig         []byte
	PaymentHash []byte
	MerkleProof *MerkleProof
}

func GenerateProof(file io.ReadSeeker, chunkSize int64, chunkIndex int) (*Proof, error) {
	sig := make([]byte, SignatureSize)
	paymentHash := make([]byte, HashSize)
	hash := make([]byte, chunkSize)

	// read schnorr signature
	_, err := file.Read(sig)
	if err != nil {
		return nil, err
	}

	// read payment hash
	_, err = file.Read(paymentHash)
	if err != nil {
		return nil, err
	}

	var (
		leaves []*Node
		offset int64 = 96 // offset at sig len (64) + paymentHash len (32)
	)

	// read all leaves from  file
	for {
		_, err := file.Seek(offset, 0)
		if err != nil {
			return nil, err
		}

		// read expected hash from file
		n, err := file.Read(hash)
		if n == 0 {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		// update offset for next chunk hash read
		offset += chunkSize * 2
		node := NewNodeFromHash(hash)
		leaves = append(leaves, node)

		clear(hash)
	}

	// generate merkleProof from leaves
	merkleProof := GenerateMerkleProof(leaves, chunkIndex)

	return &Proof{
		Sig:         sig,
		PaymentHash: paymentHash,
		MerkleProof: merkleProof,
	}, nil
}
