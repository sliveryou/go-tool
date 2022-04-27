package pkcs

import (
	"bytes"
	"fmt"
)

// Reference:
// wiki https://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS7
// https://blog.csdn.net/xz_studying/article/details/94229023

// PKCS7Padding pkcs7 padding method
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

// PKCS7Trimming pkcs7 trimming method
func PKCS7Trimming(encrypt []byte) ([]byte, error) {
	padding := encrypt[len(encrypt)-1]
	end := len(encrypt) - int(padding)
	if end <= 0 {
		return nil, fmt.Errorf("PKCS7Trimming err, err encrypt data")
	}

	return encrypt[:end], nil
}
