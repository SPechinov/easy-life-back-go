package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "0123456789"

func GenerateRandomCode(n int) (string, error) {
	code := make([]byte, n)

	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code[i] = charset[num.Int64()]
	}

	return string(code), nil
}
