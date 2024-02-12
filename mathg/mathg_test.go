package mathg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	t.Parallel()

	signIntCases := []struct {
		n    int
		want int
	}{
		{n: 100, want: 1},
		{n: 0, want: 0},
		{n: -100, want: -1},
	}

	for _, c := range signIntCases {
		get := Sign(c.n)
		assert.Equal(t, c.want, get)
	}

	signFloatCases := []struct {
		n    float64
		want int
	}{
		{n: 100.1234, want: 1},
		{n: 0, want: 0},
		{n: -100.1234, want: -1},
	}

	for _, c := range signFloatCases {
		get := Sign(c.n)
		assert.Equal(t, c.want, get)
	}
}

func TestSignIs(t *testing.T) {
	t.Parallel()

	testSignIsIntCases := []struct {
		n                 int
		wantIsPositive    bool
		wantIsNonPositive bool
		wantIsNegative    bool
		wantIsNonNegative bool
		wantIsZero        bool
	}{
		{n: 100, wantIsPositive: true, wantIsNonPositive: false, wantIsNegative: false, wantIsNonNegative: true, wantIsZero: false},
		{n: 0, wantIsPositive: false, wantIsNonPositive: true, wantIsNegative: false, wantIsNonNegative: true, wantIsZero: true},
		{n: -100, wantIsPositive: false, wantIsNonPositive: true, wantIsNegative: true, wantIsNonNegative: false, wantIsZero: false},
	}

	for _, c := range testSignIsIntCases {
		assert.Equal(t, c.wantIsPositive, IsPositive(c.n))
		assert.Equal(t, c.wantIsNonPositive, IsNonPositive(c.n))
		assert.Equal(t, c.wantIsNegative, IsNegative(c.n))
		assert.Equal(t, c.wantIsNonNegative, IsNonNegative(c.n))
		assert.Equal(t, c.wantIsZero, IsZero(c.n))
	}

	testSignIsFloatCases := []struct {
		n                 float64
		wantIsPositive    bool
		wantIsNonPositive bool
		wantIsNegative    bool
		wantIsNonNegative bool
		wantIsZero        bool
	}{
		{n: 100.1234, wantIsPositive: true, wantIsNonPositive: false, wantIsNegative: false, wantIsNonNegative: true, wantIsZero: false},
		{n: 0, wantIsPositive: false, wantIsNonPositive: true, wantIsNegative: false, wantIsNonNegative: true, wantIsZero: true},
		{n: -100.1234, wantIsPositive: false, wantIsNonPositive: true, wantIsNegative: true, wantIsNonNegative: false, wantIsZero: false},
	}

	for _, c := range testSignIsFloatCases {
		assert.Equal(t, c.wantIsPositive, IsPositive(c.n))
		assert.Equal(t, c.wantIsNonPositive, IsNonPositive(c.n))
		assert.Equal(t, c.wantIsNegative, IsNegative(c.n))
		assert.Equal(t, c.wantIsNonNegative, IsNonNegative(c.n))
		assert.Equal(t, c.wantIsZero, IsZero(c.n))
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()

	absIntCases := []struct {
		n    int
		want int
	}{
		{n: 100, want: 100},
		{n: 0, want: 0},
		{n: -100, want: 100},
	}

	for _, c := range absIntCases {
		get := Abs(c.n)
		assert.Equal(t, c.want, get)
	}

	absFloatCases := []struct {
		n    float64
		want float64
	}{
		{n: 100.1234, want: 100.1234},
		{n: 0, want: 0},
		{n: -100.1234, want: 100.1234},
	}

	for _, c := range absFloatCases {
		get := Abs(c.n)
		assert.InDelta(t, c.want, get, 0.0001)
	}
}

func TestAverage(t *testing.T) {
	t.Parallel()

	averageIntCases := []struct {
		nums []int
		want float64
	}{
		{nums: []int{1, 3, 5, 10, 2, 0}, want: 3.5},
		{nums: []int{100}, want: 100},
		{nums: []int{}, want: 0},
	}

	for _, c := range averageIntCases {
		assert.InDelta(t, c.want, Average(c.nums...), 0.0001)
	}

	averageFloatCases := []struct {
		nums []float64
		want float64
	}{
		{nums: []float64{1, 3, 5, 10, 2, 0}, want: 3.5},
		{nums: []float64{100}, want: 100},
		{nums: []float64{}, want: 0},
	}

	for _, c := range averageFloatCases {
		assert.InDelta(t, c.want, Average(c.nums...), 0.0001)
	}
}

func TestSum(t *testing.T) {
	t.Parallel()

	sumIntCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 5, 10, 2, 0}, want: 21},
		{nums: []int{100}, want: 100},
		{nums: []int{}, want: 0},
	}

	for _, c := range sumIntCases {
		assert.Equal(t, c.want, Sum(c.nums...))
	}

	sumFloatCases := []struct {
		nums []float64
		want float64
	}{
		{nums: []float64{1, 3, 5, 10, 2, 0}, want: 21},
		{nums: []float64{100}, want: 100},
		{nums: []float64{}, want: 0},
	}

	for _, c := range sumFloatCases {
		assert.InDelta(t, c.want, Sum(c.nums...), 0.0001)
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	maxIntCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 5, 10, 2, 0}, want: 10},
		{nums: []int{100}, want: 100},
	}

	for _, c := range maxIntCases {
		assert.Equal(t, c.want, Max(c.nums...))
	}

	maxFloatCases := []struct {
		nums []float64
		want float64
	}{
		{nums: []float64{1, 3, 5, 10, 2, 0}, want: 10},
		{nums: []float64{100}, want: 100},
	}

	for _, c := range maxFloatCases {
		assert.InDelta(t, c.want, Max(c.nums...), 0.0001)
	}
}

