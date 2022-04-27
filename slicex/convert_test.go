package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsToInts(t *testing.T) {
	arr1 := []int{123, 456, 789, 123, 0}
	arr2 := []string{"123.123", " 456 ", " 789.789 ", "0000123", "0000.1234"}

	assert.Equal(t, arr1, StringsToInts(arr2))
	assert.Equal(t, []int{0, 0, 0}, StringsToInts([]string{"abcd", "!@#$%^&*", "0.123"}))
	assert.Empty(t, StringsToInts([]string{}))
	assert.Empty(t, StringsToInts(nil))
}

func TestStringsToInt64s(t *testing.T) {
	arr1 := []int64{123, 456, 789, 123, 0}
	arr2 := []string{"123.123", " 456 ", " 789.789 ", "0000123", "0000.1234"}

	assert.Equal(t, arr1, StringsToInt64s(arr2))
	assert.Equal(t, []int64{0, 0, 0}, StringsToInt64s([]string{"abcd", "!@#$%^&*", "0.123"}))
	assert.Empty(t, StringsToInt64s([]string{}))
	assert.Empty(t, StringsToInt64s(nil))
}

func TestStringsToFloats(t *testing.T) {
	arr1 := []float64{123.123, 456.456, 789.789, 0.1234}
	arr2 := []string{"123.123", "  456.456  ", " 789.789", "0000.1234"}

	assert.Equal(t, arr1, StringsToFloats(arr2))
	assert.Equal(t, []float64{0, 0, 0}, StringsToFloats([]string{"abcd", "!@#$%^&*", "0.1.123"}))
	assert.Empty(t, StringsToFloats([]string{}))
	assert.Empty(t, StringsToFloats(nil))
}

func TestStringsToFloat64s(t *testing.T) {
	arr1 := []float64{123.123, 456.456, 789.789, 0.1234}
	arr2 := []string{"123.123", "  456.456  ", " 789.789", "0000.1234"}

	assert.Equal(t, arr1, StringsToFloat64s(arr2))
	assert.Equal(t, []float64{0, 0, 0}, StringsToFloat64s([]string{"abcd", "!@#$%^&*", "0.1.123"}))
	assert.Empty(t, StringsToFloat64s([]string{}))
	assert.Empty(t, StringsToFloat64s(nil))
}

func TestStringsToInterfaces(t *testing.T) {
	arr1 := []interface{}{"123.123", "456.456", "789.789"}
	arr2 := []string{"123.123", "456.456", "789.789"}

	assert.Equal(t, arr1, StringsToInterfaces(arr2))
	assert.Empty(t, StringsToInterfaces([]string{}))
	assert.Empty(t, StringsToInterfaces(nil))
}

func TestIntsToStrings(t *testing.T) {
	arr1 := []string{"123", "456", "789"}
	arr2 := []int{123, 456, 789}

	assert.Equal(t, arr1, IntsToStrings(arr2))
	assert.Empty(t, IntsToStrings([]int{}))
	assert.Empty(t, IntsToStrings(nil))
}

func TestIntsToInterfaces(t *testing.T) {
	arr1 := []interface{}{123, 456, 789}
	arr2 := []int{123, 456, 789}

	assert.Equal(t, arr1, IntsToInterfaces(arr2))
	assert.Empty(t, IntsToInterfaces([]int{}))
	assert.Empty(t, IntsToInterfaces(nil))
}

func TestInt64sToStrings(t *testing.T) {
	arr1 := []string{"123", "456", "789"}
	arr2 := []int64{123, 456, 789}

	assert.Equal(t, arr1, Int64sToStrings(arr2))
	assert.Empty(t, Int64sToStrings([]int64{}))
	assert.Empty(t, Int64sToStrings(nil))
}

func TestInt64sToInterfaces(t *testing.T) {
	arr1 := []interface{}{int64(123), int64(456), int64(789)}
	arr2 := []int64{123, 456, 789}

	assert.Equal(t, arr1, Int64sToInterfaces(arr2))
	assert.Empty(t, Int64sToInterfaces([]int64{}))
	assert.Empty(t, Int64sToInterfaces(nil))
}

