package main

import (
	"crypto/aes"
	"crypto/cipher"
)

// Decrypt the data
func decrypt(data string, key []byte) string {
	cipherData := []byte(data)

	block, err := aes.NewCipher(key)
	err_panic(err)

	if len(cipherData) < aes.BlockSize {
		return "TOO_SHORT_DATA"
	}

	iv := cipherData[:aes.BlockSize]
	cipherData = cipherData[aes.BlockSize:]

	strean := cipher.NewCFBDecrypter(block, iv)
	strean.XORKeyStream(cipherData, cipherData)

	return string(cipherData)
}
