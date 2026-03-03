package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var key = "thisisasecretkeyuseaesgcmencrypt"
var nonce = make([]byte, 12)

func init() {
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
}

func encrypt(plainText string) (string, error) {
	// 创建一个新的 cipher 加密块
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create new cipher: %v", err)
	}

	// 创建一个新的 GCM 封装器
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	// 加密数据
	cipherText := aesgcm.Seal(nil, nonce, []byte(plainText), nil)

	// 转换为 十六进制 字符串
	hexCipherText := make([]byte, hex.EncodedLen(len(cipherText)))
	hex.Encode(hexCipherText, cipherText)

	return string(hexCipherText), nil
}

func decrypt(cipherText string) (string, error) {
	if len(cipherText) == 0 {
		return "", nil
	}

	// 从 十六进制 字符串转换为 []byte
	data, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("failed to hex DecodeString: %v", err)
	}

	// 创建一个新的 cipher 加密块
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("failed to create new cipher: %v", err)
	}

	// 创建一个新的 GCM 封装器
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	// 解密数据
	plainText, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %v", err)
	}

	return string(plainText), nil
}
