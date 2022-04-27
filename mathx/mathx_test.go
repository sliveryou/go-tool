package mathx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	assert.Equal(t, 123.46, Round(123.45678, 2))
	assert.Equal(t, 123.456, RoundBank(123.4565, 3))
	assert.Equal(t, "123.46", RoundToString(123.45678, 2))
	assert.Equal(t, "123.456", RoundBankToString(123.4565, 3))
}

func TestSign(t *testing.T) {
	f := 1.23
	assert.Equal(t, 1, Sign(f))
	assert.True(t, IsPositive(f))
	assert.False(t, IsNonPositive(f))

	f = -1.23
	assert.Equal(t, -1, Sign(f))
	assert.True(t, IsNegative(f))
	assert.False(t, IsNonNegative(f))

	f = 0
	assert.Equal(t, 0, Sign(f))
	assert.True(t, IsZero(f))
	assert.True(t, IsNonPositive(f))
	assert.True(t, IsNonNegative(f))
}

func TestCompare(t *testing.T) {
	f1, f2 := 1.1234, 1.1235

	assert.Equal(t, -1, Compare(f1, f2))
	assert.False(t, Equal(f1, f2))
	assert.False(t, GreaterThan(f1, f2))
	assert.True(t, LessThan(f1, f2))
	assert.False(t, GreaterThanOrEqual(f1, f2))
	assert.True(t, LessThanOrEqual(f1, f2))
}

func TestOddAndEven(t *testing.T) {
	assert.True(t, IsOdd(123))
	assert.False(t, IsEven(123))

	assert.False(t, IsOdd(456))
	assert.True(t, IsEven(456))
}

func TestRange(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, RangeInt(0, 10))
	assert.Equal(t, []int{0, 2, 4, 6, 8}, RangeInt(0, 10, 2))
	assert.Equal(t, []int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}, RangeInt(0, -10, -1))
	assert.Empty(t, RangeInt(0, 0))
	assert.Empty(t, RangeInt(1, 1))

	assert.Equal(t, []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, RangeInt64(0, 10))
	assert.Equal(t, []int64{0, 2, 4, 6, 8}, RangeInt64(0, 10, 2))
	assert.Equal(t, []int64{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}, RangeInt64(0, -10, -1))
	assert.Empty(t, RangeInt64(0, 0))
	assert.Empty(t, RangeInt64(1, 1))

	assert.Equal(t, []float64{0.5, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}, RangeFloat(0.5, 10))
	assert.Equal(t, []float64{0.5, 2.5, 4.5, 6.5, 8.5}, RangeFloat(0.5, 10, 2))
	assert.Equal(t, []float64{0.5, -0.5, -1.5, -2.5, -3.5, -4.5, -5.5, -6.5, -7.5, -8.5, -9.5}, RangeFloat(0.5, -10, -1))
	assert.Empty(t, RangeFloat(0, 0))
	assert.Empty(t, RangeFloat(1, 1))

	assert.Equal(t, []float64{0.5, 1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5}, RangeFloat64(0.5, 10))
	assert.Equal(t, []float64{0.5, 2.5, 4.5, 6.5, 8.5}, RangeFloat64(0.5, 10, 2))
	assert.Equal(t, []float64{0.5, -0.5, -1.5, -2.5, -3.5, -4.5, -5.5, -6.5, -7.5, -8.5, -9.5}, RangeFloat64(0.5, -10, -1))
	assert.Empty(t, RangeFloat64(0, 0))
	assert.Empty(t, RangeFloat64(1, 1))
}

func TestRand(t *testing.T) {
	ri := RandInt(5, 100)
	assert.True(t, 5 <= ri && ri < 100)

	ri64 := RandInt64(5, 100)
	assert.True(t, 5 <= ri64 && ri64 < 100)

	rf := RandFloat(5, 100)
	assert.True(t, 5 <= rf && rf < 100)

	rf64 := RandFloat64(5, 100)
	assert.True(t, 5 <= rf64 && rf64 < 100)
}

