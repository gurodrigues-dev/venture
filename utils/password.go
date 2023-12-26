package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func checkPasswordHash(password, hash string) bool {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil)) == hash
}
