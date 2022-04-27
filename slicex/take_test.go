package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid slice type", func() {
		Take("err type")
	})

	arr1 := []string{"sliveryou", "sliver", "you", "go", "abc"}
	assert.Contains(t, arr1, Take(arr1))
	assert.Empty(t, Take([]string{}))
	assert.Empty(t, Take(nil))

	arr2 := []int{12, 34, 56, 78, 910}
	assert.Contains(t, arr2, Take(arr2))
	assert.Empty(t, Take([]int{}))
	assert.Empty(t, Take(nil))
}

func TestTakeString(t *testing.T) {
	arr := []string{"sliveryou", "sliver", "you", "go", "abc"}

	assert.Contains(t, arr, TakeString(arr))
	assert.Empty(t, TakeString([]string{}))
	assert.Empty(t, TakeString(nil))
}

func TestTakeBool(t *testing.T) {
	arr := []bool{true, false, true, false, true}

	assert.Contains(t, arr, TakeBool(arr))
	assert.Empty(t, TakeBool([]bool{}))
	assert.Empty(t, TakeBool(nil))
}

func TestTakeInt(t *testing.T) {
	arr := []int{12, 34, 56, 78, 910}

	assert.Contains(t, arr, TakeInt(arr))
	assert.Empty(t, TakeInt([]int{}))
	assert.Empty(t, TakeInt(nil))
}

func TestTakeInt64(t *testing.T) {
	arr := []int64{12, 34, 56, 78, 910}

	assert.Contains(t, arr, TakeInt64(arr))
	assert.Empty(t, TakeInt64([]int64{}))
	assert.Empty(t, TakeInt64(nil))
}

func TestTakeInt32(t *testing.T) {
	arr := []int32{12, 34, 56, 78, 910}

	assert.Contains(t, arr, TakeInt32(arr))
	assert.Empty(t, TakeInt32([]int32{}))
	assert.Empty(t, TakeInt32(nil))
}

func TestTakeFloat(t *testing.T) {
	arr := []float64{12.12, 34.34, 56.56, 78.78, 90.90}

	assert.Contains(t, arr, TakeFloat(arr))
	assert.Empty(t, TakeFloat([]float64{}))
	assert.Empty(t, TakeFloat(nil))
}

func TestTakeFloat64(t *testing.T) {
	arr := []float64{12.12, 34.34, 56.56, 78.78, 90.90}

	assert.Contains(t, arr, TakeFloat64(arr))
	assert.Empty(t, TakeFloat64([]float64{}))
	assert.Empty(t, TakeFloat64(nil))
}

func TestTakeFloat32(t *testing.T) {
	arr := []float32{12.12, 34.34, 56.56, 78.78, 90.90}

	assert.Contains(t, arr, TakeFloat32(arr))
	assert.Empty(t, TakeFloat32([]float32{}))
	assert.Empty(t, TakeFloat32(nil))
}
