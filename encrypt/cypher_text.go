// Package encrypt contains functions to cypher and de-cypher data
package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Reference: https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/

// createHash takes a passphrase or any string,
// hash it, then return the hash as a hexadecimal value.
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt encrypts the text with the provided passphrase value
func Encrypt(text string, passphrase string) (error, string) {
	data := []byte(text)
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err, ""
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err, ""
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return nil, string(ciphertext)
}

// Decrypt decrypts the text with the provided passphrase value
func Decrypt(text string, passphrase string) (error, string) {
	data := []byte(text)
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return err, ""
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err, ""
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err, ""
	}
	return nil, string(plaintext)
}
