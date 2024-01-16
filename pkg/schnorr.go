package bitstream

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
)

func Sign(message []byte, priv *ecdsa.PrivateKey) ([]byte, error) {
	h := sha256.Sum256(message)
	fmt.Println(h)

	return nil, nil
}
