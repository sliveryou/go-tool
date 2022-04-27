package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid slice type", func() {
		Join("err type")
	})

	assert.Equal(t, "a.b.c", Join([]string{"a", "b", "c"}, "."))
	assert.Equal(t, "a,b,c", Join([]string{"a", "b", "c"}))
	assert.Empty(t, Join([]string{}))
	assert.Empty(t, Join(nil))

	assert.Equal(t, "1.2.3", Join([]int{1, 2, 3}, "."))
	assert.Equal(t, "1,2,3", Join([]int{1, 2, 3}))
	assert.Empty(t, Join([]int{}))
	assert.Empty(t, Join(nil), ",")
}

func TestJoinStrings(t *testing.T) {
	assert.Equal(t, "a.b.c", JoinStrings([]string{"a", "b", "c"}, "."))
	assert.Equal(t, "a,b,c", JoinStrings([]string{"a", "b", "c"}))
	assert.Empty(t, JoinStrings([]string{}))
	assert.Empty(t, JoinStrings(nil))
}

func TestJoinBools(t *testing.T) {
	assert.Equal(t, "true.false.true", JoinBools([]bool{true, false, true}, "."))
	assert.Equal(t, "true,false,true", JoinBools([]bool{true, false, true}))
	assert.Empty(t, JoinBools([]bool{}))
	assert.Empty(t, JoinBools(nil))
}

func TestJoinInts(t *testing.T) {
	assert.Equal(t, "1.2.3", JoinInts([]int{1, 2, 3}, "."))
	assert.Equal(t, "1,2,3", JoinInts([]int{1, 2, 3}))
	assert.Empty(t, JoinInts([]int{}))
	assert.Empty(t, JoinInts(nil))
}

func TestJoinInt64s(t *testing.T) {
	assert.Equal(t, "1.2.3", JoinInt64s([]int64{1, 2, 3}, "."))
	assert.Equal(t, "1,2,3", JoinInt64s([]int64{1, 2, 3}))
	assert.Empty(t, JoinInt64s([]int64{}))
	assert.Empty(t, JoinInt64s(nil))
}

func TestJoinInt32s(t *testing.T) {
	assert.Equal(t, "1.2.3", JoinInt32s([]int32{1, 2, 3}, "."))
	assert.Equal(t, "1,2,3", JoinInt32s([]int32{1, 2, 3}))
	assert.Empty(t, JoinInt32s([]int32{}))
	assert.Empty(t, JoinInt32s(nil))
}

func TestJoinFloats(t *testing.T) {
	assert.Equal(t, "1.1-2.2-3.3", JoinFloats([]float64{1.1, 2.2, 3.3}, "-"))
	assert.Equal(t, "1.1,2.2,3.3", JoinFloats([]float64{1.1, 2.2, 3.3}, ","))
	assert.Empty(t, JoinFloats([]float64{}))
	assert.Empty(t, JoinFloats(nil))
}

func TestJoinFloat64s(t *testing.T) {
	assert.Equal(t, "1.1-2.2-3.3", JoinFloat64s([]float64{1.1, 2.2, 3.3}, "-"))
	assert.Equal(t, "1.1,2.2,3.3", JoinFloat64s([]float64{1.1, 2.2, 3.3}, ","))
	assert.Empty(t, JoinFloat64s([]float64{}))
	assert.Empty(t, JoinFloat64s(nil))
}

func TestJoinFloat32s(t *testing.T) {
	assert.Equal(t, "1.1-2.2-3.3", JoinFloat32s([]float32{1.1, 2.2, 3.3}, "-"))
	assert.Equal(t, "1.1,2.2,3.3", JoinFloat32s([]float32{1.1, 2.2, 3.3}, ","))
	assert.Empty(t, JoinFloat32s([]float32{}))
	assert.Empty(t, JoinFloat32s(nil))
}
