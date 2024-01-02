package chapter5

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(s string) string {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
