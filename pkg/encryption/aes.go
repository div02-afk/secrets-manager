package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

type AESProvider struct {
}

func (a *AESProvider) Encrypt(key []byte, payload []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := aead.Seal(nil, nonce, payload, nil)

	return append(nonce, cipherText...), nil
}

func (a *AESProvider) Decrypt(key []byte, payload []byte) ([]byte, error) {

	log.Println("Key", string(key))
	log.Println("Secret", string(payload))

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Couldnt create new cipher")
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("Couldnt create new gcm aead")
		return nil, err
	}

	nonceSize := aead.NonceSize()
	if len(payload) < nonceSize {
		log.Printf("payload len < nonceSize")
		return nil, fmt.Errorf("invalid encrypted DEK")
	}

	nonce, ciphertext := payload[:nonceSize], payload[nonceSize:]

	plainText, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Printf("decryption failed")
		return nil, err
	}

	return plainText, nil
}
