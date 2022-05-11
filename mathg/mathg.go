package mathg

import (
	"golang.org/x/exp/constraints"
)

// Sign returns:
//     -1 if n <  0
//      0 if n == 0
//     +1 if n >  0
func Sign[T constraints.Signed | constraints.Float](n T) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}

// IsPositive returns:
//      true if n >  0
//     false if n == 0
//     false if n <  0
func IsPositive[T constraints.Signed | constraints.Float](n T) bool {
	return Sign(n) == 1
}

// IsNonPositive returns:
//      true if n <  0
//      true if n == 0
//     false if n >  0
func IsNonPositive[T constraints.Signed | constraints.Float](n T) bool {
	return Sign(n) != 1
}

// IsNegative returns:
//      true if n <  0
//     false if n == 0
//     false if n >  0
func IsNegative[T constraints.Signed | constraints.Float](n T) bool {
	return Sign(n) == -1
}

// IsNonNegative returns:
//      true if n >  0
//      true if n == 0
//     false if n <  0
func IsNonNegative[T constraints.Signed | constraints.Float](n T) bool {
	return Sign(n) != -1
}

// IsZero returns:
//      true if n == 0
//     false if n >  0
//     false if n <  0
func IsZero[T constraints.Signed | constraints.Float](n T) bool {
	return Sign(n) == 0
}

// Abs returns the absolute int value of n.
func Abs[T constraints.Signed | constraints.Float](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// Average returns the average of nums.
func Average[T constraints.Signed | constraints.Float](nums ...T) float64 {
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
	if len(nums) < 1 {
		panic("mathg: nums length cannot be less than 1")
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// Min returns the smallest element in nums.
func Min[T constraints.Ordered](nums ...T) T {
	if len(nums) < 1 {
		panic("mathg: nums length cannot be less than 1")
	}
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
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
func IsOdd[T constraints.Signed](n T) bool {
	return n%2 != 0
}

// IsEven reports whether n is even number.
func IsEven[T constraints.Signed](n T) bool {
	return n%2 == 0
}
