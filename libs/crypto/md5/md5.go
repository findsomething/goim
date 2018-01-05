package md5

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}
