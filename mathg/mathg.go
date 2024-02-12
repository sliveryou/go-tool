package mathg

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Sign returns:
//
//	-1 if n <  0
//	 0 if n == 0
//	+1 if n >  0
func Sign[T constraints.Integer | constraints.Float](n T) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}

	return 0
}

// IsPositive returns:
//
//	 true if n >  0
//	false if n == 0
//	false if n <  0
func IsPositive[T constraints.Integer | constraints.Float](n T) bool {
	return Sign(n) == 1
}

// IsNonPositive returns:
//
//	 true if n <  0
//	 true if n == 0
//	false if n >  0
func IsNonPositive[T constraints.Integer | constraints.Float](n T) bool {
	return Sign(n) != 1
}

// IsNegative returns:
//
//	 true if n <  0
//	false if n == 0
//	false if n >  0
func IsNegative[T constraints.Integer | constraints.Float](n T) bool {
	return Sign(n) == -1
}

// IsNonNegative returns:
//
//	 true if n >  0
//	 true if n == 0
//	false if n <  0
func IsNonNegative[T constraints.Integer | constraints.Float](n T) bool {
	return Sign(n) != -1
}

// IsZero returns:
//
//	 true if n == 0
//	false if n >  0
//	false if n <  0
func IsZero[T constraints.Integer | constraints.Float](n T) bool {
	return Sign(n) == 0
}

// Abs returns the absolute int value of n.
func Abs[T constraints.Integer | constraints.Float](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

// Average returns the average of nums.
func Average[T constraints.Integer | constraints.Float](nums ...T) float64 {
	var average float64
	length := len(nums)
	if length > 0 {
		average = float64(Sum(nums...)) / float64(length)
	}

	return average
}

// Sum returns the sum of nums.
func Sum[T constraints.Ordered](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}

	return sum
}

// Max returns the largest element in nums.
func Max[T constraints.Ordered](nums ...T) T {
	var max T
	if len(nums) == 0 {
		return max
	}

	max = nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > max {
			max = num
		}
	}

	return max
}

// MaxBy return the largest element in slice using the given cmp function.
func MaxBy[T any](slice []T, cmp func(a, b T) bool) T {
	var max T
	if len(slice) == 0 {
		return max
	}

	max = slice[0]
	for i := 1; i < len(slice); i++ {
		val := slice[i]
		if cmp(val, max) {
			max = val
		}
	}

	return max
}

// Min returns the smallest element in nums.
func Min[T constraints.Ordered](nums ...T) T {
	var min T
	if len(nums) == 0 {
		return min
	}

	min = nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < min {
			min = num
		}
	}

	return min
}

// MinBy return the smallest element in slice using the given cmp function.
func MinBy[T any](slice []T, cmp func(a, b T) bool) T {
	var min T
	if len(slice) == 0 {
		return min
	}

	min = slice[0]
	for i := 1; i < len(slice); i++ {
		val := slice[i]
		if cmp(val, min) {
			min = val
		}
	}

	return min
}

// Range returns a slice, starting from start, and increments by step, and stops before stop.
// If start < stop, default step is 1, else default step is -1.
func Range[T constraints.Signed | constraints.Float](start, stop T, step ...T) []T {
	var s T
	var rs []T

	if start < stop {
		s = 1
		if len(step) != 0 && step[0] > 1 {
			s = step[0]
		}
		for {
			rs = append(rs, start)
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
			rs = append(rs, start)
			start += s
			if start <= stop {
				break
			}
		}
	}

	return rs
}

// IsOdd reports whether n is odd number.
func IsOdd[T constraints.Integer](n T) bool {
	return n%2 != 0
}

// IsEven reports whether n is even number.
func IsEven[T constraints.Integer](n T) bool {
	return n%2 == 0
}

// DivCeil returns the least integer value greater than or equal to x/y.
// If y == 0, the result will be 0.
func DivCeil[T constraints.Integer | constraints.Float](x, y T) T {
	if y == 0 {
		return 0
	}

	return T(math.Ceil(float64(x) / float64(y)))
}

// DivFloor returns the greatest integer value less than or equal to x/y.
// If y == 0, the result will be 0.
func DivFloor[T constraints.Integer | constraints.Float](x, y T) T {
	if y == 0 {
		return 0
	}

	return T(math.Floor(float64(x) / float64(y)))
}

// DivRound returns the nearest integer, rounding half away from zero to x/y.
// If y == 0, the result will be 0.
func DivRound[T constraints.Integer | constraints.Float](x, y T) T {
	if y == 0 {
		return 0
	}

	return T(math.Round(float64(x) / float64(y)))
}

// Mod returns the remainder of x/y.
// The magnitude of the result is less than y and its
// sign agrees with that of x.
// If y == 0, the result will be 0.
func Mod[T constraints.Integer | constraints.Float](x, y T) T {
	if y == 0 {
		return 0
	}

	return T(math.Mod(float64(x), float64(y)))
}

// Pow returns x**y, the base-x exponential of y.
func Pow[T constraints.Integer | constraints.Float](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

// Dim returns the maximum of x-y or 0.
func Dim[T constraints.Integer | constraints.Float](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}
