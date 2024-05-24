package passwordmanager

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var key []byte

type Config struct {
	EncryptionKey string `json:"encryption_key"`
}

func Encrypt(password string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(password), nil)
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(hashedPassword string) (string, error) {
	data, err := hex.DecodeString(hashedPassword)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func loadKeyFromConfig() ([]byte, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %v", err)
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("could not decode config JSON: %v", err)
	}

	key, err := hex.DecodeString(config.EncryptionKey)
	if err != nil {
		return nil, fmt.Errorf("invalid key format: %v", err)
	}

	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key length: %v", len(key))
	}

	return key, nil
}

func init() {
	var err error
	key, err = loadKeyFromConfig()
	if err != nil {
		log.Fatalf("Error loading key: %v", err)
	}
}
