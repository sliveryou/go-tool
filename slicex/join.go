package slicex

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/sliveryou/go-tool/v2/convert"
)

const (
	joinSeparator = ","
)

// Join returns string result of slice connected with string sep,
// default sep is ",".
// It panics if slice is invalid.
func Join(slice interface{}, sep ...string) (result string) {
	if slice == nil {
		return ""
	}
	value := reflect.ValueOf(slice)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		length := value.Len()
		if length == 0 {
			return
		}
		separator := joinSeparator
		if len(sep) != 0 {
			separator = sep[0]
		}
		var builder strings.Builder
		for i := 0; i < value.Len(); i++ {
			builder.WriteString(convert.ToString(value.Index(i).Interface()))
			if length--; length > 0 {
				builder.WriteString(separator)
			}
		}
		result = builder.String()
		return
	default:
		panic("slicex: invalid slice type")
	}
}

// JoinStrings returns string result of string slice connected with string sep,
// default sep is ",".
func JoinStrings(slice []string, sep ...string) (result string) {
	length := len(slice)
	if length == 0 {
		return
	}
	separator := joinSeparator
	if len(sep) != 0 {
		separator = sep[0]
	}
	var builder strings.Builder
	for _, value := range slice {
		builder.WriteString(value)
		if length--; length > 0 {
			builder.WriteString(separator)
		}
	}
	result = builder.String()
	return
}

// JoinBools returns string result of bool slice connected with string sep,
// default sep is ",".
func JoinBools(slice []bool, sep ...string) (result string) {
	length := len(slice)
	if length == 0 {
		return
	}
	separator := joinSeparator
	if len(sep) != 0 {
		separator = sep[0]
	}
	var builder strings.Builder
	for _, value := range slice {
		builder.WriteString(strconv.FormatBool(value))
		if length--; length > 0 {
			builder.WriteString(separator)
		}
	}
	result = builder.String()
	return
}

// JoinInts returns string result of int slice connected with string sep,
// default sep is ",".
func JoinInts(slice []int, sep ...string) (result string) {
	length := len(slice)
	if length == 0 {
		return
	}
	separator := joinSeparator
	if len(sep) != 0 {
		separator = sep[0]
	}
	var builder strings.Builder
	for _, value := range slice {
		builder.WriteString(strconv.Itoa(value))
		if length--; length > 0 {
			builder.WriteString(separator)
		}
	}
	result = builder.String()
	return
}

// JoinInt64s returns string result of int64 slice connected with string sep,
// default sep is ",".
func JoinInt64s(slice []int64, sep ...string) (result string) {
	length := len(slice)
	if length == 0 {
		return
	}
	separator := joinSeparator
	if len(sep) != 0 {
		separator = sep[0]
	}
	var builder strings.Builder
	for _, value := range slice {
		builder.WriteString(strconv.FormatInt(value, 10))
		if length--; length > 0 {
			builder.WriteString(separator)
		}
	}
	result = builder.String()
	return
}

// JoinInt32s returns string result of int32 slice connected with string sep,
// default sep is ",".
func JoinInt32s(slice []int32, sep ...string) (result string) {
	length := len(slice)
	if length == 0 {
		return
	}
	separator := joinSeparator
	if len(sep) != 0 {
		separator = sep[0]
	}
	var builder strings.Builder
	for _, value := range slice {
		builder.WriteString(strconv.FormatInt(int64(value), 10))
		if length--; length > 0 {
			builder.WriteString(separator)
		}
	}
	result = builder.String()
	return
}

// JoinFloats returns string result of float64 slice connected with string sep,
// default sep is ",".
func JoinFloats(slice []float64, sep ...string) (result string) {
	return JoinFloat64s(slice, sep...)
}

// JoinFloat64s returns string result of float64 slice connected with string sep,
// default sep is ",".
func JoinFloat64s(slice []float64, sep ...string) (result string) {
	length := len(slice)
	if length == 0 {
		return
	}
	separator := joinSeparator
	if len(sep) != 0 {
		separator = sep[0]
	}
	var builder strings.Builder
	for _, value := range slice {
		builder.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
		if length--; length > 0 {
			builder.WriteString(separator)
		}
	}
	result = builder.String()
	return
}

// JoinFloat32s returns string result of float32 slice connected with string sep,
// default sep is ",".
func JoinFloat32s(slice []float32, sep ...string) (result string) {
	length := len(slice)
	if length == 0 {
		return
	}
	separator := joinSeparator
	if len(sep) != 0 {
		separator = sep[0]
	}
	var builder strings.Builder
	for _, value := range slice {
		builder.WriteString(strconv.FormatFloat(float64(value), 'f', -1, 32))
		if length--; length > 0 {
			builder.WriteString(separator)
		}
	}
	result = builder.String()
	return
}
