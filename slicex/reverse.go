package slicex

import "reflect"

// Reverse returns the reverse order for slice.
// It panics if slice is invalid.
func Reverse(slice interface{}) []interface{} {
	if slice == nil {
		return []interface{}{}
	}

	value := reflect.ValueOf(slice)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		length := value.Len()
		result := make([]interface{}, length)
		i, j := 0, length-1

		for ; i < j; i, j = i+1, j-1 {
			result[i], result[j] = value.Index(j).Interface(), value.Index(i).Interface()
		}
		if length > 0 && length%2 == 1 {
			result[j] = value.Index(j).Interface()
		}

		return result
	default:
		panic("slicex: invalid slice type")
	}
}

// ReverseStrings returns the reverse order for string slice.
func ReverseStrings(slice []string) []string {
	length := len(slice)
	result := make([]string, length)
	i, j := 0, length-1

	for ; i < j; i, j = i+1, j-1 {
		result[i], result[j] = slice[j], slice[i]
	}
	if length > 0 && length%2 == 1 {
		result[j] = slice[j]
	}

	return result
}

// ReverseBools returns the reverse order for bool slice.
func ReverseBools(slice []bool) []bool {
	length := len(slice)
	result := make([]bool, length)
	i, j := 0, length-1

	for ; i < j; i, j = i+1, j-1 {
		result[i], result[j] = slice[j], slice[i]
	}
	if length > 0 && length%2 == 1 {
		result[j] = slice[j]
	}

	return result
}

// ReverseInts returns the reverse order for int slice.
func ReverseInts(slice []int) []int {
	length := len(slice)
	result := make([]int, length)
	i, j := 0, length-1

	for ; i < j; i, j = i+1, j-1 {
		result[i], result[j] = slice[j], slice[i]
	}
	if length > 0 && length%2 == 1 {
		result[j] = slice[j]
	}

	return result
}

// ReverseInt64s returns the reverse order for int64 slice.
func ReverseInt64s(slice []int64) []int64 {
	length := len(slice)
	result := make([]int64, length)
	i, j := 0, length-1

	for ; i < j; i, j = i+1, j-1 {
		result[i], result[j] = slice[j], slice[i]
	}
	if length > 0 && length%2 == 1 {
		result[j] = slice[j]
	}

	return result
}

// ReverseInt32s returns the reverse order for int32 slice.
func ReverseInt32s(slice []int32) []int32 {
	length := len(slice)
	result := make([]int32, length)
	i, j := 0, length-1

	for ; i < j; i, j = i+1, j-1 {
		result[i], result[j] = slice[j], slice[i]
	}
	if length > 0 && length%2 == 1 {
		result[j] = slice[j]
	}

	return result
}

// ReverseFloats returns the reverse order for float64 slice.
func ReverseFloats(slice []float64) []float64 {
	return ReverseFloat64s(slice)
}

// ReverseFloat64s returns the reverse order for float64 slice.
func ReverseFloat64s(slice []float64) []float64 {
	length := len(slice)
	result := make([]float64, length)
	i, j := 0, length-1

	for ; i < j; i, j = i+1, j-1 {
		result[i], result[j] = slice[j], slice[i]
	}
	if length > 0 && length%2 == 1 {
		result[j] = slice[j]
	}

	return result
}

// ReverseFloat32s returns the reverse order for float64 slice.
func ReverseFloat32s(slice []float32) []float32 {
	length := len(slice)
	result := make([]float32, length)
	i, j := 0, length-1

	for ; i < j; i, j = i+1, j-1 {
		result[i], result[j] = slice[j], slice[i]
	}
	if length > 0 && length%2 == 1 {
		result[j] = slice[j]
	}

	return result
}
