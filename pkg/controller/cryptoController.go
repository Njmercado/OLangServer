package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

// Hash works to create a new Hash data.
// data string --> data to be hashed
// return string
func Hash(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// createNonce works to create some nonce array based on random data.
// size []byte --> nonce size
// return ( []byte, []byte )
func createNonce(size int) []byte {
	nonce := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	return nonce
}

// getNonce works to split cypher data into nonce and real cypher data
// data []byte
// nonceSize int
// return ( []byte, []byte )
func getNonce(data []byte, nonceSize int) ([]byte, []byte) {
	return data[:nonceSize], data[nonceSize:]
}

// Encrypt works to handle data Encryp action. Encrypt(data []byte, key string) []byte
func Encrypt(data []byte, key string) []byte {
	hash := Hash(key)
	block, _ := aes.NewCipher([]byte(hash))

	counter, _ := cipher.NewGCM(block)
	nonce := createNonce(counter.NonceSize())

	finalText := counter.Seal(nonce, nonce, data, nil)
	// cypherText := append(finalText, nonce...)

	return finalText
}

// Decrypt works to handle data decryp action. Decrypt(data []byte, key string) string
func Decrypt(data []byte, key string) []byte {
	hash := Hash(key)
	block, err := aes.NewCipher([]byte(hash))
	if err != nil {
		log.Print("This data could not been decrypted")
		return nil
	}

	counter, _ := cipher.NewGCM(block)
	nonce, cypherData := getNonce(data, counter.NonceSize())
	decryptedData, err := counter.Open(nil, nonce, cypherData, nil)
	if err != nil {
		log.Print("This data could not been decrypted")
		return nil
	}

	return decryptedData
}
