package base62

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStdSource(t *testing.T) {
	const expect = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	assert.Equal(t, expect, StdSource())
}

func TestMustNewEncoder(t *testing.T) {
	assert.PanicsWithError(t, "base62: encoding source is not 62-bytes long", func() {
		MustNewEncoder("test")
	})
	assert.NotPanics(t, func() {
		MustNewEncoder("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	})
	assert.NotNil(t,
		MustNewEncoder("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	)
}

func TestNewEncoder(t *testing.T) {
	enc, err := NewEncoder("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	assert.NoError(t, err)
	assert.NotNil(t, enc)

	_, err = NewEncoder("test")
	assert.EqualError(t, err, "base62: encoding source is not 62-bytes long")
}

func TestStdEncoding_Encode(t *testing.T) {
	cases := map[int64]string{
		0:             "0",
		57:            "v",
		math.MaxInt8:  "23",
		math.MaxInt16: "8WV",
		math.MaxInt32: "2LKcb1",
		math.MaxInt64: "AzL8n0Y58m7",
	}

	for k, v := range cases {
		e := StdEncoding.Encode(k)
		assert.Equal(t, v, e)
	}
}

func TestStdEncoding_Decode(t *testing.T) {
	cases := map[int64]string{
		0:             "0",
		57:            "v",
		math.MaxInt8:  "23",
		math.MaxInt16: "8WV",
		math.MaxInt32: "2LKcb1",
		math.MaxInt64: "AzL8n0Y58m7",
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
		0:             "0",
		57:            "v",
		math.MaxInt8:  "23",
		math.MaxInt16: "8WV",
		math.MaxInt32: "2LKcb1",
		math.MaxInt64: "AzL8n0Y58m7",
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
		0:             "0",
		57:            "v",
		math.MaxInt8:  "23",
		math.MaxInt16: "8WV",
		math.MaxInt32: "2LKcb1",
		math.MaxInt64: "AzL8n0Y58m7",
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