func TestMaxBy(t *testing.T) {
	t.Parallel()

	maxIntCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 5, 10, 2, 0}, want: 10},
		{nums: []int{100}, want: 100},
	}

	for _, c := range maxIntCases {
		assert.Equal(t, c.want, MaxBy(c.nums, func(a, b int) bool {
			return a > b
		}))
	}

	maxFloatCases := []struct {
		nums []float64
		want float64
	}{
		{nums: []float64{1, 3, 5, 10, 2, 0}, want: 10},
		{nums: []float64{100}, want: 100},
	}

	for _, c := range maxFloatCases {
		assert.InDelta(t, c.want, MaxBy(c.nums, func(a, b float64) bool {
			return a > b
		}), 0.0001)
	}
}

func TestMin(t *testing.T) {
	t.Parallel()

	minIntCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 5, 10, 2, 0}, want: 0},
		{nums: []int{100}, want: 100},
	}

	for _, c := range minIntCases {
		assert.Equal(t, c.want, Min(c.nums...))
	}

	minFloatCases := []struct {
		nums []float64
		want float64
	}{
		{nums: []float64{1, 3, 5, 10, 2, 0}, want: 0},
		{nums: []float64{100}, want: 100},
	}

	for _, c := range minFloatCases {
		assert.InDelta(t, c.want, Min(c.nums...), 0.0001)
	}
}

func TestMinBy(t *testing.T) {
	t.Parallel()

	minIntCases := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 3, 5, 10, 2, 0}, want: 0},
		{nums: []int{100}, want: 100},
	}

	for _, c := range minIntCases {
		assert.Equal(t, c.want, MinBy(c.nums, func(a, b int) bool {
			return a < b
		}))
	}

	minFloatCases := []struct {
		nums []float64
		want float64
	}{
		{nums: []float64{1, 3, 5, 10, 2, 0}, want: 0},
		{nums: []float64{100}, want: 100},
	}

	for _, c := range minFloatCases {
		assert.InDelta(t, c.want, MinBy(c.nums, func(a, b float64) bool {
			return a < b
		}), 0.0001)
	}
}

