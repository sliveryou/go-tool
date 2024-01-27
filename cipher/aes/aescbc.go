package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/sliveryou/go-tool/v2/cipher/pkcs"
)

// There aescbc add pkcs7Padding to be same as php's aescbc

const (
	// Cbc128KeyLen key len: 16
	Cbc128KeyLen = 16
	// Cbc192KeyLen key len: 24
	Cbc192KeyLen = 24
	// Cbc256KeyLen key len: 32
	Cbc256KeyLen = 32
	// IvLen iv len: 16
	IvLen = 16
)

// Cbc the base aes cbc structure
type Cbc struct {
	key   []byte
	iv    []byte
	block cipher.Block
	// Originally want to multiple call BlockMode.CryptBlocks to
	// reduce memory apply and release
	// But BlockMode's iv will be change and it not support reset iv
	// So we should only new BlockMode when we encrypt/decrypt
	// encrypter cipher.BlockMode
	// decrypter cipher.BlockMode
}

// NewCbc new aes cbc cipher
// aescbc support key len 16 24 32 match aescbc-128 aescbc-192 aescbc-256
// iv len must be 16
func NewCbc(key, iv []byte) (*Cbc, error) {
	k := len(key)
	switch k {
	default:
		return nil, fmt.Errorf("key len must be 16,24,32 your key is %d", k)
	case Cbc128KeyLen, Cbc192KeyLen, Cbc256KeyLen:
	}

	if len(iv) != IvLen {
		return nil, fmt.Errorf("iv len must be 16 your iv is %d", len(iv))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &Cbc{
		key:   key,
		iv:    iv,
		block: block,
	}, nil
}

// Encrypt the aes cbc encrypt method
func (c *Cbc) Encrypt(src []byte) ([]byte, error) {
	paddingText := pkcs.PKCS7Padding(src, aes.BlockSize)

	encrypter := cipher.NewCBCEncrypter(c.block, c.iv)
	cipherText := make([]byte, len(paddingText))
	encrypter.CryptBlocks(cipherText, paddingText)

	return cipherText, nil
}

// Decrypt the aes cbc decrypt method
func (c *Cbc) Decrypt(src []byte) ([]byte, error) {
	decrypter := cipher.NewCBCDecrypter(c.block, c.iv)

	plainText := make([]byte, len(src))
	decrypter.CryptBlocks(plainText, src)

	return pkcs.PKCS7Trimming(plainText)
}

// The follow functions are used for easy to call test
// or different key to cipher

// CbcEncrypt the aes cbc encrypt method
func CbcEncrypt(key, iv, src []byte) ([]byte, error) {
	a, err := NewCbc(key, iv)
	if err != nil {
		return nil, err
	}

	return a.Encrypt(src)
}

// CbcDecrypt the ase cbc decrypt method
func CbcDecrypt(key, iv, src []byte) ([]byte, error) {
	a, err := NewCbc(key, iv)
	if err != nil {
		return nil, err
	}

	return a.Decrypt(src)
}

// CbcEncryptHex return hex result
func CbcEncryptHex(key, iv, src []byte) (string, error) {
	a, err := NewCbc(key, iv)
	if err != nil {
		return "", err
	}
	dst, err := a.Encrypt(src)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(dst), nil
}

// CbcDecryptHex decrypt hex msg
func CbcDecryptHex(key, iv []byte, msg string) ([]byte, error) {
	a, err := NewCbc(key, iv)
	if err != nil {
		return nil, err
	}

	data, err := hex.DecodeString(msg)
	if err != nil {
		return nil, err
	}

	return a.Decrypt(data)
}

// CbcEncryptBase64 return base64 result
func CbcEncryptBase64(key, iv, src []byte) (string, error) {
	a, err := NewCbc(key, iv)
	if err != nil {
		return "", err
	}
	dst, err := a.Encrypt(src)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(dst), nil
}

// CbcDecryptBase64 decrypt base64 msg
func CbcDecryptBase64(key, iv []byte, msg string) ([]byte, error) {
	a, err := NewCbc(key, iv)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return nil, err
	}

	return a.Decrypt(data)
}
