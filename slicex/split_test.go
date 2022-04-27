package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitStrings(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}

	assert.Equal(t, arr, SplitStrings("a,b,c,d,e"))
	assert.Equal(t, arr, SplitStrings("a.b.c.d.e", "."))
	assert.Equal(t, []string{"abc"}, SplitStrings("abc", "."))
	assert.Empty(t, SplitStrings(""))
	assert.Empty(t, SplitStrings("", ""))
}

func TestSplitBools(t *testing.T) {
	arr := []bool{true, true, false, false, true}

	assert.Equal(t, arr, SplitBools("true,true,false,false,true"))
	assert.Equal(t, arr, SplitBools("true.true.false.false.true", "."))
	assert.Equal(t, []bool{true}, SplitBools("true", "."))
	assert.Equal(t, []bool{false}, SplitBools("true,true,true", "."))
	assert.Empty(t, SplitBools(""))
	assert.Empty(t, SplitBools("", ""))
}

func TestSplitInts(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	assert.Equal(t, arr, SplitInts("1,2,3,4,5"))
	assert.Equal(t, arr, SplitInts("1.2.3.4.5", "."))
	assert.Equal(t, []int{123}, SplitInts("123", "."))
	assert.Equal(t, []int{0}, SplitInts("1,2,3", "."))
	assert.Empty(t, SplitInts(""))
	assert.Empty(t, SplitInts("", ""))
}

func TestSplitInt64s(t *testing.T) {
	arr := []int64{1, 2, 3, 4, 5}

	assert.Equal(t, arr, SplitInt64s("1,2,3,4,5"))
	assert.Equal(t, arr, SplitInt64s("1.2.3.4.5", "."))
	assert.Equal(t, []int64{123}, SplitInt64s("123", "."))
	assert.Equal(t, []int64{0}, SplitInt64s("1,2,3", "."))
	assert.Empty(t, SplitInt64s(""))
	assert.Empty(t, SplitInt64s("", ""))
}

func TestSplitInt32s(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5}

	assert.Equal(t, arr, SplitInt32s("1,2,3,4,5"))
	assert.Equal(t, arr, SplitInt32s("1.2.3.4.5", "."))
	assert.Equal(t, []int32{123}, SplitInt32s("123", "."))
	assert.Equal(t, []int32{0}, SplitInt32s("1,2,3", "."))
	assert.Empty(t, SplitInt32s(""))
	assert.Empty(t, SplitInt32s("", ""))
}

func TestSplitFloats(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	assert.Equal(t, arr, SplitFloats("1.1,2.2,3.3,4.4,5.5"))
	assert.Equal(t, arr, SplitFloats("1.1+2.2+3.3+4.4+5.5", "+"))
	assert.Equal(t, []float64{123.123}, SplitFloats("123.123", "+"))
	assert.Equal(t, []float64{0}, SplitFloats("1.2.3", "+"))
	assert.Empty(t, SplitFloats(""))
	assert.Empty(t, SplitFloats("", ""))
}

func TestSplitFloat64s(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	assert.Equal(t, arr, SplitFloat64s("1.1,2.2,3.3,4.4,5.5"))
	assert.Equal(t, arr, SplitFloat64s("1.1+2.2+3.3+4.4+5.5", "+"))
	assert.Equal(t, []float64{123.123}, SplitFloat64s("123.123", "+"))
	assert.Equal(t, []float64{0}, SplitFloat64s("1.2.3", "+"))
	assert.Empty(t, SplitFloat64s(""))
	assert.Empty(t, SplitFloat64s("", ""))
}

func TestSplitFloat32s(t *testing.T) {
	arr := []float32{1.1, 2.2, 3.3, 4.4, 5.5}

	assert.Equal(t, arr, SplitFloat32s("1.1,2.2,3.3,4.4,5.5"))
	assert.Equal(t, arr, SplitFloat32s("1.1+2.2+3.3+4.4+5.5", "+"))
	assert.Equal(t, []float32{123.123}, SplitFloat32s("123.123", "+"))
	assert.Equal(t, []float32{0}, SplitFloat32s("1.2.3", "+"))
	assert.Empty(t, SplitFloat32s(""))
	assert.Empty(t, SplitFloat32s("", ""))
}