func TestRange(t *testing.T) {
	t.Parallel()

	rangeIntCases := []struct {
		start int
		stop  int
		step  int
		want  []int
	}{
		{start: 1, stop: 10, step: 0, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{start: 1, stop: 10, step: 2, want: []int{1, 3, 5, 7, 9}},
		{start: 10, stop: 1, step: -2, want: []int{10, 8, 6, 4, 2}},
	}

	for _, c := range rangeIntCases {
		get := Range(c.start, c.stop, c.step)
		assert.Equal(t, c.want, get)
	}

	rangeFloatCases := []struct {
		start float64
		stop  float64
		step  float64
		want  []float64
	}{
		{start: 1, stop: 10, step: 0, want: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{start: 1, stop: 10, step: 2, want: []float64{1, 3, 5, 7, 9}},
		{start: 10, stop: 1, step: -2, want: []float64{10, 8, 6, 4, 2}},
	}

	for _, c := range rangeFloatCases {
		get := Range(c.start, c.stop, c.step)
		assert.Equal(t, c.want, get)
	}
}

func TestSignedNumberIs(t *testing.T) {
	t.Parallel()

	signedNumberIsIntCases := []struct {
		n          int
		wantIsOdd  bool
		wantIsEven bool
	}{
		{n: 100, wantIsOdd: false, wantIsEven: true},
		{n: 0, wantIsOdd: false, wantIsEven: true},
		{n: -99, wantIsOdd: true, wantIsEven: false},
	}

	for _, c := range signedNumberIsIntCases {
		assert.Equal(t, c.wantIsOdd, IsOdd(c.n))
		assert.Equal(t, c.wantIsEven, IsEven(c.n))
	}

	signedNumberIsInt64Cases := []struct {
		n          int64
		wantIsOdd  bool
		wantIsEven bool
	}{
		{n: 100, wantIsOdd: false, wantIsEven: true},
		{n: 0, wantIsOdd: false, wantIsEven: true},
		{n: -99, wantIsOdd: true, wantIsEven: false},
	}

	for _, c := range signedNumberIsInt64Cases {
		assert.Equal(t, c.wantIsOdd, IsOdd(c.n))
		assert.Equal(t, c.wantIsEven, IsEven(c.n))
	}
}

func TestDivCeil(t *testing.T) {
	t.Parallel()

	int64Cases := []struct {
		x      int64
		y      int64
		expect int64
	}{
		{x: 100, y: 24, expect: 5},
		{x: 100, y: -24, expect: -4},
		{x: 90, y: 23, expect: 4},
		{x: 80, y: 22, expect: 4},
		{x: 70, y: 0, expect: 0},
		{x: 0, y: 70, expect: 0},
	}

	for _, c := range int64Cases {
		out := DivCeil(c.x, c.y)
		assert.Equal(t, c.expect, out)
	}
}

func TestDivFloor(t *testing.T) {
	t.Parallel()

	int64Cases := []struct {
		x      int64
		y      int64
		expect int64
	}{
		{x: 100, y: 24, expect: 4},
		{x: 100, y: -24, expect: -5},
		{x: 90, y: 23, expect: 3},
		{x: 80, y: 22, expect: 3},
		{x: 70, y: 0, expect: 0},
		{x: 0, y: 70, expect: 0},
	}

	for _, c := range int64Cases {
		out := DivFloor(c.x, c.y)
		assert.Equal(t, c.expect, out)
	}
}

func TestDivRound(t *testing.T) {
	t.Parallel()

	int64Cases := []struct {
		x      int64
		y      int64
		expect int64
	}{
		{x: 100, y: 24, expect: 4},
		{x: 100, y: -24, expect: -4},
		{x: 90, y: 23, expect: 4},
		{x: 80, y: 22, expect: 4},
		{x: 70, y: 0, expect: 0},
		{x: 0, y: 70, expect: 0},
	}

	for _, c := range int64Cases {
		out := DivRound(c.x, c.y)
		assert.Equal(t, c.expect, out)
	}
}

func TestMod(t *testing.T) {
	t.Parallel()

	int64Cases := []struct {
		x      int64
		y      int64
		expect int64
	}{
		{x: 100, y: 24, expect: 4},
		{x: 100, y: -24, expect: 4},
		{x: -100, y: 24, expect: -4},
		{x: 90, y: 23, expect: 21},
		{x: 80, y: 22, expect: 14},
		{x: 70, y: 0, expect: 0},
		{x: 0, y: 70, expect: 0},
	}

	for _, c := range int64Cases {
		out := Mod(c.x, c.y)
		assert.Equal(t, c.expect, out)
	}
}

func TestPow(t *testing.T) {
	t.Parallel()

	int64Cases := []struct {
		x      int64
		y      int64
		expect int64
	}{
		{x: 1, y: 0, expect: 1},
		{x: 1, y: 3, expect: 1},
		{x: 2, y: 0, expect: 1},
		{x: 2, y: 1, expect: 2},
		{x: 2, y: 3, expect: 8},
		{x: 5, y: 0, expect: 1},
		{x: 5, y: 1, expect: 5},
		{x: 5, y: 3, expect: 125},
	}

	for _, c := range int64Cases {
		out := Pow(c.x, c.y)
		assert.Equal(t, c.expect, out)
	}
}

func TestDim(t *testing.T) {
	t.Parallel()

	int64Cases := []struct {
		x      int64
		y      int64
		expect int64
	}{
		{x: 1, y: 0, expect: 1},
		{x: 1, y: 3, expect: 0},
		{x: 2, y: 0, expect: 2},
		{x: 2, y: 1, expect: 1},
		{x: 2, y: 3, expect: 0},
		{x: 5, y: 0, expect: 5},
		{x: 5, y: 1, expect: 4},
		{x: 5, y: 8, expect: 0},
	}

	for _, c := range int64Cases {
		out := Dim(c.x, c.y)
		assert.Equal(t, c.expect, out)
	}
}
