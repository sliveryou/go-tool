package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid slice type", func() {
		Extract("err type", 3)
	})
	assert.PanicsWithValue(t, "slicex: num cannot be less than 0", func() {
		Extract([]string{}, -1)
	})

	arr1 := []string{"sliveryou", "sliver", "you", "a", "b", "c", "d"}
	assert.Len(t, Extract(arr1, 3), 3)
	assert.Len(t, Extract(arr1, 10), 7)
	assert.Empty(t, Extract(arr1, 0))
	assert.Empty(t, Extract(nil, 0))

	arr2 := []int{5, 6, 7, 8, 9, 10, 11}
	assert.Len(t, Extract(arr2, 3), 3)
	assert.Len(t, Extract(arr2, 10), 7)
	assert.Empty(t, Extract(arr2, 0))
	assert.Empty(t, Extract(nil, 10))
}

func TestExtractStrings(t *testing.T) {
	arr := []string{"sliveryou", "sliver", "you", "a", "b", "c", "d"}

	assert.Len(t, ExtractStrings(arr, 3), 3)
	assert.Len(t, ExtractStrings(arr, 10), 7)
	assert.Empty(t, ExtractStrings(arr, 0))
	assert.Empty(t, ExtractStrings([]string{}, 10))
	assert.Empty(t, ExtractStrings(nil, 10))
}

func TestExtractBools(t *testing.T) {
	arr := []bool{true, false, true, true, true, false, false}

	assert.Len(t, ExtractBools(arr, 3), 3)
	assert.Len(t, ExtractBools(arr, 10), 7)
	assert.Empty(t, ExtractBools(arr, 0))
	assert.Empty(t, ExtractBools([]bool{}, 10))
	assert.Empty(t, ExtractBools(nil, 10))
}

func TestExtractInts(t *testing.T) {
	arr := []int{5, 6, 7, 8, 9, 10, 11}

	assert.Len(t, ExtractInts(arr, 3), 3)
	assert.Len(t, ExtractInts(arr, 10), 7)
	assert.Empty(t, ExtractInts(arr, 0))
	assert.Empty(t, ExtractInts([]int{}, 10))
	assert.Empty(t, ExtractInts(nil, 10))
}

func TestExtractInt64s(t *testing.T) {
	arr := []int64{5, 6, 7, 8, 9, 10, 11}

	assert.Len(t, ExtractInt64s(arr, 3), 3)
	assert.Len(t, ExtractInt64s(arr, 10), 7)
	assert.Empty(t, ExtractInt64s(arr, 0))
	assert.Empty(t, ExtractInt64s([]int64{}, 10))
	assert.Empty(t, ExtractInt64s(nil, 10))
}

func TestExtractInt32s(t *testing.T) {
	arr := []int32{5, 6, 7, 8, 9, 10, 11}

	assert.Len(t, ExtractInt32s(arr, 3), 3)
	assert.Len(t, ExtractInt32s(arr, 10), 7)
	assert.Empty(t, ExtractInt32s(arr, 0))
	assert.Empty(t, ExtractInt32s([]int32{}, 10))
	assert.Empty(t, ExtractInt32s(nil, 10))
}

func TestExtractFloats(t *testing.T) {
	arr := []float64{5.5, 6.6, 7.7, 8.8, 9.9, 10.10, 11.11}

	assert.Len(t, ExtractFloats(arr, 3), 3)
	assert.Len(t, ExtractFloats(arr, 10), 7)
	assert.Empty(t, ExtractFloats(arr, 0))
	assert.Empty(t, ExtractFloats([]float64{}, 10))
	assert.Empty(t, ExtractFloats(nil, 10))
}

func TestExtractFloat64s(t *testing.T) {
	arr := []float64{5.5, 6.6, 7.7, 8.8, 9.9, 10.10, 11.11}

	assert.Len(t, ExtractFloat64s(arr, 3), 3)
	assert.Len(t, ExtractFloat64s(arr, 10), 7)
	assert.Empty(t, ExtractFloat64s(arr, 0))
	assert.Empty(t, ExtractFloat64s([]float64{}, 10))
	assert.Empty(t, ExtractFloat64s(nil, 10))
}

func TestExtractFloat32s(t *testing.T) {
	arr := []float32{5.5, 6.6, 7.7, 8.8, 9.9, 10.10, 11.11}

	assert.Len(t, ExtractFloat32s(arr, 3), 3)
	assert.Len(t, ExtractFloat32s(arr, 10), 7)
	assert.Empty(t, ExtractFloat32s(arr, 0))
	assert.Empty(t, ExtractFloat32s([]float32{}, 10))
	assert.Empty(t, ExtractFloat32s(nil, 10))
}
