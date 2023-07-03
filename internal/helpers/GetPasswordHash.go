package helpers

import (
	"crypto/sha1"
	"encoding/hex"
)

func GetPasswordHash(password string) string {
	hasher := sha1.New()

	if _, err := hasher.Write([]byte(password)); err == nil {
		hash := hex.EncodeToString(hasher.Sum(nil))
		return hash
	}

	return ""

}
