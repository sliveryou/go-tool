package pkcs

import (
	"bytes"
	"fmt"
)

// Reference:
// https://en.wikipedia.org/wiki/Padding_(cryptography)#PKCS5
// https://blog.csdn.net/xz_studying/article/details/94229023

// PKCS5Padding pkcs5 padding method
func PKCS5Padding(cipherText []byte) []byte {
	padding := 8 - len(cipherText)%8
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

// PKCS5Trimming pkcs5 trimming method
func PKCS5Trimming(encrypt []byte) ([]byte, error) {
	padding := encrypt[len(encrypt)-1]
	end := len(encrypt) - int(padding)
	if end <= 0 {
		return nil, fmt.Errorf("PKCS5Trimming err, err encrypt data")
	}

	return encrypt[:end], nil
}
