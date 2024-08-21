package helpers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func ToSha1(text, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(text))
	hex.EncodeToString(h.Sum(nil))
}
