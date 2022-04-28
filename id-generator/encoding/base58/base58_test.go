package base58

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStdSource(t *testing.T) {
	const expect = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	assert.Equal(t, expect, StdSource())
}

func TestMustNewEncoder(t *testing.T) {
	assert.PanicsWithError(t, "base58: encoding source is not 58-bytes long", func() {
		MustNewEncoder("test")
	})
	assert.NotPanics(t, func() {
		MustNewEncoder("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	})
	assert.NotNil(t,
		MustNewEncoder("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"),
	)
}

func TestNewEncoder(t *testing.T) {
	enc, err := NewEncoder("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
	assert.NoError(t, err)
	assert.NotNil(t, enc)

	_, err = NewEncoder("test")
	assert.EqualError(t, err, "base58: encoding source is not 58-bytes long")
}

func TestStdEncoding_Encode(t *testing.T) {
	cases := map[int64]string{
		0:             "1",
		57:            "z",
		math.MaxInt8:  "3C",
		math.MaxInt16: "Ajx",
		math.MaxInt32: "4GmR58",
		math.MaxInt64: "NQm6nKp8qFC",
	}

	for k, v := range cases {
		e := StdEncoding.Encode(k)
		assert.Equal(t, v, e)
	}
}

func TestStdEncoding_Decode(t *testing.T) {
	cases := map[int64]string{
		0:             "1",
		57:            "z",
		math.MaxInt8:  "3C",
		math.MaxInt16: "Ajx",
		math.MaxInt32: "4GmR58",
		math.MaxInt64: "NQm6nKp8qFC",
	}

	for k, v := range cases {
		d, err := StdEncoding.Decode(v)
		if assert.Nil(t, err) {
			assert.Equal(t, k, d)
		}
	}
}

func TestEncoder_Encode(t *testing.T) {
	cases := map[int64]string{
		0:             "1",
		57:            "z",
		math.MaxInt8:  "3C",
		math.MaxInt16: "Ajx",
		math.MaxInt32: "4GmR58",
		math.MaxInt64: "NQm6nKp8qFC",
	}

	enc := MustNewEncoder(StdSource())
	assert.NotNil(t, enc)

	for k, v := range cases {
		e := enc.Encode(k)
		assert.Equal(t, v, e)
	}
}

func TestEncoder_Decode(t *testing.T) {
	cases := map[int64]string{
		0:             "1",
		57:            "z",
		math.MaxInt8:  "3C",
		math.MaxInt16: "Ajx",
		math.MaxInt32: "4GmR58",
		math.MaxInt64: "NQm6nKp8qFC",
	}

	enc := MustNewEncoder(StdSource())
	assert.NotNil(t, enc)

	for k, v := range cases {
		d, err := enc.Decode(v)
		if assert.Nil(t, err) {
			assert.Equal(t, k, d)
		}
	}
}

func BenchmarkEncoder_Encode(b *testing.B) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	enc := MustNewEncoder(StdSource())
	benchFunc := func(b *testing.B, id int64) {
		for i := 0; i < b.N; i++ {
			enc.Encode(id)
		}
	}

	for i := 0; i < 5; i++ {
		id := s.Int63()
		b.Run(fmt.Sprintf("%d", id), func(b *testing.B) {
			benchFunc(b, id)
		})
	}
}

func BenchmarkEncoder_Decode(b *testing.B) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	enc := MustNewEncoder(StdSource())
	benchFunc := func(b *testing.B, id int64, idStr string) {
		for i := 0; i < b.N; i++ {
			_, _ = enc.Decode(idStr)
		}
	}

	for i := 0; i < 5; i++ {
		id := s.Int63()
		idStr := enc.Encode(id)
		b.Run(fmt.Sprintf("%d - %s", id, idStr), func(b *testing.B) {
			benchFunc(b, id, idStr)
		})
	}
}
