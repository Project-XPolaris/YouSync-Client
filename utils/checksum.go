package utils

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Checksum(data []byte) string {
	hash := sha256.New()
	sum := hash.Sum(data)
	return fmt.Sprintf("%x", sum)
}
