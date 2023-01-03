package util

import (
	"encoding/base64"
)

func Base64Encrypt(data string) (result string) {
	result = base64.StdEncoding.EncodeToString([]byte(data))
	return
}

func Base64Decode(data string) (result string) {
	decode, _ := base64.StdEncoding.DecodeString(data)
	result = string(decode)
	return
}
