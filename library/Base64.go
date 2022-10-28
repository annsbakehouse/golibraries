package library

import (
	"encoding/base64"
)

func Base64Encode(data string) string {
	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	return encodedString
}

func Base64Decode(data string) string {
	var decodedByte, _ = base64.StdEncoding.DecodeString(data)
	var decodedString = string(decodedByte)
	return decodedString
}
