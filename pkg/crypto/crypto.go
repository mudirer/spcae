package crypto

import (
	"crypto/sha256"
	"fmt"
	"io"
)

// Sha256Hex hash function
func Sha256Hex(data string) string {
	h256 := sha256.New()
	io.WriteString(h256, data)
	return fmt.Sprintf("%x", h256.Sum(nil))
}