func TestMax(t *testing.T) {
	assert.Equal(t, 10.1, Max(1, 5.0, 2.0, 3.5, 8, 0, 10.1, 7, 4))
	assert.Equal(t, 9, MaxInt(1, 3, 9, 7, 5))
	assert.Equal(t, int64(9), MaxInt64(1, 3, 9, 7, 5))
	assert.Equal(t, 10.1, MaxFloat(1, 5.0, 2.0, 3.5, 8, 0, 10.1, 7, 4))
	assert.Equal(t, 10.1, MaxFloat64(1, 5.0, 2.0, 3.5, 8, 0, 10.1, 7, 4))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 0.0, Min(1, 5.0, 2.0, 3.5, 8, 0, 10.1, 7, 4))
	assert.Equal(t, 1, MinInt(1, 3, 9, 7, 5))
	assert.Equal(t, int64(1), MinInt64(1, 3, 9, 7, 5))
	assert.Equal(t, 0.0, MinFloat(1, 5.0, 2.0, 3.5, 8, 0, 10.1, 7, 4))
	assert.Equal(t, 0.0, MinFloat64(1, 5.0, 2.0, 3.5, 8, 0, 10.1, 7, 4))
}

func TestSum(t *testing.T) {
	assert.Equal(t, 16.5, Sum(1.1, 2.2, 3.3, 4.4, 5.5))
	assert.Equal(t, 15, SumInt(1, 2, 3, 4, 5))
	assert.Equal(t, int64(15), SumInt64(1, 2, 3, 4, 5))
	assert.Equal(t, 16.5, SumFloat(1.1, 2.2, 3.3, 4.4, 5.5))
	assert.Equal(t, 16.5, SumFloat64(1.1, 2.2, 3.3, 4.4, 5.5))
}

func TestAverage(t *testing.T) {
	assert.Equal(t, 3.3, Average(1.1, 2.2, 3.3, 4.4, 5.5))
	assert.Equal(t, 3.0, AverageInt(1, 2, 3, 4, 5))
	assert.Equal(t, 3.0, AverageInt64(1, 2, 3, 4, 5))
	assert.Equal(t, 3.3, AverageFloat(1.1, 2.2, 3.3, 4.4, 5.5))
	assert.Equal(t, 3.3, AverageFloat64(1.1, 2.2, 3.3, 4.4, 5.5))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 100, AbsInt(-100))
	assert.Equal(t, int64(100), AbsInt64(-100))
	assert.Equal(t, 123.456, AbsFloat(-123.456))
	assert.Equal(t, 123.456, AbsFloat64(-123.456))
	assert.Equal(t, 0.0, AbsFloat64(0))
}

func TestPercent(t *testing.T) {
	assert.Equal(t, 15.0, Percent(15, 100))
	assert.Greater(t, Percent("10", "98"), 10.2)
	assert.Equal(t, 12.34, Percent(12.34, "100"))
	assert.Equal(t, 0.0, Percent(123, "abcd"))
	assert.Equal(t, 0.0, Percent(0, "123"))
}

func TestSizeFormat(t *testing.T) {
	assert.Equal(t, "117.7376 MB", SizeFormat(123456789, 4))
	assert.Equal(t, "11.2283TB", SizeFormat(12345671234567, 4, ""))
	assert.Equal(t, "1.149729 GB", SizeFormat(1234512345, 6, " "))
	assert.Equal(t, "0.000000 B", SizeFormat(0, 6, " "))
}

func TestNumberFormat(t *testing.T) {
	assert.Equal(t, "123,456,789.12346", NumberFormat(123456789.123456789, 5))
	assert.Equal(t, "123-456-789.12346", NumberFormat(123456789.123456789, 5, "-"))
	assert.Equal(t, "123,456,789.12346", NumberFormat(123456789.123456789, 5, ","))
	assert.Equal(t, "0.00000", NumberFormat(0, 5))
}
