package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid slice type", func() {
		Count("err type")
	})

	m1 := map[interface{}]int{"a": 2, 1: 2, true: 2, 3.3: 3}
	m2 := Count([]interface{}{"a", 1, true, 3.3, "a", 1, true, 3.3, 3.3})

	assert.Equal(t, m1, m2)
	assert.Empty(t, Count([]interface{}{}))
	assert.Empty(t, Count(nil))
}

func TestCountString(t *testing.T) {
	m1 := map[string]int{"a": 2, "b": 2, "c": 1}
	m2 := CountString([]string{"a", "a", "b", "c", "b"})

	assert.Equal(t, m1, m2)
	assert.Empty(t, CountString([]string{}))
	assert.Empty(t, CountString(nil))
}

func TestCountBool(t *testing.T) {
	m1 := map[bool]int{true: 2, false: 3}
	m2 := CountBool([]bool{true, false, true, false, false})

	assert.Equal(t, m1, m2)
	assert.Empty(t, CountBool([]bool{}))
	assert.Empty(t, CountBool(nil))
}

func TestCountInt(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	m2 := CountInt([]int{1, 2, 3, 2, 3, 3})

	assert.Equal(t, m1, m2)
	assert.Empty(t, CountInt([]int{}))
	assert.Empty(t, CountInt(nil))
}

func TestCountInt64(t *testing.T) {
	m1 := map[int64]int{1: 1, 2: 2, 3: 3}
	m2 := CountInt64([]int64{1, 2, 3, 2, 3, 3})

	assert.Equal(t, m1, m2)
	assert.Empty(t, CountInt64([]int64{}))
	assert.Empty(t, CountInt64(nil))
}

func TestCountInt32(t *testing.T) {
	m1 := map[int32]int{1: 1, 2: 2, 3: 3}
	m2 := CountInt32([]int32{1, 2, 3, 2, 3, 3})

	assert.Equal(t, m1, m2)
	assert.Empty(t, CountInt32([]int32{}))
	assert.Empty(t, CountInt32(nil))
}

func TestCountFloat(t *testing.T) {
	m1 := map[float64]int{1.1: 1, 2.2: 2, 3.3: 3}
	m2 := CountFloat([]float64{1.1, 2.2, 3.3, 2.2, 3.3, 3.3})

	assert.Equal(t, m1, m2)
	assert.Empty(t, CountFloat([]float64{}))
	assert.Empty(t, CountFloat(nil))
}

func TestCountFloat64(t *testing.T) {
	m1 := map[float64]int{1.1: 1, 2.2: 2, 3.3: 3}
	m2 := CountFloat64([]float64{1.1, 2.2, 3.3, 2.2, 3.3, 3.3})

	assert.Equal(t, m1, m2)
	assert.Empty(t, CountFloat64([]float64{}))
	assert.Empty(t, CountFloat64(nil))
}

func TestCountFloat32(t *testing.T) {
	m1 := map[float32]int{1.1: 1, 2.2: 2, 3.3: 3}
	m2 := CountFloat32([]float32{1.1, 2.2, 3.3, 2.2, 3.3, 3.3})

	assert.Equal(t, m1, m2)
	assert.Empty(t, CountFloat32([]float32{}))
	assert.Empty(t, CountFloat32(nil))
}
