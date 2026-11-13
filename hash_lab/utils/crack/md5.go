package crack

import (
    "crypto/md5"
    "encoding/hex"
)

func MD5Hash(password string) string {
    data := []byte(password)

    hash := md5.Sum(data)

    return hex.EncodeToString(hash[:])
}
