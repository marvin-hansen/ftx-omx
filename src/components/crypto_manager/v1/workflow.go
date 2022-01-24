// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package v1

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"ftx-omx/src/cfg"
	"ftx-omx/src/utils/crypto"
	"io"
)

// Symmetric AES encryption / decryption
// https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes

func (c *CryptoManager) GenerateKey(length uint) (key string, ok bool) {
	bytes := make([]byte, length) //generate a random 64 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		msg := "Failed to generate key to encrypt / decrypt"
		DbgPrint(msg)
		return "", false
	}
	key = hex.EncodeToString(bytes) //encode key in bytes to string to store in safe place!
	return key, true
}

func (c *CryptoManager) EncodeBase64(stringToEncode string) (encodedString string) {
	// https://gobyexample.com/base64-encoding
	encodedString = base64.StdEncoding.EncodeToString([]byte(stringToEncode))
	return encodedString
}

func (c *CryptoManager) DecodeBase64(stringToDecode string) (decodedString string) {
	bytes, _ := base64.StdEncoding.DecodeString(stringToDecode)
	decodedString = string(bytes)
	return decodedString
}

// https://gist.github.com/mickelsonm/e1bf365a149f3fe59119
func (c *CryptoManager) Encrypt(stringToEncrypt string) (encryptedString, msg string, ok bool) {

	// DbgPrint("Decode Master key")
	cryptKey := crypto.DecodeKey(cfg.GetMasterKey())

	// plaintext must be char array
	plainText := []byte(stringToEncrypt)

	block, err := aes.NewCipher([]byte(cryptKey))
	if err != nil {
		msg = "error creating cipher "
		return "", msg, false
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		msg = "error creating IV"
		return "", msg, false
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	// converts byte array to base64 encoded string
	encryptedString = base64.URLEncoding.EncodeToString(cipherText)
	msg = "ok"
	return encryptedString, msg, true
}

func (c *CryptoManager) Decrypt(encryptedString string) (decryptedString, msg string, ok bool) {
	//DbgPrint("Decode Master key")
	keyString := crypto.DecodeKey(cfg.GetMasterKey())

	// decode base64 back to byte array of encrypted data
	cipherText, err := base64.URLEncoding.DecodeString(encryptedString)
	if err != nil {
		msg = "error decoding base64 char "
		return "", msg, false
	}

	// creates new cipher
	block, err := aes.NewCipher([]byte(keyString))
	if err != nil {
		msg = "error creating cyper block "
		return "", msg, false
	}

	if len(cipherText) < aes.BlockSize {
		msg = "Ciphertext block size is too short! "
		return "", msg, false
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Stream which decrypts with cipher feedback mode using the given Block & IV
	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	// convert byte array back to string.
	decryptedString = string(cipherText)

	msg = "ok"
	return decryptedString, msg, true
}
