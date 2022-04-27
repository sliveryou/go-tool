package slicex

import (
	"math/rand"
	"reflect"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Extract returns randomly extracted num elements from slice.
// It panics if slice or num is invalid.
func Extract(slice interface{}, num int) []interface{} {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}
	if slice == nil {
		return nil
	}
	value := reflect.ValueOf(slice)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		length := value.Len()
		if length == 0 {
			return nil
		}
		if num > length {
			num = length
		}
		result := make([]interface{}, num)
		for i, v := range rand.Perm(length) {
			if i < num {
				result[i] = value.Index(v).Interface()
			} else {
				break
			}
		}
		return result
	default:
		panic("slicex: invalid slice type")
	}
}

// ExtractStrings returns randomly extracted num elements from string slice.
// It panics if num is invalid.
func ExtractStrings(slice []string, num int) []string {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}
	length := len(slice)
	if length == 0 {
		return nil
	}
	if num > length {
		num = length
	}
	result := make([]string, num)
	for i, v := range rand.Perm(length) {
		if i < num {
			result[i] = slice[v]
		} else {
			break
		}
	}
	return result
}

// ExtractBools returns randomly extracted num elements from bool slice.
// It panics if num is invalid.
func ExtractBools(slice []bool, num int) []bool {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}
	length := len(slice)
	if length == 0 {
		return nil
	}
	if num > length {
		num = length
	}
	result := make([]bool, num)
	for i, v := range rand.Perm(length) {
		if i < num {
			result[i] = slice[v]
		} else {
			break
		}
	}
	return result
}

// ExtractInts returns randomly extracted num elements from int slice.
// It panics if num is invalid.
func ExtractInts(slice []int, num int) []int {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}
	length := len(slice)
	if length == 0 {
		return nil
	}
	if num > length {
		num = length
	}
	result := make([]int, num)
	for i, v := range rand.Perm(length) {
		if i < num {
			result[i] = slice[v]
		} else {
			break
		}
	}
	return result
}

// ExtractInt64s returns randomly extracted num elements from int64 slice.
// It panics if num is invalid.
func ExtractInt64s(slice []int64, num int) []int64 {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}
	length := len(slice)
	if length == 0 {
		return nil
	}
	if num > length {
		num = length
	}
	result := make([]int64, num)
	for i, v := range rand.Perm(length) {
		if i < num {
			result[i] = slice[v]
		} else {
			break
		}
	}
	return result
}

// ExtractInt32s returns randomly extracted num elements from int64 slice.
// It panics if num is invalid.
func ExtractInt32s(slice []int32, num int) []int32 {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}
	length := len(slice)
	if length == 0 {
		return nil
	}
	if num > length {
		num = length
	}
	result := make([]int32, num)
	for i, v := range rand.Perm(length) {
		if i < num {
			result[i] = slice[v]
		} else {
			break
		}
	}
	return result
}

// ExtractFloats returns randomly extracted num elements from float64 slice.
// It panics if num is invalid.
func ExtractFloats(slice []float64, num int) []float64 {
	return ExtractFloat64s(slice, num)
}

// ExtractFloat64s returns randomly extracted num elements from float64 slice.
// It panics if num is invalid.
func ExtractFloat64s(slice []float64, num int) []float64 {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}
	length := len(slice)
	if length == 0 {
		return nil
	}
	if num > length {
		num = length
	}
	result := make([]float64, num)
	for i, v := range rand.Perm(length) {
		if i < num {
			result[i] = slice[v]
		} else {
			break
		}
	}
	return result
}

// ExtractFloat32s returns randomly extracted num elements from float32 slice.
// It panics if num is invalid.
func ExtractFloat32s(slice []float32, num int) []float32 {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}
	length := len(slice)
	if length == 0 {
		return nil
	}
	if num > length {
		num = length
	}
	result := make([]float32, num)
	for i, v := range rand.Perm(length) {
		if i < num {
			result[i] = slice[v]
		} else {
			break
		}
	}
	return result
}
