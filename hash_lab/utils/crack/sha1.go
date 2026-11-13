package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1HASH(password string) string {
	data := []byte(password)

	hash := sha1.Sum(data)

	return hex.EncodeToString(hash[:])
}