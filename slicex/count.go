package slicex

import (
	"reflect"
)

// Count returns value count map by slice.
func Count(slice interface{}) map[interface{}]int {
	if slice == nil {
		return map[interface{}]int{}
	}

	value := reflect.ValueOf(slice)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		kinds := []reflect.Kind{
			reflect.Interface, reflect.Slice, reflect.Array,
			reflect.Map, reflect.Struct, reflect.Func,
		}

		countMap := make(map[interface{}]int, value.Len())
		for i := 0; i < value.Len(); i++ {
			item := value.Index(i).Interface()
			kind := reflect.ValueOf(item).Kind()
			if Contain(kinds, kind) == -1 {
				countMap[item]++
			}
		}

		return countMap
	default:
		panic("slicex: invalid slice type")
	}
}

// CountString returns string value count map by string slice.
func CountString(slice []string) map[string]int {
	result := make(map[string]int, len(slice))
	for _, v := range slice {
		result[v]++
	}

	return result
}

// CountBool returns bool value count map by bool slice.
func CountBool(slice []bool) map[bool]int {
	result := make(map[bool]int, len(slice))
	for _, v := range slice {
		result[v]++
	}

	return result
}

// CountInt returns int value count map by int slice.
func CountInt(slice []int) map[int]int {
	result := make(map[int]int, len(slice))
	for _, v := range slice {
		result[v]++
	}

	return result
}

// CountInt64 returns int64 value count map by int64 slice.
func CountInt64(slice []int64) map[int64]int {
	result := make(map[int64]int, len(slice))
	for _, v := range slice {
		result[v]++
	}

	return result
}

// CountInt32 returns int32 value count map by int32 slice.
func CountInt32(slice []int32) map[int32]int {
	result := make(map[int32]int, len(slice))
	for _, v := range slice {
		result[v]++
	}

	return result
}

// CountFloat returns float64 value count map by float64 slice.
func CountFloat(slice []float64) map[float64]int {
	return CountFloat64(slice)
}

// CountFloat64 returns float64 value count map by float64 slice.
func CountFloat64(slice []float64) map[float64]int {
	result := make(map[float64]int, len(slice))
	for _, v := range slice {
		result[v]++
	}

	return result
}

// CountFloat32 returns float32 value count map by float32 slice.
func CountFloat32(slice []float32) map[float32]int {
	result := make(map[float32]int, len(slice))
	for _, v := range slice {
		result[v]++
	}

	return result
}
