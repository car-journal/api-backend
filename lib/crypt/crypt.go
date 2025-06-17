package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type CryptInterface interface {
	// With B64 Std
	EncodeBase64Std(value string) string
	DecodeBase64Std(value string) (string, error)

	// With Secret
	DecryptString(value string) (response string, err error)
	EncryptString(value string) (response string, err error)
}

type crypt struct {
	secretKey string
}

func NewCrypt(secretKey string) CryptInterface {
	return &crypt{
		secretKey: secretKey,
	}
}

func (c crypt) EncodeBase64Std(value string) string {
	return c.encodeBase64Std([]byte(value))
}

func (c crypt) DecodeBase64Std(value string) (string, error) {
	res, err := c.decodeBase64Std(value)
	if nil != err {
		return "", err
	}
	return string(res), nil
}

func (c crypt) DecryptString(value string) (response string, err error) {
	key, err := c.decodeBase64Std(c.key())
	if err != nil {
		return "", errors.New("seems like you provide a key in base64 format, but it's not valid")
	}

	chiperBytes, err := c.decodeBase64Std(value)
	if err != nil {
		return "", errors.New("ciphertext value must in base64 format")
	}

	var payload struct {
		IV    string
		Value string
		Mac   string
	}
	err = json.Unmarshal(chiperBytes, &payload)
	if err != nil {
		return "", errors.New("ciphertext value must be valid")
	}

	iv, err := c.decodeBase64Std(payload.IV)
	if err != nil {
		return "", errors.New("iv in payload must be valid base64 format")
	}

	encryptedText, err := c.decodeBase64Std(payload.Value)
	if err != nil {
		return "", errors.New("encrypted text must be valid base64 format")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptedText, encryptedText)

	encryptedTextLength := len(encryptedText)
	unpadding := int(encryptedText[encryptedTextLength-1])
	encryptedText = encryptedText[:(encryptedTextLength - unpadding)]

	response = string(encryptedText)

	return response, nil
}

func (c crypt) EncryptString(value string) (response string, err error) {
	plainBytes := []byte(value)
	plainBytesLength := aes.BlockSize - len(plainBytes)%aes.BlockSize
	padding := []byte{byte(plainBytesLength)}
	padding = bytes.Repeat(padding, plainBytesLength)
	plainBytes = append(plainBytes, padding...)

	key, err := c.decodeBase64Std(c.key())
	if err != nil {
		return "", errors.New("seems like you provide a key in base64 format, but it's not valid")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(plainBytes, plainBytes)

	payload := make(map[string]string)
	payload["iv"] = c.encodeBase64Std(iv)
	payload["value"] = c.encodeBase64Std(plainBytes)

	h := hmac.New(sha256.New, []byte(key))
	_, err = io.WriteString(h, payload["iv"]+payload["value"])
	if err != nil {
		return "", err
	}

	payload["mac"] = fmt.Sprintf("%x", h.Sum(nil))

	chiperBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	response = c.encodeBase64Std(chiperBytes)

	return response, nil
}

func (c crypt) key() (response string) {
	if strings.HasPrefix(c.secretKey, "base64:") {
		return string(c.secretKey[7:])
	}
	return c.secretKey
}

func (c crypt) encodeBase64Std(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}

func (c crypt) decodeBase64Std(value string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(value)
}
