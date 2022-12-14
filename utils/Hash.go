package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512(text string) string {
	algorithm := sha512.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
