package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: num cannot be less than 0", func() {
		Fill([]interface{}{}, -1)
	})

	assert.Equal(t, []interface{}{"abc", "abc", "abc"}, Fill("abc", 3))
	assert.Empty(t, Fill("abc", 0))

	assert.Equal(t, []interface{}{123, 123, 123}, Fill(123, 3))
	assert.Empty(t, Fill(123, 0))
}

func TestFillString(t *testing.T) {
	assert.Equal(t, []string{"abc", "abc", "abc"}, FillString("abc", 3))
	assert.Empty(t, FillString("abc", 0))
}

func TestFillBool(t *testing.T) {
	assert.Equal(t, []bool{true, true, true}, FillBool(true, 3))
	assert.Empty(t, FillBool(true, 0))
}

func TestFillInt(t *testing.T) {
	assert.Equal(t, []int{123, 123, 123}, FillInt(123, 3))
	assert.Empty(t, FillInt(123, 0))
}

func TestFillInt64(t *testing.T) {
	assert.Equal(t, []int64{123, 123, 123}, FillInt64(123, 3))
	assert.Empty(t, FillInt64(123, 0))
}

func TestFillInt32(t *testing.T) {
	assert.Equal(t, []int32{123, 123, 123}, FillInt32(123, 3))
	assert.Empty(t, FillInt32(123, 0))
}

func TestFillFloat(t *testing.T) {
	assert.Equal(t, []float64{123.123, 123.123, 123.123}, FillFloat(123.123, 3))
	assert.Empty(t, FillFloat(123.123, 0))
}

func TestFillFloat64(t *testing.T) {
	assert.Equal(t, []float64{123.123, 123.123, 123.123}, FillFloat64(123.123, 3))
	assert.Empty(t, FillFloat64(123.123, 0))
}

func TestFillFloat32(t *testing.T) {
	assert.Equal(t, []float32{123.123, 123.123, 123.123}, FillFloat32(123.123, 3))
	assert.Empty(t, FillFloat32(123.123, 0))
}
