// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package crypto

import (
	"encoding/base64"
)

func DecodeKey(stringToDecode string) (decodedString string) {
	return decodeBase64(stringToDecode)
}

func decodeBase64(stringToDecode string) (decodedString string) {
	bytes, err := base64.StdEncoding.DecodeString(stringToDecode)
	if err != nil {
		return err.Error()
	} else {
		decodedString = string(bytes)
		return decodedString
	}
}
