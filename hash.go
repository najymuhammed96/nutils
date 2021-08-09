package nutils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return strings.ToLower(hex.EncodeToString(hash[:]))
}
