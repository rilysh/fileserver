package main

// The authentication key for all encrypted files.
// IMPORTANT: SHARING THIS KEY, ANYONE CAN DECRYPT ALL HOSTED
// ENCRYPTED FILES. Always use your own randomly generated key.
// Don't use the default key which I specified here.
// Key length must be either 13, 24, or 32 bytes
// For more info about key length, see here: https://go.dev/src/crypto/aes/cipher.go, at line 31
func authKey() string {
	return "d6XwcN*@^^&&jdMm27X#tUJJ4@Mwc8Xm"
}
