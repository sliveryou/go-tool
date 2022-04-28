package base62

import (
	"encoding/binary"
	"errors"
)

var (
	// StdEncoding base62 standard encoder.
	StdEncoding = MustNewEncoder(StdSource())

	// stdSource base62 standard source string.
	stdSource = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// Encoder base62 encoder.
type Encoder struct {
	encode    [62]byte
	decodeMap [256]int
}

// StdSource returns base62 standard source string.
func StdSource() string {
	return stdSource
}

// MustNewEncoder must new a base62 encoder.
func MustNewEncoder(source string) *Encoder {
	enc, err := NewEncoder(source)
	if err != nil {
		panic(err)
	}
	return enc
}

// NewEncoder new a base62 encoder.
func NewEncoder(source string) (*Encoder, error) {
	if len(source) != 62 {
		return nil, errors.New("base62: encoding source is not 62-bytes long")
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

// Encode base62 encodes a int64 id.
func (enc *Encoder) Encode(id int64) string {
	if id == 0 {
		return string(enc.encode[:1])
	}

	bin := make([]byte, 0, binary.MaxVarintLen64)
	for id > 0 {
		bin = append(bin, enc.encode[id%62])
		id /= 62
	}

	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	return string(bin)
}

// Decode base62 decodes a string id.
func (enc *Encoder) Decode(id string) (int64, error) {
	if id == "" {
		return 0, errors.New("base62: decoding id should not be empty")
	}

	var n int64
	for i := range id {
		u := enc.decodeMap[id[i]]
		if u < 0 {
			return 0, errors.New("base62: invalid decoding character - " + string(id[i]))
		}
		n = n*62 + int64(u)
	}

	return n, nil
}
