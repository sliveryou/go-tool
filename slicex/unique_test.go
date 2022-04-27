package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid slice type", func() {
		Unique("err type")
	})

	arr1 := []interface{}{"sliveryou", "sliver", "you"}
	arr2 := []string{"sliveryou", "sliver", "sliver", "you", "you", "sliveryou"}

	assert.Equal(t, arr1, Unique(arr2))
	assert.Empty(t, Unique([]string{}))
	assert.Empty(t, Unique(nil))

	arr3 := []interface{}{1, 2, 3}
	arr4 := []int{1, 1, 2, 2, 3, 3}

	assert.Equal(t, arr3, Unique(arr4))
	assert.Empty(t, Unique([]int{}))
	assert.Empty(t, Unique(nil))
}

func TestUniqueStrings(t *testing.T) {
	arr1 := []string{"sliveryou", "sliver", "you"}
	arr2 := []string{"sliveryou", "sliver", "sliver", "you", "you", "sliveryou"}

	assert.Equal(t, arr1, UniqueStrings(arr2))
	assert.Empty(t, UniqueStrings([]string{}))
	assert.Empty(t, UniqueStrings(nil))
}

func TestUniqueBools(t *testing.T) {
	arr1 := []bool{true, false}
	arr2 := []bool{true, true, true, false, false, false}

	assert.Equal(t, arr1, UniqueBools(arr2))
	assert.Empty(t, UniqueBools([]bool{}))
	assert.Empty(t, UniqueBools(nil))
}

func TestUniqueInts(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 1, 2, 2, 3, 3}

	assert.Equal(t, arr1, UniqueInts(arr2))
	assert.Empty(t, UniqueInts([]int{}))
	assert.Empty(t, UniqueInts(nil))
}

func TestUniqueInt64s(t *testing.T) {
	arr1 := []int64{1, 2, 3}
	arr2 := []int64{1, 1, 2, 2, 3, 3}

	assert.Equal(t, arr1, UniqueInt64s(arr2))
	assert.Empty(t, UniqueInt64s([]int64{}))
	assert.Empty(t, UniqueInt64s(nil))
}

func TestUniqueInt32s(t *testing.T) {
	arr1 := []int32{1, 2, 3}
	arr2 := []int32{1, 1, 2, 2, 3, 3}

	assert.Equal(t, arr1, UniqueInt32s(arr2))
	assert.Empty(t, UniqueInt32s([]int32{}))
	assert.Empty(t, UniqueInt32s(nil))
}

func TestUniqueFloats(t *testing.T) {
	arr1 := []float64{1.1, 2.2, 3.3}
	arr2 := []float64{1.1, 1.1, 2.2, 2.2, 3.3, 3.3}

	assert.Equal(t, arr1, UniqueFloats(arr2))
	assert.Empty(t, UniqueFloats([]float64{}))
	assert.Empty(t, UniqueFloats(nil))
}

func TestUniqueFloat64s(t *testing.T) {
	arr1 := []float64{1.1, 2.2, 3.3}
	arr2 := []float64{1.1, 1.1, 2.2, 2.2, 3.3, 3.3}

	assert.Equal(t, arr1, UniqueFloat64s(arr2))
	assert.Empty(t, UniqueFloat64s([]float64{}))
	assert.Empty(t, UniqueFloat64s(nil))
}

func TestUniqueFloat32s(t *testing.T) {
	arr1 := []float32{1.1, 2.2, 3.3}
	arr2 := []float32{1.1, 1.1, 2.2, 2.2, 3.3, 3.3}

	assert.Equal(t, arr1, UniqueFloat32s(arr2))
	assert.Empty(t, UniqueFloat32s([]float32{}))
	assert.Empty(t, UniqueFloat32s(nil))
}
