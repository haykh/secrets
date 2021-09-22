package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

/*
 * Encrypt takes in a key and plaintext, and then returns a hex
 * representation of the encrypted value.
 * Code based on standard library example at:
 * - https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
 */
func Encrypt(key, plaintext string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}

/*
 * Decrypt takes in a key and cipherhex (hex representation of
 * the ciphertext) and decrypts it.
 * Code based on standard library example at:
 * - https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
 */
func Decrypt(key, cipherhex string) (string, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(cipherhex)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream works in-place if both arguments are the same
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
