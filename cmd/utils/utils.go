package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/pbkdf2"
)

func FileOrDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func EncryptFile(passphrase string, data []byte) []byte {
	key := deriveKey(passphrase)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("cypher Block err:")
		cobra.CheckErr(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("cypher GCM err:")
		cobra.CheckErr(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("nonce err:")
		cobra.CheckErr(err)
	}
	if err != nil {
		cobra.CheckErr(err)
	}
	cypherText := gcm.Seal(nonce, nonce, data, nil)

	return cypherText
}

func deriveKey(passphrase string) []byte {
	salt := []byte("randomsaltvalue")
	return pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha256.New)
}

func DecryptFile(passphrase string, cypherText []byte) ([]byte, error) {

	key := deriveKey(passphrase)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("cypher Block err:")
		cobra.CheckErr(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("cypher GCM err:")
		cobra.CheckErr(err)
	}

	nonce := cypherText[:gcm.NonceSize()]
	cypherText = cypherText[gcm.NonceSize():]

	plainFile, err := gcm.Open(nil, nonce, cypherText, nil)
	if err != nil {
		fmt.Println("decrypt file err:")
		cobra.CheckErr(err)
	}
	return plainFile, nil
}
