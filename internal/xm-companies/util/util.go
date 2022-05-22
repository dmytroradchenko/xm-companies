package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func CheckHash(str, hash string) bool {
	strHash := Hash(str)
	return strHash == hash
}
