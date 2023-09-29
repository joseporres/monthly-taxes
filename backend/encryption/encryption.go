package encryption

import (
	"backend/global"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// const key=

func Encrypt(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(global.KEY))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(global.KEY))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func ToBase64(plaintext []byte) string {
	return base64.RawStdEncoding.EncodeToString(plaintext)
}

func FromBase64(ciphertext string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(ciphertext)
}
