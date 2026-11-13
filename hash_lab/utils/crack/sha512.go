package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

func SHA512HASH(password string) string {
	data := []byte(password)

	hash := sha512.Sum512(data)

	return hex.EncodeToString(hash[:])
}