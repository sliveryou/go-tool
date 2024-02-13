package cipher

// The interface is used for usual cipher.
// Now it support aescbc.
// Plan to support aesecb rsa.
import (
	"github.com/sliveryou/go-tool/v2/cipher/aes"
)

var _ Cipher = (*aes.Cbc)(nil)

// The cipher will deal with some diffirent between php/nodejs cipher
// Such as aescbc add pkcs7Padding to be same as php's aescbc

// Cipher the cipher interface.
type Cipher interface {
	// Encrypt the encrypt method
	Encrypt(src []byte) ([]byte, error)
	// Decrypt the decrypt method
	Decrypt(src []byte) ([]byte, error)
}

// The follow is a aescbc demo

// import "github.com/sliveryou/go-tool/cipher"
//
// 	cr, err := cipher.NewCbc(key, iv)
// 	if err != nil {
// 		// deal error
// 	}
// 	// encrypt
// 	dst, err := cr.Encrypt(src)
// 	if err != nil {
// 		// deal error
// 	}
//
// 	// decrypt
// 	dst, err := cr.Decrypt(dst)
// 	if err != nil {
// 		// deal error
// 	}

// NewAesCbc support aescbc-128  aescbc-192 aescbc-256,
// match key len     16          24         32.
func NewAesCbc(key, iv string) (*aes.Cbc, error) {
	return aes.NewCbc([]byte(key), []byte(iv))
}

// MustNewAesCbc NewAesCbc err will panic, be careful.
func MustNewAesCbc(key, iv string) *aes.Cbc {
	c, err := aes.NewCbc([]byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}

	return c
}
