// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package crypto_manager

import v1 "web_socket/src/components/crypto_manager/v1"

type CryptoManager interface {
	GenerateKey(length uint) (key string, ok bool)
	EncodeBase64(stringToEncode string) (encodedString string)
	DecodeBase64(stringToDecode string) (encodedString string)
	Encrypt(stringToEncrypt string) (encryptedString, msg string, ok bool)
	Decrypt(encryptedString string) (decryptedString, msg string, ok bool)
}

func NewNewCryptoManager() (cryptoManager CryptoManager) {
	cryptoManager = v1.NewCryptoManager()
	return cryptoManager
}
