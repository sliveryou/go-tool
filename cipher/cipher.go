package cipher

// The  interface is used for usual cipher
// Now it support aescbc
// Plan to support aesecb rsa
import (
	"github.com/sliveryou/go-tool/v2/cipher/aes"
)

// If you want convert encrypted content to appoint format
// You can use convertUtil gitlab.33.cn\util\go-kit\convert package
// Such as hex convertUtil.BytesEncodeHex(bytes)

// The cipher will deal with some diffirent between php/nodejs cipher
// Such as aescbc add pkcs7Padding to be same as php's aescbc

// Cipher the cipher interface
type Cipher interface {
	// Encrypt the encrypt method
	Encrypt(src []byte) ([]byte, error)
	// Decrypt the decrypt method
	Decrypt(src []byte) ([]byte, error)
}

// The follow is a aescbc demo

// import cipherUtil "gitlab.33.cn/utils/go-kit/cipher"
//  cipher, err := cipherUtil.NewAesCbc(key, iv)
//	if err != nil {
//		//deal error
//	}
//  //encrypt
//  dst, err := cipher.Encrypt(src)
//	if err != nil {
//		//deal error
//	}
//
//	//decrypt
//	dst, err := cipher.Decrypt(dst)
//	if err != nil {
//		//deal error
//	}

// NewAesCbc support aescbc-128  aescbc-192 aescbc-256
// match key len     16          24         32
func NewAesCbc(key, iv string) (Cipher, error) {
	return aes.NewAesCbc([]byte(key), []byte(iv))
}

// MustNewAesCbc NewAesCbc err will panic , be careful
func MustNewAesCbc(key, iv string) Cipher {
	c, err := aes.NewAesCbc([]byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}

	return c
}
