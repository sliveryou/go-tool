package slicex

import (
	"reflect"

	"github.com/sliveryou/go-tool/v2/mathx"
)

// Equal reports whether arr1 equals arr2.
// It panics if arr1 or arr2 is invalid.
func Equal(arr1, arr2 interface{}) bool {
	if arr1 == nil || arr2 == nil {
		return arr1 == arr2
	}

	kind1, kind2 := reflect.TypeOf(arr1).Kind(), reflect.TypeOf(arr2).Kind()
	if kind1 == reflect.Slice && kind2 == reflect.Slice {
		return reflect.DeepEqual(arr1, arr2)
	} else if kind1 == reflect.Array && kind2 == reflect.Array {
		return arr1 == arr2
	}

	panic("slicex: invalid arr1 or arr2 type")
}

// EqualStrings reports whether string arr1 equals string arr2.
func EqualStrings(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}

// EqualBools reports whether bool arr1 equals bool arr2.
func EqualBools(arr1, arr2 []bool) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}

// EqualInts reports whether int arr1 equals int arr2.
func EqualInts(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}

// EqualInt64s reports whether int64 arr1 equals int64 arr2.
func EqualInt64s(arr1, arr2 []int64) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}

// EqualInt32s reports whether int32 arr1 equals int32 arr2.
func EqualInt32s(arr1, arr2 []int32) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}

// EqualFloats reports whether float64 arr1 equals float64 arr2.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func EqualFloats(arr1, arr2 []float64, places ...int) bool {
	return EqualFloat64s(arr1, arr2, places...)
}

// EqualFloat64s reports whether float64 arr1 equals float64 arr2.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func EqualFloat64s(arr1, arr2 []float64, places ...int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	for i, v := range arr1 {
		if !mathx.Equal(v, arr2[i], places...) {
			return false
		}
	}

	return true
}

// EqualFloat32s reports whether float32 arr1 equals float32 arr2.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func EqualFloat32s(arr1, arr2 []float32, places ...int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	if (arr1 == nil) != (arr2 == nil) {
		return false
	}

	for i, v := range arr1 {
		if !mathx.Equal(float64(v), float64(arr2[i]), places...) {
			return false
		}
	}

	return true
}
