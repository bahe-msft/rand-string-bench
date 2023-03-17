package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
)

const letters = "0123456789abcdefghijklmnopqrstuvwxyz"

type cryptoStringGenerator struct{}

// NewCryptoStringGenerator creates a secure string generator with letter set 0123456789abcdefghijklmnopqrstuvwxyz.
func NewCryptoStringGenerator() *cryptoStringGenerator {
	return &cryptoStringGenerator{}
}

func (r *cryptoStringGenerator) Generate(length int) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("invalid length")
	}

	// ref: https://github.com/kubernetes/cluster-bootstrap/blob/v0.26.2/token/util/helpers.go
	// len("0123456789abcdefghijklmnopqrstuvwxyz") = 36 which doesn't evenly divide
	// the possible values of a byte: 256 mod 36 = 4. Discard any random bytes we
	// read that are >= 252 so the bytes we evenly divide the character set.
	const maxByteValue = 252

	var (
		b     byte
		err   error
		token = make([]byte, length)
	)

	reader := bufio.NewReaderSize(rand.Reader, length*2)
	for i := range token {
		for {
			if b, err = reader.ReadByte(); err != nil {
				return "", err
			}
			if b < maxByteValue {
				break
			}
		}

		token[i] = letters[int(b)%len(letters)]
	}

	return string(token), nil
}
