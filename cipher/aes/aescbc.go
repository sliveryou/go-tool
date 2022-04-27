package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/sliveryou/go-tool/cipher/pkcs"
)

// There aescbc add pkcs7Padding to be same as php's aescbc

const (
	// AesCbc128KeyLen key len: 16
	AesCbc128KeyLen = 16
	// AesCbc192KeyLen key len: 24
	AesCbc192KeyLen = 24
	// AesCbc256KeyLen key len: 32
	AesCbc256KeyLen = 32
	// IvLen iv len: 16
	IvLen = 16
)

// AesCbc the base aes cbc structure
type AesCbc struct {
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

// NewAesCbc new aes cbc cipher
// aescbc support key len 16 24 32 match aescbc-128 aescbc-192 aescbc-256
// iv len must be 16
func NewAesCbc(key, iv []byte) (*AesCbc, error) {
	k := len(key)
	switch k {
	default:
		return nil, fmt.Errorf("key len must be 16,24,32 your key is %d", k)
	case AesCbc128KeyLen, AesCbc192KeyLen, AesCbc256KeyLen:
		break
	}

	if len(iv) != IvLen {
		return nil, fmt.Errorf("iv len must be 16 your iv is %d", len(iv))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return &AesCbc{
		key:   key,
		iv:    iv,
		block: block,
	}, nil
}

// Encrypt the aes cbc encrypt method
func (a *AesCbc) Encrypt(src []byte) ([]byte, error) {
	return a.encrypt(src)
}

// Decrypt the aes cbc decrypt method
func (a *AesCbc) Decrypt(src []byte) ([]byte, error) {
	return a.decrypt(src)
}

func (a *AesCbc) decrypt(src []byte) ([]byte, error) {
	decrypter := cipher.NewCBCDecrypter(a.block, a.iv)

	plainText := make([]byte, len(src))
	decrypter.CryptBlocks(plainText, src)

	return pkcs.PKCS7Trimming(plainText)
}

func (a *AesCbc) encrypt(src []byte) ([]byte, error) {
	paddingText := pkcs.PKCS7Padding(src, aes.BlockSize)

	encrypter := cipher.NewCBCEncrypter(a.block, a.iv)
	cipherText := make([]byte, len(paddingText))
	encrypter.CryptBlocks(cipherText, paddingText)

	return cipherText, nil
}

// The follow functions are used for easy to call test
// or diffirent key to cipher

// AesCbcEncrypt the aes cbc encrypt method
func AesCbcEncrypt(key, iv, src []byte) ([]byte, error) {
	a, err := NewAesCbc(key, iv)
	if err != nil {
		return nil, err
	}

	return a.encrypt(src)
}

// AesCbcDecrypt the ase cbc decrypt method
func AesCbcDecrypt(key, iv, src []byte) ([]byte, error) {
	a, err := NewAesCbc(key, iv)
	if err != nil {
		return nil, err
	}

	return a.decrypt(src)
}

// AesCbcEncryptHex return hex result
func AesCbcEncryptHex(key, iv, src []byte) (string, error) {
	a, err := NewAesCbc(key, iv)
	if err != nil {
		return "", err
	}
	dst, err := a.encrypt(src)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(dst), nil
}

// AesCbcDecryptHex decrypt hex msg
func AesCbcDecryptHex(key, iv []byte, msg string) ([]byte, error) {
	a, err := NewAesCbc(key, iv)
	if err != nil {
		return nil, err
	}

	data, err := hex.DecodeString(msg)
	if err != nil {
		return nil, err
	}

	return a.decrypt(data)
}

// AesCbcEncryptBase64 return base64 result
func AesCbcEncryptBase64(key, iv, src []byte) (string, error) {
	a, err := NewAesCbc(key, iv)
	if err != nil {
		return "", err
	}
	dst, err := a.encrypt(src)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(dst), nil
}

// AesCbcDecryptBase64 decrypt base64 msg
func AesCbcDecryptBase64(key, iv []byte, msg string) ([]byte, error) {
	a, err := NewAesCbc(key, iv)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return nil, err
	}

	return a.decrypt(data)
}
