package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContain(t *testing.T) {
	assertion := assert.New(t)
	assertion.PanicsWithValue("slicex: invalid slice type", func() {
		Contain("err type", "err")
	})
	arr := []interface{}{"sliveryou", 123456, 123.456, "go"}

	assertion.Equal(3, Contain(arr, "go"))
	assertion.Equal(1, Contain(arr, 123456))
	assertion.Equal(-1, Contain(arr, "sliver"))
	assertion.Equal(-1, Contain([]string{}, "sliver"))
	assertion.Equal(2, Contain([]string{"sliver", "sliveryou", "go"}, "go"))
	assertion.Equal(1, Contain([]int{123, 456, 789}, 456))
	assertion.Equal(-1, Contain(nil, "nil"))
	assertion.Equal(-1, Contain(nil, 0))
}

func TestContainString(t *testing.T) {
	assertion := assert.New(t)
	arr := []string{"sliveryou", "sliver", "go"}

	assertion.Equal(1, ContainString(arr, "sliver"))
	assertion.Equal(2, ContainString(arr, "go"))
	assertion.Equal(-1, ContainString(arr, "you"))
	assertion.Equal(-1, ContainString([]string{}, "you"))
	assertion.Equal(-1, ContainString(nil, "nil"))
}

func TestContainRune(t *testing.T) {
	assertion := assert.New(t)
	arr := []rune{'a', 'b', '你', '好'}

	assertion.Equal(1, ContainRune(arr, 'b'))
	assertion.Equal(2, ContainRune(arr, '你'))
	assertion.Equal(-1, ContainRune(arr, '啦'))
	assertion.Equal(-1, ContainRune([]rune{}, 's'))
	assertion.Equal(-1, ContainRune(nil, 'i'))
}

func TestContainBool(t *testing.T) {
	assertion := assert.New(t)
	arr := []bool{false, false, true}

	assertion.Equal(0, ContainBool(arr, false))
	assertion.Equal(2, ContainBool(arr, true))
	assertion.Equal(-1, ContainBool([]bool{}, true))
	assertion.Equal(-1, ContainBool(nil, true))
}

func TestContainInt(t *testing.T) {
	assertion := assert.New(t)
	arr := []int{1, 3, 5, 7, 9}

	assertion.Equal(1, ContainInt(arr, 3))
	assertion.Equal(3, ContainInt(arr, 7))
	assertion.Equal(-1, ContainInt(arr, 10))
	assertion.Equal(-1, ContainInt([]int{}, 1))
	assertion.Equal(-1, ContainInt(nil, 1))
}

func TestContainInt64(t *testing.T) {
	assertion := assert.New(t)
	arr := []int64{1, 3, 5, 7, 9}

	assertion.Equal(1, ContainInt64(arr, 3))
	assertion.Equal(3, ContainInt64(arr, 7))
	assertion.Equal(-1, ContainInt64(arr, 10))
	assertion.Equal(-1, ContainInt64([]int64{}, 1))
	assertion.Equal(-1, ContainInt64(nil, 1))
}

func TestContainInt32(t *testing.T) {
	assertion := assert.New(t)
	arr := []int32{1, 3, 5, 7, 9}

	assertion.Equal(1, ContainInt32(arr, 3))
	assertion.Equal(3, ContainInt32(arr, 7))
	assertion.Equal(-1, ContainInt32(arr, 10))
	assertion.Equal(-1, ContainInt32([]int32{}, 1))
	assertion.Equal(-1, ContainInt32(nil, 1))
}

func TestContainFloat(t *testing.T) {
	assertion := assert.New(t)
	arr := []float64{0.3, 1.2, 3.4, 5.6}
	v1, v2 := 0.1, 0.2

	t.Log(v1 + v2)
	assertion.Equal(0, ContainFloat(arr, v1+v2))
	assertion.Equal(-1, ContainFloat(arr, v1+v2, 17))
	assertion.Equal(2, ContainFloat(arr, 3.4))
	assertion.Equal(3, ContainFloat(arr, 5.6))
	assertion.Equal(-1, ContainFloat(arr, 12.34))
	assertion.Equal(-1, ContainFloat([]float64{}, 1.2))
	assertion.Equal(-1, ContainFloat(nil, 1.2))
}

func TestContainFloat64(t *testing.T) {
	assertion := assert.New(t)
	arr := []float64{0.3, 1.2, 3.4, 5.6}
	v1, v2 := 0.1, 0.2

	t.Log(v1 + v2)
	assertion.Equal(0, ContainFloat64(arr, v1+v2))
	assertion.Equal(-1, ContainFloat64(arr, v1+v2, 17))
	assertion.Equal(2, ContainFloat64(arr, 3.4))
	assertion.Equal(3, ContainFloat64(arr, 5.6))
	assertion.Equal(-1, ContainFloat64(arr, 12.34))
	assertion.Equal(-1, ContainFloat64([]float64{}, 1.2))
	assertion.Equal(-1, ContainFloat64(nil, 1.2))
}

func TestContainFloat32(t *testing.T) {
	assertion := assert.New(t)
	arr := []float32{0.3, 1.2, 3.4, 5.6}
	v1, v2 := float32(0.1), float32(0.2)

	t.Log(v1 + v2)
	assertion.Equal(0, ContainFloat32(arr, v1+v2))
	assertion.Equal(0, ContainFloat32(arr, v1+v2, 17))
	assertion.Equal(2, ContainFloat32(arr, 3.4))
	assertion.Equal(3, ContainFloat32(arr, 5.6))
	assertion.Equal(-1, ContainFloat32(arr, 12.34))
	assertion.Equal(-1, ContainFloat32([]float32{}, 1.2))
	assertion.Equal(-1, ContainFloat32(nil, 1.2))
}
