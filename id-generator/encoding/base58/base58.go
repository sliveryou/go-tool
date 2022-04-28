package base58

import (
	"encoding/binary"
	"errors"
)

var (
	// StdEncoding base58 standard encoder.
	StdEncoding = MustNewEncoder(StdSource())

	// stdSource base58 standard source string.
	stdSource = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

// Encoder base58 encoder.
type Encoder struct {
	encode    [58]byte
	decodeMap [256]int
}

// StdSource returns base58 standard source string.
func StdSource() string {
	return stdSource
}

// MustNewEncoder must new a base58 encoder.
func MustNewEncoder(source string) *Encoder {
	enc, err := NewEncoder(source)
	if err != nil {
		panic(err)
	}
	return enc
}

// NewEncoder new a base58 encoder.
func NewEncoder(source string) (*Encoder, error) {
	if len(source) != 58 {
		return nil, errors.New("base58: encoding source is not 58-bytes long")
	}

	enc := new(Encoder)

	for i := range enc.decodeMap {
		enc.decodeMap[i] = -1
	}

	for i := range source {
		enc.encode[i] = source[i]
		enc.decodeMap[enc.encode[i]] = i
	}

	return enc, nil
}

// Encode base58 encodes a int64 id.
func (enc *Encoder) Encode(id int64) string {
	if id == 0 {
		return string(enc.encode[:1])
	}

	bin := make([]byte, 0, binary.MaxVarintLen64)
	for id > 0 {
		bin = append(bin, enc.encode[id%58])
		id /= 58
	}

	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	return string(bin)
}

// Decode base58 decodes a string id.
func (enc *Encoder) Decode(id string) (int64, error) {
	if id == "" {
		return 0, errors.New("base58: decoding id should not be empty")
	}

	var n int64
	for i := range id {
		u := enc.decodeMap[id[i]]
		if u < 0 {
			return 0, errors.New("base58: invalid decoding character - " + string(id[i]))
		}
		n = n*58 + int64(u)
	}

	return n, nil
}
