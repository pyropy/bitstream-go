package main

import (
	"flag"
	"fmt"
	bitstream "github.com/pyropy/bitstream-go/pkg"
	"io"
	"log"
	"os"
)

var (
	inPath    string
	outPath   string
	preimage  string
	chunkSize int64
)

func init() {
	flag.StringVar(&inPath, "in", "", "File to encrypt/decrypt")
	flag.StringVar(&outPath, "out", "", "Path to output encrypted/decrypted file")
	flag.StringVar(&preimage, "preimage", "", "Preimage to use for encryption/decryption")
	flag.Int64Var(&chunkSize, "size", 32, "Size of chunks to encrypt/decrypt")
}

// TODO: Implement encrypt and decrypt cli commands
func main() {
	flag.Parse()

	err := encrypt(inPath, outPath)
	if err != nil {
		panic(err)
	}
}

func encrypt(inPath, outPath string) error {
	f, err := os.Open(inPath)
	if err != nil {
		return err
	}

	encryptedFile, tree, err := bitstream.EncryptFile([]byte(preimage), f, chunkSize)
	if err != nil {
		return err
	}

	l := fmt.Sprintf("Merkle tree hash 0x%x", tree.GetHash())
	log.Println(l)

	b, err := io.ReadAll(encryptedFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(outPath, b, 0644)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
