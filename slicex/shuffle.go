package slicex

import (
	"reflect"
)

// Shuffle returns shuffled slice,
// it is equivalent to Extract(slice, len(slice)).
func Shuffle(slice interface{}) []interface{} {
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
		return Extract(slice, length)
	default:
		panic("slicex: invalid slice type")
	}
}

// ShuffleStrings returns shuffled string slice,
// it is equivalent to ExtractStrings(slice, len(slice)).
func ShuffleStrings(slice []string) []string {
	length := len(slice)
	if length == 0 {
		return nil
	}
	return ExtractStrings(slice, length)
}

// ShuffleBools returns shuffled bool slice,
// it is equivalent to ExtractBools(slice, len(slice)).
func ShuffleBools(slice []bool) []bool {
	length := len(slice)
	if length == 0 {
		return nil
	}
	return ExtractBools(slice, length)
}

// ShuffleInts returns shuffled int slice,
// it is equivalent to ExtractInts(slice, len(slice)).
func ShuffleInts(slice []int) []int {
	length := len(slice)
	if length == 0 {
		return nil
	}
	return ExtractInts(slice, length)
}

// ShuffleInt64s returns shuffled int64 slice,
// it is equivalent to ExtractInt64s(slice, len(slice)).
func ShuffleInt64s(slice []int64) []int64 {
	length := len(slice)
	if length == 0 {
		return nil
	}
	return ExtractInt64s(slice, length)
}

// ShuffleInt32s returns shuffled int32 slice,
// it is equivalent to ExtractInt32s(slice, len(slice)).
func ShuffleInt32s(slice []int32) []int32 {
	length := len(slice)
	if length == 0 {
		return nil
	}
	return ExtractInt32s(slice, length)
}

// ShuffleFloats returns shuffled float64 slice,
// it is equivalent to ExtractFloats(slice, len(slice)).
func ShuffleFloats(slice []float64) []float64 {
	return ShuffleFloat64s(slice)
}

// ShuffleFloat64s returns shuffled float64 slice,
// it is equivalent to ExtractFloat64s(slice, len(slice)).
func ShuffleFloat64s(slice []float64) []float64 {
	length := len(slice)
	if length == 0 {
		return nil
	}
	return ExtractFloat64s(slice, length)
}

// ShuffleFloat32s returns shuffled float32 slice,
// it is equivalent to ExtractFloat32s(slice, len(slice)).
func ShuffleFloat32s(slice []float32) []float32 {
	length := len(slice)
	if length == 0 {
		return nil
	}
	return ExtractFloat32s(slice, length)
}
