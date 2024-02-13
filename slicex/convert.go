package slicex

import (
	"strconv"

	"github.com/sliveryou/go-tool/convert"
)

// StringsToInts returns converted int slice by string slice.
func StringsToInts(slice []string) []int {
	result := make([]int, 0, len(slice))
	for _, s := range slice {
		result = append(result, convert.ToInt(s))
	}

	return result
}

// StringsToInt64s returns converted int64 slice by string slice.
func StringsToInt64s(slice []string) []int64 {
	result := make([]int64, 0, len(slice))
	for _, s := range slice {
		result = append(result, convert.ToInt64(s))
	}

	return result
}

// StringsToFloats returns converted float64 slice by string slice.
func StringsToFloats(slice []string) []float64 {
	return StringsToFloat64s(slice)
}

// StringsToFloat64s returns converted float64 slice by string slice.
func StringsToFloat64s(slice []string) []float64 {
	result := make([]float64, 0, len(slice))
	for _, s := range slice {
		result = append(result, convert.ToFloat64(s))
	}

	return result
}

// StringsToInterfaces returns converted interface slice by string slice.
func StringsToInterfaces(slice []string) []interface{} {
	result := make([]interface{}, 0, len(slice))
	for _, s := range slice {
		result = append(result, s)
	}

	return result
}

// IntsToStrings returns converted string slice by int slice.
func IntsToStrings(slice []int) []string {
	result := make([]string, 0, len(slice))
	for _, i := range slice {
		result = append(result, strconv.Itoa(i))
	}

	return result
}

// IntsToInterfaces returns converted interface slice by int slice.
func IntsToInterfaces(slice []int) []interface{} {
	result := make([]interface{}, 0, len(slice))
	for _, i := range slice {
		result = append(result, i)
	}

	return result
}

// Int64sToStrings returns converted string slice by int64 slice.
func Int64sToStrings(slice []int64) []string {
	result := make([]string, 0, len(slice))
	for _, i := range slice {
		result = append(result, strconv.FormatInt(i, 10))
	}

	return result
}

// Int64sToInterfaces returns converted interface slice by int64 slice.
func Int64sToInterfaces(slice []int64) []interface{} {
	result := make([]interface{}, 0, len(slice))
	for _, i := range slice {
		result = append(result, i)
	}

	return result
}

// FloatsToStrings returns converted string slice by float64 slice.
func FloatsToStrings(slice []float64) []string {
	return Float64sToStrings(slice)
}

// FloatsToInterfaces returns converted interface slice by float64 slice.
func FloatsToInterfaces(slice []float64) []interface{} {
	return Float64sToInterfaces(slice)
}

// Float64sToStrings returns converted string slice by float64 slice.
func Float64sToStrings(slice []float64) []string {
	result := make([]string, 0, len(slice))
	for _, f := range slice {
		result = append(result, strconv.FormatFloat(f, 'f', -1, 64))
	}

	return result
}

// Float64sToInterfaces returns converted interface slice by float64 slice.
func Float64sToInterfaces(slice []float64) []interface{} {
	result := make([]interface{}, 0, len(slice))
	for _, f := range slice {
		result = append(result, f)
	}

	return result
}

// InterfacesToStrings returns converted string slice by interface slice.
func InterfacesToStrings(slice []interface{}) []string {
	result := make([]string, 0, len(slice))
	for _, i := range slice {
		result = append(result, convert.ToString(i))
	}

	return result
}

// InterfacesToInts returns converted int slice by interface slice.
func InterfacesToInts(slice []interface{}) []int {
	result := make([]int, 0, len(slice))
	for _, i := range slice {
		result = append(result, convert.ToInt(i))
	}

	return result
}

// InterfacesToInt64s returns converted int64 slice by interface slice.
func InterfacesToInt64s(slice []interface{}) []int64 {
	result := make([]int64, 0, len(slice))
	for _, i := range slice {
		result = append(result, convert.ToInt64(i))
	}

	return result
}

// InterfacesToFloats returns converted float64 slice by interface slice.
func InterfacesToFloats(slice []interface{}) []float64 {
	return InterfacesToFloat64s(slice)
}

// InterfacesToFloat64s returns converted float64 slice by interface slice.
func InterfacesToFloat64s(slice []interface{}) []float64 {
	result := make([]float64, 0, len(slice))
	for _, i := range slice {
		result = append(result, convert.ToFloat64(i))
	}

	return result
}