func TestFloatsToStrings(t *testing.T) {
	arr1 := []string{"123.123", "456.456", "789.789"}
	arr2 := []float64{123.123, 456.456, 789.789}

	assert.Equal(t, arr1, FloatsToStrings(arr2))
	assert.Empty(t, FloatsToStrings([]float64{}))
	assert.Empty(t, FloatsToStrings(nil))
}

func TestFloatsToInterfaces(t *testing.T) {
	arr1 := []interface{}{123.123, 456.456, 789.789}
	arr2 := []float64{123.123, 456.456, 789.789}

	assert.Equal(t, arr1, FloatsToInterfaces(arr2))
	assert.Empty(t, FloatsToInterfaces([]float64{}))
	assert.Empty(t, FloatsToInterfaces(nil))
}

func TestFloat64sToStrings(t *testing.T) {
	arr1 := []float64{123.123, 456.456, 789.789}
	arr2 := []string{"123.123", "456.456", "789.789"}

	assert.Equal(t, arr2, Float64sToStrings(arr1))
	assert.Empty(t, Float64sToStrings([]float64{}))
	assert.Empty(t, Float64sToStrings(nil))
}

func TestFloat64sToInterfaces(t *testing.T) {
	arr1 := []interface{}{123.123, 456.456, 789.789}
	arr2 := []float64{123.123, 456.456, 789.789}

	assert.Equal(t, arr1, Float64sToInterfaces(arr2))
	assert.Empty(t, Float64sToInterfaces([]float64{}))
	assert.Empty(t, Float64sToInterfaces(nil))
}

func TestInterfacesToStrings(t *testing.T) {
	arr1 := []string{"123.123", "456", "789.789", "true", "false"}
	arr2 := []interface{}{123.123, 456, "789.789", true, false}

	assert.Equal(t, arr1, InterfacesToStrings(arr2))
	assert.Empty(t, InterfacesToStrings([]interface{}{}))
	assert.Empty(t, InterfacesToStrings(nil))
}

func TestInterfacesToInts(t *testing.T) {
	arr1 := []int{123, 456, 789, 1, 0}
	arr2 := []interface{}{123.123, 456, " 789.789 ", true, false}

	assert.Equal(t, arr1, InterfacesToInts(arr2))
	assert.Equal(t, []int{0, 0, 0}, InterfacesToInts([]interface{}{"abcd", "!@#$%^&*", "0.123"}))
	assert.Empty(t, InterfacesToInts([]interface{}{}))
	assert.Empty(t, InterfacesToInts(nil))
}

func TestInterfacesToInt64s(t *testing.T) {
	arr1 := []int64{123, 456, 789, 1, 0}
	arr2 := []interface{}{123.123, 456, "789.789", true, false}

	assert.Equal(t, arr1, InterfacesToInt64s(arr2))
	assert.Equal(t, []int64{0, 0, 0}, InterfacesToInt64s([]interface{}{"abcd", "!@#$%^&*", "0.123"}))
	assert.Empty(t, InterfacesToInt64s([]interface{}{}))
	assert.Empty(t, InterfacesToInt64s(nil))
}

func TestInterfacesToFloats(t *testing.T) {
	arr1 := []float64{123.123, 456, 789.789, 1, 0}
	arr2 := []interface{}{123.123, 456, " 789.789 ", true, false}

	assert.Equal(t, arr1, InterfacesToFloats(arr2))
	assert.Equal(t, []float64{0, 0, 0}, InterfacesToFloats([]interface{}{"abcd", "!@#$%^&*", "0.1.123"}))
	assert.Empty(t, InterfacesToFloats([]interface{}{}))
	assert.Empty(t, InterfacesToFloats(nil))
}

func TestInterfacesToFloat64s(t *testing.T) {
	arr1 := []float64{123.123, 456, 789.789, 1, 0}
	arr2 := []interface{}{123.123, 456, " 789.789 ", true, false}

	assert.Equal(t, arr1, InterfacesToFloat64s(arr2))
	assert.Equal(t, []float64{0, 0, 0}, InterfacesToFloat64s([]interface{}{"abcd", "!@#$%^&*", "0.1.123"}))
	assert.Empty(t, InterfacesToFloat64s([]interface{}{}))
	assert.Empty(t, InterfacesToFloat64s(nil))
}
