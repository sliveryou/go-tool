package slicex

import (
	"reflect"

	"github.com/sliveryou/go-tool/mathx"
)

const (
	// ValueIsNotContained means the value is not contained in the slice.
	ValueIsNotContained = -1
)

// Contain returns the index of the first instance of value in slice,
// or -1 if value is not present in slice.
// It panics if slice is invalid.
func Contain(slice, value interface{}) (index int) {
	index = ValueIsNotContained
	if slice == nil {
		return
	}
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice, reflect.Array:
		v := reflect.ValueOf(slice)
		for i := 0; i < v.Len(); i++ {
			if reflect.DeepEqual(value, v.Index(i).Interface()) {
				index = i
				return
			}
		}
	default:
		panic("slicex: invalid slice type")
	}
	return
}

// ContainString returns the index of the first instance of string value in string slice,
// or -1 if value is not present in string slice.
func ContainString(slice []string, value string) (index int) {
	index = ValueIsNotContained
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			index = i
			return
		}
	}
	return
}

// ContainBool returns the index of the first instance of bool value in bool slice,
// or -1 if value is not present in bool slice.
func ContainBool(slice []bool, value bool) (index int) {
	index = ValueIsNotContained
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			index = i
			return
		}
	}
	return
}

// ContainInt returns the index of the first instance of int value in int slice,
// or -1 if value is not present in int slice.
func ContainInt(slice []int, value int) (index int) {
	index = ValueIsNotContained
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			index = i
			return
		}
	}
	return
}

// ContainInt64 returns the index of the first instance of int64 value in int64 slice,
// or -1 if value is not present in int64 slice.
func ContainInt64(slice []int64, value int64) (index int) {
	index = ValueIsNotContained
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			index = i
			return
		}
	}
	return
}

// ContainInt32 returns the index of the first instance of int32 value in int32 slice,
// or -1 if value is not present in int32 slice.
func ContainInt32(slice []int32, value int32) (index int) {
	index = ValueIsNotContained
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			index = i
			return
		}
	}
	return
}

// ContainFloat returns the index of the first instance of float64 value in float64 slice,
// or -1 if value is not present in float64 slice.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func ContainFloat(slice []float64, value float64, places ...int) (index int) {
	return ContainFloat64(slice, value, places...)
}

// ContainFloat64 returns the index of the first instance of float64 value in float64 slice,
// or -1 if value is not present in float64 slice.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func ContainFloat64(slice []float64, value float64, places ...int) (index int) {
	index = ValueIsNotContained
	for i := 0; i < len(slice); i++ {
		if mathx.Equal(slice[i], value, places...) {
			index = i
			return
		}
	}
	return
}

// ContainFloat32 returns the index of the first instance of float32 value in float32 slice,
// or -1 if value is not present in float32 slice.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func ContainFloat32(slice []float32, value float32, places ...int) (index int) {
	index = ValueIsNotContained
	for i := 0; i < len(slice); i++ {
		if mathx.Equal(float64(slice[i]), float64(value), places...) {
			index = i
			return
		}
	}
	return
}
