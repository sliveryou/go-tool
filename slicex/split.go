package slicex

import (
	"strings"

	"github.com/sliveryou/go-tool/convert"
)

const (
	splitSeparator = ","
)

// SplitStrings slices string str into substrings separated by string sep and
// returns string slice of the converted substrings between those separators,
// default sep is ",".
func SplitStrings(str string, sep ...string) []string {
	if str == "" {
		return []string{}
	}

	separator := splitSeparator
	if len(sep) > 0 {
		separator = sep[0]
	}

	return strings.Split(str, separator)
}

// SplitBools slices string str into substrings separated by string sep and
// returns bool slice of the converted substrings between those separators,
// default sep is ",".
func SplitBools(str string, sep ...string) []bool {
	if str == "" {
		return []bool{}
	}

	separator := splitSeparator
	if len(sep) > 0 {
		separator = sep[0]
	}

	s := strings.Split(str, separator)
	result := make([]bool, 0, len(s))
	for _, v := range s {
		result = append(result, convert.ToBool(v))
	}

	return result
}

// SplitInts slices string str into substrings separated by string sep and
// returns int slice of the converted substrings between those separators,
// default sep is ",".
func SplitInts(str string, sep ...string) []int {
	if str == "" {
		return []int{}
	}

	separator := splitSeparator
	if len(sep) > 0 {
		separator = sep[0]
	}

	s := strings.Split(str, separator)
	result := make([]int, 0, len(s))
	for _, v := range s {
		result = append(result, convert.ToInt(v))
	}

	return result
}

// SplitInt64s slices string str into substrings separated by string sep and
// returns int64 slice of the converted substrings between those separators,
// default sep is ",".
func SplitInt64s(str string, sep ...string) []int64 {
	if str == "" {
		return []int64{}
	}

	separator := splitSeparator
	if len(sep) > 0 {
		separator = sep[0]
	}

	s := strings.Split(str, separator)
	result := make([]int64, 0, len(s))
	for _, v := range s {
		result = append(result, convert.ToInt64(v))
	}

	return result
}

// SplitInt32s slices string str into substrings separated by string sep and
// returns int32 slice of the converted substrings between those separators,
// default sep is ",".
func SplitInt32s(str string, sep ...string) []int32 {
	if str == "" {
		return []int32{}
	}

	separator := splitSeparator
	if len(sep) > 0 {
		separator = sep[0]
	}

	s := strings.Split(str, separator)
	result := make([]int32, 0, len(s))
	for _, v := range s {
		result = append(result, convert.ToInt32(v))
	}

	return result
}

// SplitFloats slices string str into substrings separated by string sep and
// returns float64 slice of the converted substrings between those separators,
// default sep is ",".
func SplitFloats(str string, sep ...string) []float64 {
	return SplitFloat64s(str, sep...)
}

// SplitFloat64s slices string str into substrings separated by string sep and
// returns float64 slice of the converted substrings between those separators,
// default sep is ",".
func SplitFloat64s(str string, sep ...string) []float64 {
	if str == "" {
		return []float64{}
	}

	separator := splitSeparator
	if len(sep) > 0 {
		separator = sep[0]
	}

	s := strings.Split(str, separator)
	result := make([]float64, 0, len(s))
	for _, v := range s {
		result = append(result, convert.ToFloat64(v))
	}

	return result
}

// SplitFloat32s slices string str into substrings separated by string sep and
// returns float32 slice of the converted substrings between those separators,
// default sep is ",".
func SplitFloat32s(str string, sep ...string) []float32 {
	if str == "" {
		return []float32{}
	}

	separator := splitSeparator
	if len(sep) > 0 {
		separator = sep[0]
	}

	s := strings.Split(str, separator)
	result := make([]float32, 0, len(s))
	for _, v := range s {
		result = append(result, convert.ToFloat32(v))
	}

	return result
}
