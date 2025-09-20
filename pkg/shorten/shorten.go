package shorten

import (
	"crypto/sha256"
	"encoding/base64"
)

const BASE62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeBase62(num int) string {
	if num == 0 {
		return string(BASE62[0])
	}

	result := []byte{}

	for num > 0 {
		rem := num % 62
		num = num / 62
		result = append(result, BASE62[rem])
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func HashURL(url string, length int) string {
	sha := sha256.Sum256([]byte(url))
	b64 := base64.URLEncoding.EncodeToString(sha[:])

	if length > len(b64) {
		return b64
	}

	return b64[:length]
}
