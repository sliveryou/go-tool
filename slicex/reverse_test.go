package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid slice type", func() {
		Reverse("err type")
	})

	assert.Equal(t,
		[]interface{}{"you", "sliver", "sliveryou"},
		Reverse([]string{"sliveryou", "sliver", "you"}),
	)
	assert.Empty(t, Reverse([]string{}))
	assert.Empty(t, Reverse(nil))

	assert.Equal(t,
		[]interface{}{4, 3, 2, 1},
		Reverse([]int{1, 2, 3, 4}),
	)
	assert.Empty(t, Reverse([]int{}))
	assert.Empty(t, Reverse(nil))
}

func TestReverseStrings(t *testing.T) {
	assert.Equal(t,
		[]string{"you", "sliver", "sliveryou"},
		ReverseStrings([]string{"sliveryou", "sliver", "you"}),
	)
	assert.Empty(t, ReverseStrings([]string{}))
	assert.Empty(t, ReverseStrings(nil))
}

func TestReverseBools(t *testing.T) {
	assert.Equal(t,
		[]bool{false, true, false, true},
		ReverseBools([]bool{true, false, true, false}),
	)
	assert.Empty(t, ReverseBools([]bool{}))
	assert.Empty(t, ReverseBools(nil))
}

func TestReverseInts(t *testing.T) {
	assert.Equal(t,
		[]int{4, 3, 2, 1},
		ReverseInts([]int{1, 2, 3, 4}),
	)
	assert.Empty(t, ReverseInts([]int{}))
	assert.Empty(t, ReverseInts(nil))
}

func TestReverseInt64s(t *testing.T) {
	assert.Equal(t,
		[]int64{4, 3, 2, 1},
		ReverseInt64s([]int64{1, 2, 3, 4}),
	)
	assert.Empty(t, ReverseInt64s([]int64{}))
	assert.Empty(t, ReverseInt64s(nil))
}

func TestReverseInt32s(t *testing.T) {
	assert.Equal(t,
		[]int32{4, 3, 2, 1},
		ReverseInt32s([]int32{1, 2, 3, 4}),
	)
	assert.Empty(t, ReverseInt32s([]int32{}))
	assert.Empty(t, ReverseInt32s(nil))
}

func TestReverseFloats(t *testing.T) {
	assert.Equal(t,
		[]float64{4.4, 3.3, 2.2, 1.1},
		ReverseFloats([]float64{1.1, 2.2, 3.3, 4.4}),
	)
	assert.Empty(t, ReverseFloats([]float64{}))
	assert.Empty(t, ReverseFloats(nil))
}

func TestReverseFloat64s(t *testing.T) {
	assert.Equal(t,
		[]float64{4.4, 3.3, 2.2, 1.1},
		ReverseFloat64s([]float64{1.1, 2.2, 3.3, 4.4}),
	)
	assert.Empty(t, ReverseFloat64s([]float64{}))
	assert.Empty(t, ReverseFloat64s(nil))
}

func TestReverseFloat32s(t *testing.T) {
	assert.Equal(t,
		[]float32{4.4, 3.3, 2.2, 1.1},
		ReverseFloat32s([]float32{1.1, 2.2, 3.3, 4.4}),
	)
	assert.Empty(t, ReverseFloat32s([]float32{}))
	assert.Empty(t, ReverseFloat32s(nil))
}
