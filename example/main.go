package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	bitstream "github.com/pyropy/bitstream-go/pkg"
)

var (
	action      string
	inPath      string
	outPath     string
	preimage    string
	paymentHash string
	chunkSize   int64
	pk          *btcec.PrivateKey
)

func init() {
	flag.StringVar(&action, "action", "encrypt", "encrypt or decrypt")
	flag.StringVar(&inPath, "in", "", "File to encrypt/decrypt")
	flag.StringVar(&outPath, "out", "", "Path to output encrypted/decrypted file")
	flag.StringVar(&preimage, "preimage", "", "Preimage to use for encryption/decryption")
	flag.StringVar(&paymentHash, "hash", "", "Payment hash")
	flag.Int64Var(&chunkSize, "size", 32, "Size of chunks to encrypt/decrypt")

	pkBytes, err := hex.DecodeString("80faa7d1c150a903a4028bf87ca8800aff507b24df90e8434bfa1f34d639c053")
	if err != nil {
		panic(err)
	}

	pk, _ = btcec.PrivKeyFromBytes(pkBytes)
}

// TODO: Implement encrypt and decrypt cli commands
func main() {
	flag.Parse()

	var err error

	switch action {
	case "encrypt":
		err = encrypt(inPath, outPath, preimage, chunkSize)
	case "decrypt":
		err = decrypt(inPath, outPath, preimage, chunkSize)
	}

	if err != nil {
		panic(err)
	}
}

func encrypt(inPath, outPath, preimage string, chunkSize int64) error {
	f, err := os.Open(inPath)
	if err != nil {
		return err
	}

	defer f.Close()

	out, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	preimageBytes, err := hex.DecodeString(preimage)
	if err != nil {
		panic(err)
	}

	paymentHashBytes, err := hex.DecodeString(paymentHash)
	if err != nil {
		return err
	}

	err = bitstream.Encrypt(pk, paymentHashBytes, preimageBytes, f, out, chunkSize)
	if err != nil {
		return err
	}

	return nil
}

func decrypt(inPath, outPath, preimage string, chunkSize int64) error {
	f, err := os.Open(inPath)
	if err != nil {
		return err
	}

	defer f.Close()

	preimageBytes, err := hex.DecodeString(preimage)
	if err != nil {
		return err
	}

	decryptedFile, index, err := bitstream.Decrypt(preimageBytes, f, chunkSize)
	if err != nil {
		if index != -1 {
			proof, err := bitstream.GenerateProof(f, chunkSize, 2*index)
			if err != nil {
				return err
			}

			root := fmt.Sprintf("%x", proof.MerkleProof.Root)
			log.Println("Generated proof", "proof root", root)
		}

		return err
	}

	decrypted, err := io.ReadAll(decryptedFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(outPath, decrypted, 0644)
	if err != nil {
		return err
	}

	return nil
}
