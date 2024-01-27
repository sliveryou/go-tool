package mathx

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/sliveryou/go-tool/convert"
)

// math package constants.
const (
	floatPlaces    = 9
	floatSeparator = ","
)

// math package variables.
var (
	sizeUnits = []string{
		"B", "KB", "MB", "GB", "TB", "PB",
		"EB", "ZB", "YB", "BB", "NB", "DB", "CB", "XB", "UnKnown",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Round rounds the f to places.
func Round(f float64, places int) float64 {
	result, _ := decimal.NewFromFloat(f).Round(int32(places)).Float64()
	return result
}

// RoundBank banker rounds the f to places.
func RoundBank(f float64, places int) float64 {
	result, _ := decimal.NewFromFloat(f).RoundBank(int32(places)).Float64()
	return result
}

// RoundToString returns a rounded fixed-point string with places digits after
// the f point.
func RoundToString(f float64, places int) string {
	return decimal.NewFromFloat(f).StringFixed(int32(places))
}

// RoundBankToString returns a banker rounded fixed-point string with places digits
// after the f point.
func RoundBankToString(f float64, places int) string {
	return decimal.NewFromFloat(f).StringFixedBank(int32(places))
}

// Sign returns:
//
//	-1 if f <  0
//	 0 if f == 0
//	+1 if f >  0
func Sign(f float64) int {
	if math.Float64bits(f) != 0 {
		if f > 0 {
			return 1
		}
		return -1
	}
	return 0
}

// IsPositive returns:
//
//	 true if d >  0
//	false if d == 0
//	false if d <  0
func IsPositive(f float64) bool {
	return Sign(f) == 1
}

// IsNonPositive returns:
//
//	 true if d <  0
//	 true if d == 0
//	false if d >  0
func IsNonPositive(f float64) bool {
	return Sign(f) == -1 || Sign(f) == 0
}

// IsNegative returns:
//
//	 true if d <  0
//	false if d == 0
//	false if d >  0
func IsNegative(f float64) bool {
	return Sign(f) == -1
}

// IsNonNegative returns:
//
//	 true if d >  0
//	 true if d == 0
//	false if d <  0
func IsNonNegative(f float64) bool {
	return Sign(f) == 1 || Sign(f) == 0
}

// IsZero returns:
//
//	 true if d == 0
//	false if d >  0
//	false if d <  0
func IsZero(f float64) bool {
	return Sign(f) == 0
}

// Compare compares f1 and f2 and returns:
//
//	-1 if   f1 < f2
//	 0 if | f1 - f2 | <= 10 ^ -places, default places is 9
//	+1 if   f1 > f2
func Compare(f1, f2 float64, places ...int) int {
	pla := floatPlaces
	if len(places) != 0 && places[0] >= 1 {
		pla = places[0]
	}
	isEqual := math.Abs(f1-f2) <= math.Pow10(-pla)
	if !isEqual {
		if f1 > f2 {
			return 1
		}
		return -1
	}
	return 0
}

// Equal reports whether f1 and f2 are equal.
func Equal(f1, f2 float64, places ...int) bool {
	return Compare(f1, f2, places...) == 0
}

// GreaterThan (GT) returns true when f1 is greater than f2.
func GreaterThan(f1, f2 float64, places ...int) bool {
	return Compare(f1, f2, places...) == 1
}

// LessThan (LT) returns true when f1 is less than f2.
func LessThan(f1, f2 float64, places ...int) bool {
	return Compare(f1, f2, places...) == -1
}

// GreaterThanOrEqual (GTE) returns true when f1 is greater than or equal to f2.
func GreaterThanOrEqual(f1, f2 float64, places ...int) bool {
	cmp := Compare(f1, f2, places...)
	return cmp == 1 || cmp == 0
}

// LessThanOrEqual (LTE) returns true when f1 is less than or equal to f2.
func LessThanOrEqual(f1, f2 float64, places ...int) bool {
	cmp := Compare(f1, f2, places...)
	return cmp == -1 || cmp == 0
}

// IsOdd reports whether num is odd number.
func IsOdd(num int64) bool {
	return num%2 != 0
}

// IsEven reports whether num is even number.
func IsEven(num int64) bool {
	return num%2 == 0
}

// RangeInt returns a int slice, starting from int start, and increments by int step,
// and stops before int stop.
// If start < stop, default step is 1, else default step is -1.
func RangeInt(start, stop int, step ...int) []int {
	var result []int
	var s int
	if start < stop {
		s = 1
		if len(step) != 0 && step[0] > 1 {
			s = step[0]
		}
		for {
			result = append(result, start)
			start += s
			if start >= stop {
				break
			}
		}
	} else if start > stop {
		s = -1
		if len(step) != 0 && step[0] < -1 {
			s = step[0]
		}
		for {
			result = append(result, start)
			start += s
			if start <= stop {
				break
			}
		}
	}
	return result
}

// RangeInt64 returns a int64 slice, starting from int64 start, and increments by int64 step,
// and stops before int64 stop.
// If start < stop, default step is 1, else default step is -1.
func RangeInt64(start, stop int64, step ...int64) []int64 {
	var result []int64
	var s int64
	if start < stop {
		s = 1
		if len(step) != 0 && step[0] > 1 {
			s = step[0]
		}
		for {
			result = append(result, start)
			start += s
			if start >= stop {
				break
			}
		}
	} else if start > stop {
		s = -1
		if len(step) != 0 && step[0] < -1 {
			s = step[0]
		}
		for {
			result = append(result, start)
			start += s
			if start <= stop {
				break
			}
		}
	}
	return result
}

// RangeFloat returns a float64 slice, starting from float64 start, and increments by float64 step,
// and stops before float64 stop.
// If start < stop, default step is 1, else default step is -1.
func RangeFloat(start, stop float64, step ...float64) []float64 {
	return RangeFloat64(start, stop, step...)
}

// RangeFloat64 returns a float64 slice, starting from float64 start, and increments by float64 step,
// and stops before float64 stop.
// If start < stop, default step is 1, else default step is -1.
func RangeFloat64(start, stop float64, step ...float64) []float64 {
	var result []float64
	var s float64
	if start < stop {
		s = 1.0
		if len(step) != 0 && step[0] > 0 {
			s = step[0]
		}
		for {
			result = append(result, start)
			start += s
			if GreaterThanOrEqual(start, stop) {
				break
			}
		}
	} else if start > stop {
		s = -1.0
		if len(step) != 0 && step[0] < 0 {
			s = step[0]
		}
		for {
			result = append(result, start)
			start += s
			if LessThanOrEqual(start, stop) {
				break
			}
		}
	}
	return result
}

// RandInt returns int pseudo-random number in [min, max).
func RandInt(min, max int) int {
	if min > max {
		panic("mathx: min cannot be greater than max")
	}
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}

// RandInt64 returns int64 pseudo-random number in [min, max).
func RandInt64(min, max int64) int64 {
	if min > max {
		panic("mathx: min cannot be greater than max")
	}
	if min == max {
		return min
	}
	return rand.Int63n(max-min) + min
}

// RandFloat returns float64 pseudo-random number in [min, max).
func RandFloat(min, max float64) float64 {
	return RandFloat64(min, max)
}

// RandFloat64 returns float64 pseudo-random number in [min, max).
func RandFloat64(min, max float64) float64 {
	if min > max {
		panic("mathx: min cannot be greater than max")
	}
	return rand.Float64()*(max-min) + min
}

// Max returns the largest float64 number in nums.
func Max(nums ...interface{}) float64 {
	if len(nums) < 1 {
		panic("mathx: nums length cannot be less than 1")
	}
	max := convert.ToFloat64(nums[0])
	for _, num := range nums {
		max = math.Max(max, convert.ToFloat64(num))
	}
	return max
}

// MaxInt returns the largest int number in int nums.
func MaxInt(nums ...int) int {
	if len(nums) < 1 {
		panic("mathx: nums length cannot be less than 1")
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// MaxInt64 returns the largest int64 number in int64 nums.
func MaxInt64(nums ...int64) int64 {
	if len(nums) < 1 {
		panic("mathx: nums length cannot be less than 1")
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// MaxFloat returns the largest float64 number in float64 nums.
func MaxFloat(nums ...float64) float64 {
	return MaxFloat64(nums...)
}

// MaxFloat64 returns the largest float64 number in float64 nums.
func MaxFloat64(nums ...float64) float64 {
	if len(nums) < 1 {
		panic("mathx: nums length cannot be less than 1")
	}
	max := nums[0]
	for _, num := range nums {
		max = math.Max(max, num)
	}
	return max
}

// Min returns the smallest float64 number in nums.
func Min(nums ...interface{}) float64 {
	if len(nums) < 1 {
		panic("mathx: nums length cannot be less than 1")
	}
	min := convert.ToFloat64(nums[0])
	for _, num := range nums {
		min = math.Min(min, convert.ToFloat64(num))
	}
	return min
}

// MinInt returns the smallest int number in int nums.
func MinInt(nums ...int) int {
	if len(nums) < 1 {
		panic("mathx: nums length cannot be less than 1")
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// MinInt64 returns the smallest int64 number in int64 nums.
func MinInt64(nums ...int64) int64 {
	if len(nums) < 1 {
		panic("mathx: nums length cannot be less than 1")
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// MinFloat returns the smallest float64 number in float64 nums.
func MinFloat(nums ...float64) float64 {
	return MinFloat64(nums...)
}

// MinFloat64 returns the smallest float64 number in float64 nums.
func MinFloat64(nums ...float64) float64 {
	if len(nums) < 1 {
		panic("mathx: nums length cannot be less than 1")
	}
	min := nums[0]
	for _, v := range nums {
		min = math.Min(min, v)
	}
	return min
}

// Sum returns the float64 sum of nums.
func Sum(nums ...interface{}) float64 {
	var sum float64
	for _, num := range nums {
		sum += convert.ToFloat64(num)
	}
	return sum
}

// SumInt returns the int sum of int nums.
func SumInt(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

// SumInt64 returns the int64 sum of int64 nums.
func SumInt64(nums ...int64) int64 {
	var sum int64
	for _, num := range nums {
		sum += num
	}
	return sum
}

// SumFloat returns the float64 sum of float64 nums.
func SumFloat(nums ...float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

// SumFloat64 returns the float64 sum of float64 nums.
func SumFloat64(nums ...float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Average returns the float64 average of nums.
func Average(nums ...interface{}) float64 {
	var average float64
	length := len(nums)
	if length > 0 {
		average = Sum(nums...) / float64(length)
	}
	return average
}

// AverageInt returns the float64 average of int nums.
func AverageInt(nums ...int) float64 {
	var average float64
	length := len(nums)
	if length > 0 {
		average = float64(SumInt(nums...)) / float64(length)
	}
	return average
}

// AverageInt64 returns the float64 average of int64 nums.
func AverageInt64(nums ...int64) float64 {
	var average float64
	length := len(nums)
	if length > 0 {
		average = float64(SumInt64(nums...)) / float64(length)
	}
	return average
}

// AverageFloat returns the float64 average of float64 nums.
func AverageFloat(nums ...float64) float64 {
	return AverageFloat64(nums...)
}

// AverageFloat64 returns the float64 average of float64 nums.
func AverageFloat64(nums ...float64) float64 {
	var average float64
	length := len(nums)
	if length > 0 {
		average = SumFloat64(nums...) / float64(length)
	}
	return average
}

// AbsInt returns the absolute int value of int num.
func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// AbsInt64 returns the absolute int64 value of int64 num.
func AbsInt64(num int64) int64 {
	sign := num >> 63
	return (num ^ sign) - sign
}

// AbsFloat returns the absolute float64 value of float64 num.
func AbsFloat(num float64) float64 {
	return AbsFloat64(num)
}

// AbsFloat64 returns the absolute float64 value of float64 num.
func AbsFloat64(num float64) float64 {
	return math.Abs(num)
}

// Percent returns the float64 percentage of num in total.
func Percent(num, total interface{}) float64 {
	t := convert.ToFloat64(total)
	if t == 0 {
		return 0
	}
	n := convert.ToFloat64(num)
	return (n / t) * 100
}

// SizeFormat returns the formatted size with storage unit.
func SizeFormat(size float64, places int, separator ...string) string {
	sep := " "
	if len(separator) != 0 {
		sep = separator[0]
	}
	index := 0
	for {
		if size < 1024 {
			break
		}
		size /= 1024
		index++
	}
	if index >= len(sizeUnits) {
		index = len(sizeUnits) - 1
	}
	return RoundToString(size, places) + sep + sizeUnits[index]
}

// NumberFormat returns the formatted num with grouped thousands.
func NumberFormat(num float64, places int, separator ...string) string {
	isNegative := false
	if num < 0 {
		num = -num
		isNegative = true
	}
	roundStr := RoundToString(num, places)
	var prefix, suffix string
	if places > 0 {
		splits := strings.Split(roundStr, ".")
		prefix = splits[0]
		suffix = splits[1]
	} else {
		prefix = roundStr
	}
	sep := []byte(floatSeparator)
	if len(separator) != 0 {
		sep = []byte(separator[0])
	}
	count, preLength, sepLength := 0, len(prefix), len(sep)
	sepNum := (preLength - 1) / 3
	prefixNew := make([]byte, sepLength*sepNum+preLength) // len(prefixNew) = len(prefix) + sepNum * len(sep)
	prefixNewPos := len(prefixNew) - 1
	for prefixPos := preLength - 1; prefixPos >= 0; prefixPos, count, prefixNewPos = prefixPos-1, count+1, prefixNewPos-1 {
		if sepLength > 0 && count > 0 && count%3 == 0 {
			for j := range sep {
				prefixNew[prefixNewPos] = sep[sepLength-j-1]
				prefixNewPos--
			}
		}
		prefixNew[prefixNewPos] = prefix[prefixPos]
	}
	result := string(prefixNew)
	if places > 0 {
		result += "." + suffix
	}
	if isNegative {
		result = "-" + result
	}
	return result
}
