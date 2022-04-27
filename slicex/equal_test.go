package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid arr1 or arr2 type", func() {
		Equal("err type1", "err type2")
	})

	arr1 := []string{"sliveryou", "sliver", "you"}
	arr2 := []string{"sliveryou", "sliver", "you", "go"}

	assert.False(t, Equal(arr1, arr2))
	assert.True(t, Equal(arr1, arr1))
	assert.False(t, Equal([]string{}, []int{}))
	assert.False(t, Equal([]string{}, nil))
	assert.True(t, Equal([]string{}, []string{}))
	assert.True(t, Equal(nil, nil))
}

func TestEqualStrings(t *testing.T) {
	arr1 := []string{"sliveryou", "sliver", "you"}
	arr2 := []string{"sliveryou", "you", "go"}

	assert.False(t, EqualStrings(arr1, arr2))
	assert.True(t, EqualStrings(arr1, arr1))
	assert.False(t, EqualStrings([]string{}, nil))
	assert.True(t, EqualStrings([]string{}, []string{}))
	assert.True(t, EqualStrings(nil, nil))
}

func TestEqualBools(t *testing.T) {
	arr1 := []bool{true, true, false}
	arr2 := []bool{false, true, false}

	assert.False(t, EqualBools(arr1, arr2))
	assert.True(t, EqualBools(arr1, arr1))
	assert.False(t, EqualBools([]bool{}, nil))
	assert.True(t, EqualBools([]bool{}, []bool{}))
	assert.True(t, EqualBools(nil, nil))
}

func TestEqualInts(t *testing.T) {
	arr1 := []int{12, 23, 34}
	arr2 := []int{23, 34, 45}

	assert.False(t, EqualInts(arr1, arr2))
	assert.True(t, EqualInts(arr1, arr1))
	assert.False(t, EqualInts([]int{}, nil))
	assert.True(t, EqualInts([]int{}, []int{}))
	assert.True(t, EqualInts(nil, nil))
}

func TestEqualInt64s(t *testing.T) {
	arr1 := []int64{12, 23, 34}
	arr2 := []int64{23, 34, 45}

	assert.False(t, EqualInt64s(arr1, arr2))
	assert.True(t, EqualInt64s(arr1, arr1))
	assert.False(t, EqualInt64s([]int64{}, nil))
	assert.True(t, EqualInt64s([]int64{}, []int64{}))
	assert.True(t, EqualInt64s(nil, nil))
}

func TestEqualInt32s(t *testing.T) {
	arr1 := []int32{12, 23, 34}
	arr2 := []int32{23, 34, 45}

	assert.False(t, EqualInt32s(arr1, arr2))
	assert.True(t, EqualInt32s(arr1, arr1))
	assert.False(t, EqualInt32s([]int32{}, nil))
	assert.True(t, EqualInt32s([]int32{}, []int32{}))
	assert.True(t, EqualInt32s(nil, nil))
}

func TestEqualFloats(t *testing.T) {
	arr1 := []float64{1.1, 2.2, 3.3}
	arr2 := []float64{2.2, 3.3, 4.4}

	assert.False(t, EqualFloats(arr1, arr2))
	assert.True(t, EqualFloats(arr1, arr1))
	assert.False(t, EqualFloats([]float64{}, nil))
	assert.True(t, EqualFloats([]float64{}, []float64{}))
	assert.True(t, EqualFloats(nil, nil))
}

func TestEqualFloat64s(t *testing.T) {
	arr1 := []float64{1.1, 2.2, 3.3}
	arr2 := []float64{2.2, 3.3, 4.4}

	assert.False(t, EqualFloat64s(arr1, arr2))
	assert.True(t, EqualFloat64s(arr1, arr1))
	assert.False(t, EqualFloat64s([]float64{}, nil))
	assert.True(t, EqualFloat64s([]float64{}, []float64{}))
	assert.True(t, EqualFloat64s(nil, nil))
}

func TestEqualFloat32s(t *testing.T) {
	arr1 := []float32{1.1, 2.2, 3.3}
	arr2 := []float32{2.2, 3.3, 4.4}

	assert.False(t, EqualFloat32s(arr1, arr2))
	assert.True(t, EqualFloat32s(arr1, arr1))
	assert.False(t, EqualFloat32s([]float32{}, nil))
	assert.True(t, EqualFloat32s([]float32{}, []float32{}))
	assert.True(t, EqualFloat32s(nil, nil))
}
