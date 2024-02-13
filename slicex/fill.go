package slicex

// Fill returns slice filled with value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func Fill(value interface{}, num int) []interface{} {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}

	result := make([]interface{}, num)
	for i := 0; i < num; i++ {
		result[i] = value
	}

	return result
}

// FillString returns string slice filled with string value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func FillString(value string, num int) []string {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}

	result := make([]string, num)
	for i := 0; i < num; i++ {
		result[i] = value
	}

	return result
}

// FillBool returns bool slice filled with bool value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func FillBool(value bool, num int) []bool {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}

	result := make([]bool, num)
	for i := 0; i < num; i++ {
		result[i] = value
	}

	return result
}

// FillInt returns int slice filled with int value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func FillInt(value, num int) []int {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}

	result := make([]int, num)
	for i := 0; i < num; i++ {
		result[i] = value
	}

	return result
}

// FillInt64 returns int64 slice filled with int64 value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func FillInt64(value int64, num int) []int64 {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}

	result := make([]int64, num)
	for i := 0; i < num; i++ {
		result[i] = value
	}

	return result
}

// FillInt32 returns int32 slice filled with int32 value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func FillInt32(value int32, num int) []int32 {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}

	result := make([]int32, num)
	for i := 0; i < num; i++ {
		result[i] = value
	}

	return result
}

// FillFloat returns float64 slice filled with float64 value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func FillFloat(value float64, num int) []float64 {
	return FillFloat64(value, num)
}

// FillFloat64 returns float64 slice filled with float64 value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func FillFloat64(value float64, num int) []float64 {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}

	result := make([]float64, num)
	for i := 0; i < num; i++ {
		result[i] = value
	}

	return result
}

// FillFloat32 returns float32 slice filled with float32 value,
// where num is the number of value should be filled.
// It panics if num is invalid.
func FillFloat32(value float32, num int) []float32 {
	if num < 0 {
		panic("slicex: num cannot be less than 0")
	}

	result := make([]float32, num)
	for i := 0; i < num; i++ {
		result[i] = value
	}

	return result
}
