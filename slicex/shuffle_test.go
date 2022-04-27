package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid slice type", func() {
		Shuffle("err type")
	})

	arr1 := []string{"sliveryou", "sliver", "you", "go", "abc"}
	assert.ElementsMatch(t, arr1, Shuffle(arr1))
	assert.Empty(t, Shuffle([]string{}))
	assert.Empty(t, Shuffle(nil))

	arr2 := []int{12, 34, 56, 78, 910}
	assert.ElementsMatch(t, arr2, Shuffle(arr2))
	assert.Empty(t, Shuffle([]int{}))
	assert.Empty(t, Shuffle(nil))
}

func TestShuffleStrings(t *testing.T) {
	arr := []string{"sliveryou", "sliver", "you", "go", "abc"}

	assert.ElementsMatch(t, arr, ShuffleStrings(arr))
	assert.Empty(t, ShuffleStrings([]string{}))
	assert.Empty(t, ShuffleStrings(nil))
}

func TestShuffleBools(t *testing.T) {
	arr := []bool{true, true, false, false, true}

	assert.ElementsMatch(t, arr, ShuffleBools(arr))
	assert.Empty(t, ShuffleBools([]bool{}))
	assert.Empty(t, ShuffleBools(nil))
}

func TestShuffleInts(t *testing.T) {
	arr := []int{12, 34, 56, 78, 910}

	assert.ElementsMatch(t, arr, ShuffleInts(arr))
	assert.Empty(t, ShuffleInts([]int{}))
	assert.Empty(t, ShuffleInts(nil))
}

func TestShuffleInt64s(t *testing.T) {
	arr := []int64{12, 34, 56, 78, 910}

	assert.ElementsMatch(t, arr, ShuffleInt64s(arr))
	assert.Empty(t, ShuffleInt64s([]int64{}))
	assert.Empty(t, ShuffleInt64s(nil))
}

func TestShuffleInt32s(t *testing.T) {
	arr := []int32{12, 34, 56, 78, 910}

	assert.ElementsMatch(t, arr, ShuffleInt32s(arr))
	assert.Empty(t, ShuffleInt32s([]int32{}))
	assert.Empty(t, ShuffleInt32s(nil))
}

func TestShuffleFloats(t *testing.T) {
	arr := []float64{12, 34, 56, 78, 910}

	assert.ElementsMatch(t, arr, ShuffleFloats(arr))
	assert.Empty(t, ShuffleFloats([]float64{}))
	assert.Empty(t, ShuffleFloats(nil))
}

func TestShuffleFloat64s(t *testing.T) {
	arr := []float64{12, 34, 56, 78, 910}

	assert.ElementsMatch(t, arr, ShuffleFloat64s(arr))
	assert.Empty(t, ShuffleFloat64s([]float64{}))
	assert.Empty(t, ShuffleFloat64s(nil))
}

func TestShuffleFloat32s(t *testing.T) {
	arr := []float32{12, 34, 56, 78, 910}

	assert.ElementsMatch(t, arr, ShuffleFloat32s(arr))
	assert.Empty(t, ShuffleFloat32s([]float32{}))
	assert.Empty(t, ShuffleFloat32s(nil))
}
