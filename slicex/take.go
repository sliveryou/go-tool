package slicex

import (
	"reflect"
)

// Take returns the randomly taken element from slice,
// it is equivalent to Extract(slice, 1)[0].
func Take(slice interface{}) interface{} {
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
		return Extract(slice, 1)[0]
	default:
		panic("slicex: invalid slice type")
	}
}

// TakeString returns the randomly taken string element from string slice,
// it is equivalent to ExtractStrings(slice, 1)[0].
func TakeString(slice []string) string {
	length := len(slice)
	if length == 0 {
		return ""
	}
	return ExtractStrings(slice, 1)[0]
}

// TakeBool returns the randomly taken bool element from bool slice,
// it is equivalent to ExtractBools(slice, 1)[0].
func TakeBool(slice []bool) bool {
	length := len(slice)
	if length == 0 {
		return false
	}
	return ExtractBools(slice, 1)[0]
}

// TakeInt returns the randomly taken int element from int slice,
// it is equivalent to ExtractInts(slice, 1)[0].
func TakeInt(slice []int) int {
	length := len(slice)
	if length == 0 {
		return 0
	}
	return ExtractInts(slice, 1)[0]
}

// TakeInt64 returns the randomly taken int64 element from int64 slice,
// it is equivalent to ExtractInt64s(slice, 1)[0].
func TakeInt64(slice []int64) int64 {
	length := len(slice)
	if length == 0 {
		return 0
	}
	return ExtractInt64s(slice, 1)[0]
}

// TakeInt32 returns the randomly taken int32 element from int32 slice,
// it is equivalent to ExtractInt32s(slice, 1)[0].
func TakeInt32(slice []int32) int32 {
	length := len(slice)
	if length == 0 {
		return 0
	}
	return ExtractInt32s(slice, 1)[0]
}

// TakeFloat returns the randomly taken float64 element from float64 slice,
// it is equivalent to ExtractFloats(slice, 1)[0].
func TakeFloat(slice []float64) float64 {
	return TakeFloat64(slice)
}

// TakeFloat64 returns the randomly taken float64 element from float64 slice,
// it is equivalent to ExtractFloat64s(slice, 1)[0].
func TakeFloat64(slice []float64) float64 {
	length := len(slice)
	if length == 0 {
		return 0
	}
	return ExtractFloat64s(slice, 1)[0]
}

// TakeFloat32 returns the randomly taken float32 element from float32 slice,
// it is equivalent to ExtractFloat32s(slice, 1)[0].
func TakeFloat32(slice []float32) float32 {
	length := len(slice)
	if length == 0 {
		return 0
	}
	return ExtractFloat32s(slice, 1)[0]
}
