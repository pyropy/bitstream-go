package bitstream

import "io"

type FraudProof struct {
	Preimage []byte
}

func GenerateFraudProof(preimage []byte, file io.Reader) (io.ReadCloser, error) {
	return nil, nil
}
