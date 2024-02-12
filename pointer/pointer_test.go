package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	t.Parallel()

	result1 := Of(123)
	result2 := Of("abc")

	assert.Equal(t, 123, *result1)
	assert.Equal(t, "abc", *result2)
}

func TestUnwrap(t *testing.T) {
	t.Parallel()

	a := 123
	b := "abc"

	assert.Equal(t, a, Unwrap(&a))
	assert.Equal(t, b, Unwrap(&b))
}

func TestUnwrapOr(t *testing.T) {
	t.Parallel()

	a := 123
	b := "abc"

	var c *int
	var d *string

	assert.Equal(t, a, UnwrapOr(&a, 456))
	assert.Equal(t, b, UnwrapOr(&b, "abc"))
	assert.Equal(t, 456, UnwrapOr(c, 456))
	assert.Equal(t, "def", UnwrapOr(d, "def"))
}

func TestUnwrapOrDefault(t *testing.T) {
	t.Parallel()

	a := 123
	b := "abc"

	var c *int
	var d *string

	assert.Equal(t, a, UnwrapOrDefault(&a))
	assert.Equal(t, b, UnwrapOrDefault(&b))
	assert.Equal(t, 0, UnwrapOrDefault(c))
	assert.Equal(t, "", UnwrapOrDefault(d))
}

func TestExtractPointer(t *testing.T) {
	t.Parallel()

	a := 1
	b := &a
	c := &b
	d := &c

	assert.Equal(t, 1, ExtractPointer(d))
}
