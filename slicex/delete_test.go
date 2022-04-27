package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	assert.PanicsWithValue(t, "slicex: invalid slice type", func() {
		Delete("err type", "error", 1)
	})

	arr := []interface{}{"1", 2, "a", "1", 1, "a", "b", "1", 1, 2, 1}

	cases := []struct {
		value     interface{}
		n         int
		expectArr []interface{}
		expectNum int
	}{
		{value: "1", n: 0, expectArr: []interface{}{"1", 2, "a", "1", 1, "a", "b", "1", 1, 2, 1}, expectNum: 0},
		{value: "1", n: 2, expectArr: []interface{}{2, "a", 1, "a", "b", "1", 1, 2, 1}, expectNum: 2},
		{value: "1", n: 5, expectArr: []interface{}{2, "a", 1, "a", "b", 1, 2, 1}, expectNum: 3},
		{value: "1", n: -1, expectArr: []interface{}{2, "a", 1, "a", "b", 1, 2, 1}, expectNum: 3},
		{value: 1, n: 0, expectArr: []interface{}{"1", 2, "a", "1", 1, "a", "b", "1", 1, 2, 1}, expectNum: 0},
		{value: 1, n: 2, expectArr: []interface{}{"1", 2, "a", "1", "a", "b", "1", 2, 1}, expectNum: 2},
		{value: 1, n: 5, expectArr: []interface{}{"1", 2, "a", "1", "a", "b", "1", 2}, expectNum: 3},
		{value: 1, n: -1, expectArr: []interface{}{"1", 2, "a", "1", "a", "b", "1", 2}, expectNum: 3},
	}

	for _, c := range cases {
		getArr, getNum := Delete(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}

func TestDeleteString(t *testing.T) {
	arr := []string{"a", "b", "b", "a", "a", "c", "d", "a"}

	cases := []struct {
		value     string
		n         int
		expectArr []string
		expectNum int
	}{
		{value: "h", n: 10, expectArr: []string{"a", "b", "b", "a", "a", "c", "d", "a"}, expectNum: 0},
		{value: "a", n: 0, expectArr: []string{"a", "b", "b", "a", "a", "c", "d", "a"}, expectNum: 0},
		{value: "a", n: 1, expectArr: []string{"b", "b", "a", "a", "c", "d", "a"}, expectNum: 1},
		{value: "a", n: 2, expectArr: []string{"b", "b", "a", "c", "d", "a"}, expectNum: 2},
		{value: "a", n: 10, expectArr: []string{"b", "b", "c", "d"}, expectNum: 4},
		{value: "a", n: -1, expectArr: []string{"b", "b", "c", "d"}, expectNum: 4},
	}

	for _, c := range cases {
		getArr, getNum := DeleteString(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}

func TestDeleteBool(t *testing.T) {
	arr := []bool{true, false, true, false, true, false, false}

	cases := []struct {
		value     bool
		n         int
		expectArr []bool
		expectNum int
	}{
		{value: true, n: 0, expectArr: []bool{true, false, true, false, true, false, false}, expectNum: 0},
		{value: true, n: 2, expectArr: []bool{false, false, true, false, false}, expectNum: 2},
		{value: true, n: 5, expectArr: []bool{false, false, false, false}, expectNum: 3},
		{value: true, n: -1, expectArr: []bool{false, false, false, false}, expectNum: 3},
	}

	for _, c := range cases {
		getArr, getNum := DeleteBool(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}

func TestDeleteInt(t *testing.T) {
	arr := []int{1, 2, 3, 1, 2, 3, 1, 1, 2, 3}

	cases := []struct {
		value     int
		n         int
		expectArr []int
		expectNum int
	}{
		{value: 1, n: 0, expectArr: []int{1, 2, 3, 1, 2, 3, 1, 1, 2, 3}, expectNum: 0},
		{value: 1, n: 2, expectArr: []int{2, 3, 2, 3, 1, 1, 2, 3}, expectNum: 2},
		{value: 1, n: 5, expectArr: []int{2, 3, 2, 3, 2, 3}, expectNum: 4},
		{value: 1, n: -1, expectArr: []int{2, 3, 2, 3, 2, 3}, expectNum: 4},
	}

	for _, c := range cases {
		getArr, getNum := DeleteInt(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}

func TestDeleteInt64(t *testing.T) {
	arr := []int64{1, 2, 3, 1, 2, 3, 1, 1, 2, 3}

	cases := []struct {
		value     int64
		n         int
		expectArr []int64
		expectNum int
	}{
		{value: 1, n: 0, expectArr: []int64{1, 2, 3, 1, 2, 3, 1, 1, 2, 3}, expectNum: 0},
		{value: 1, n: 2, expectArr: []int64{2, 3, 2, 3, 1, 1, 2, 3}, expectNum: 2},
		{value: 1, n: 5, expectArr: []int64{2, 3, 2, 3, 2, 3}, expectNum: 4},
		{value: 1, n: -1, expectArr: []int64{2, 3, 2, 3, 2, 3}, expectNum: 4},
	}

	for _, c := range cases {
		getArr, getNum := DeleteInt64(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}

func TestDeleteInt32(t *testing.T) {
	arr := []int32{1, 2, 3, 1, 2, 3, 1, 1, 2, 3}

	cases := []struct {
		value     int32
		n         int
		expectArr []int32
		expectNum int
	}{
		{value: 1, n: 0, expectArr: []int32{1, 2, 3, 1, 2, 3, 1, 1, 2, 3}, expectNum: 0},
		{value: 1, n: 2, expectArr: []int32{2, 3, 2, 3, 1, 1, 2, 3}, expectNum: 2},
		{value: 1, n: 5, expectArr: []int32{2, 3, 2, 3, 2, 3}, expectNum: 4},
		{value: 1, n: -1, expectArr: []int32{2, 3, 2, 3, 2, 3}, expectNum: 4},
	}

	for _, c := range cases {
		getArr, getNum := DeleteInt32(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}

func TestDeleteFloat(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3, 1.1, 2.2, 1.1, 2.2, 3.3, 1.1}

	cases := []struct {
		value     float64
		n         int
		expectArr []float64
		expectNum int
	}{
		{value: 1.1, n: 0, expectArr: []float64{1.1, 2.2, 3.3, 1.1, 2.2, 1.1, 2.2, 3.3, 1.1}, expectNum: 0},
		{value: 1.1, n: 2, expectArr: []float64{2.2, 3.3, 2.2, 1.1, 2.2, 3.3, 1.1}, expectNum: 2},
		{value: 1.1, n: 5, expectArr: []float64{2.2, 3.3, 2.2, 2.2, 3.3}, expectNum: 4},
		{value: 1.1, n: -1, expectArr: []float64{2.2, 3.3, 2.2, 2.2, 3.3}, expectNum: 4},
	}

	for _, c := range cases {
		getArr, getNum := DeleteFloat(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}

func TestDeleteFloat64(t *testing.T) {
	arr := []float64{1.1, 2.2, 3.3, 1.1, 2.2, 1.1, 2.2, 3.3, 1.1}

	cases := []struct {
		value     float64
		n         int
		expectArr []float64
		expectNum int
	}{
		{value: 1.1, n: 0, expectArr: []float64{1.1, 2.2, 3.3, 1.1, 2.2, 1.1, 2.2, 3.3, 1.1}, expectNum: 0},
		{value: 1.1, n: 2, expectArr: []float64{2.2, 3.3, 2.2, 1.1, 2.2, 3.3, 1.1}, expectNum: 2},
		{value: 1.1, n: 5, expectArr: []float64{2.2, 3.3, 2.2, 2.2, 3.3}, expectNum: 4},
		{value: 1.1, n: -1, expectArr: []float64{2.2, 3.3, 2.2, 2.2, 3.3}, expectNum: 4},
	}

	for _, c := range cases {
		getArr, getNum := DeleteFloat64(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}

func TestDeleteFloat32(t *testing.T) {
	arr := []float32{1.1, 2.2, 3.3, 1.1, 2.2, 1.1, 2.2, 3.3, 1.1}

	cases := []struct {
		value     float32
		n         int
		expectArr []float32
		expectNum int
	}{
		{value: 1.1, n: 0, expectArr: []float32{1.1, 2.2, 3.3, 1.1, 2.2, 1.1, 2.2, 3.3, 1.1}, expectNum: 0},
		{value: 1.1, n: 2, expectArr: []float32{2.2, 3.3, 2.2, 1.1, 2.2, 3.3, 1.1}, expectNum: 2},
		{value: 1.1, n: 5, expectArr: []float32{2.2, 3.3, 2.2, 2.2, 3.3}, expectNum: 4},
		{value: 1.1, n: -1, expectArr: []float32{2.2, 3.3, 2.2, 2.2, 3.3}, expectNum: 4},
	}

	for _, c := range cases {
		getArr, getNum := DeleteFloat32(arr, c.value, c.n)
		assert.Equal(t, c.expectArr, getArr)
		assert.Equal(t, c.expectNum, getNum)
	}
}
