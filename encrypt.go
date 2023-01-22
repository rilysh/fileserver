package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Encrypt the data with a
func encrypt(data string, key []byte) string {
	block, err := aes.NewCipher(key)
	err_panic(err)

	cipherData := make([]byte, aes.BlockSize+len(data))
	iv := cipherData[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, iv)
	err_panic(err)

	strean := cipher.NewCFBEncrypter(block, iv)
	strean.XORKeyStream(cipherData[aes.BlockSize:], []byte(data))

	return string(cipherData)
}
