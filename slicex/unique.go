package slicex

import (
	"reflect"
)

// Unique returns the unique slice.
// It panics if slice is invalid.
func Unique(slice interface{}) []interface{} {
	if slice == nil {
		return []interface{}{}
	}

	value := reflect.ValueOf(slice)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		result := make([]interface{}, 0, value.Len())
		kinds := []reflect.Kind{
			reflect.Interface, reflect.Slice, reflect.Array,
			reflect.Map, reflect.Struct, reflect.Func,
		}
		uniqueSlice, uniqueMap := make([]interface{}, 0), make(map[interface{}]struct{})

		for i := 0; i < value.Len(); i++ {
			item := value.Index(i).Interface()
			kind := reflect.ValueOf(item).Kind()
			if Contain(kinds, kind) == -1 {
				if _, ok := uniqueMap[item]; !ok {
					uniqueMap[item] = struct{}{}
					result = append(result, item)
				}
			} else {
				if Contain(uniqueSlice, item) == -1 {
					uniqueSlice = append(uniqueSlice, item)
					result = append(result, item)
				}
			}
		}

		return result
	default:
		panic("slicex: invalid slice type")
	}
}

// UniqueStrings returns the unique string slice.
func UniqueStrings(slice []string) []string {
	result := make([]string, 0, len(slice))
	uniqueMap := make(map[string]struct{}, len(slice))

	for _, value := range slice {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// UniqueBools returns the unique bool slice.
func UniqueBools(slice []bool) []bool {
	result := make([]bool, 0, len(slice))
	uniqueMap := make(map[bool]struct{}, len(slice))

	for _, value := range slice {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// UniqueInts returns the unique int slice.
func UniqueInts(slice []int) []int {
	result := make([]int, 0, len(slice))
	uniqueMap := make(map[int]struct{}, len(slice))

	for _, value := range slice {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// UniqueInt64s returns the unique int64 slice.
func UniqueInt64s(slice []int64) []int64 {
	result := make([]int64, 0, len(slice))
	uniqueMap := make(map[int64]struct{}, len(slice))

	for _, value := range slice {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// UniqueInt32s returns the unique int32 slice.
func UniqueInt32s(slice []int32) []int32 {
	result := make([]int32, 0, len(slice))
	uniqueMap := make(map[int32]struct{}, len(slice))

	for _, value := range slice {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// UniqueFloats returns the unique float64 slice.
func UniqueFloats(slice []float64) []float64 {
	return UniqueFloat64s(slice)
}

// UniqueFloat64s returns the unique float64 slice.
func UniqueFloat64s(slice []float64) []float64 {
	result := make([]float64, 0, len(slice))
	uniqueMap := make(map[float64]struct{}, len(slice))

	for _, value := range slice {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// UniqueFloat32s returns the unique float32 slice.
func UniqueFloat32s(slice []float32) []float32 {
	result := make([]float32, 0, len(slice))
	uniqueMap := make(map[float32]struct{}, len(slice))

	for _, value := range slice {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}
