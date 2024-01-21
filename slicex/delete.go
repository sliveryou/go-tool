package slicex

import (
	"reflect"

	"github.com/sliveryou/go-tool/v2/mathx"
)

// Delete returns the slice that deletes n specified value and the number of actual deletions.
// If n < 0, it will delete all specified value.
func Delete(slice, value interface{}, n int) ([]interface{}, int) {
	if slice == nil {
		return nil, 0
	}
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice, reflect.Array:
		var result []interface{}
		var count int
		v := reflect.ValueOf(slice)
		for i := 0; i < v.Len(); i++ {
			if !reflect.DeepEqual(value, v.Index(i).Interface()) {
				result = append(result, v.Index(i).Interface())
			} else {
				if n < 0 {
					count++
					continue
				}
				if count >= n {
					result = append(result, v.Index(i).Interface())
				} else {
					count++
				}
			}
		}
		return result, count
	default:
		panic("slicex: invalid slice type")
	}
}

// DeleteString returns the slice that deletes n specified string value and the number of actual deletions.
// If n < 0, it will delete all specified string value.
func DeleteString(slice []string, value string, n int) ([]string, int) {
	var result []string
	var count int
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		} else {
			if n < 0 {
				count++
				continue
			}
			if count >= n {
				result = append(result, v)
			} else {
				count++
			}
		}
	}
	return result, count
}

// DeleteBool returns the slice that deletes n specified bool value and the number of actual deletions.
// If n < 0, it will delete all specified bool value.
func DeleteBool(slice []bool, value bool, n int) ([]bool, int) {
	var result []bool
	var count int
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		} else {
			if n < 0 {
				count++
				continue
			}
			if count >= n {
				result = append(result, v)
			} else {
				count++
			}
		}
	}
	return result, count
}

// DeleteInt returns the slice that deletes n specified int value and the number of actual deletions.
// If n < 0, it will delete all specified int value.
func DeleteInt(slice []int, value, n int) ([]int, int) {
	var result []int
	var count int
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		} else {
			if n < 0 {
				count++
				continue
			}
			if count >= n {
				result = append(result, v)
			} else {
				count++
			}
		}
	}
	return result, count
}

// DeleteInt64 returns the slice that deletes n specified int64 value and the number of actual deletions.
// If n < 0, it will delete all specified int64 value.
func DeleteInt64(slice []int64, value int64, n int) ([]int64, int) {
	var result []int64
	var count int
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		} else {
			if n < 0 {
				count++
				continue
			}
			if count >= n {
				result = append(result, v)
			} else {
				count++
			}
		}
	}
	return result, count
}

// DeleteInt32 returns the slice that deletes n specified int32 value and the number of actual deletions.
// If n < 0, it will delete all specified int32 value.
func DeleteInt32(slice []int32, value int32, n int) ([]int32, int) {
	var result []int32
	var count int
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		} else {
			if n < 0 {
				count++
				continue
			}
			if count >= n {
				result = append(result, v)
			} else {
				count++
			}
		}
	}
	return result, count
}

// DeleteFloat returns the slice that deletes n specified float64 value and the number of actual deletions.
// If n < 0, it will delete all specified float64 value.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func DeleteFloat(slice []float64, value float64, n int, places ...int) ([]float64, int) {
	return DeleteFloat64(slice, value, n, places...)
}

// DeleteFloat64 returns the slice that deletes n specified float64 value and the number of actual deletions.
// If n < 0, it will delete all specified float64 value.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func DeleteFloat64(slice []float64, value float64, n int, places ...int) ([]float64, int) {
	var result []float64
	var count int
	for _, v := range slice {
		if !mathx.Equal(v, value, places...) {
			result = append(result, v)
		} else {
			if n < 0 {
				count++
				continue
			}
			if count >= n {
				result = append(result, v)
			} else {
				count++
			}
		}
	}
	return result, count
}

// DeleteFloat32 returns the slice that deletes n specified float32 value and the number of actual deletions.
// If n < 0, it will delete all specified float32 value.
// It will be judged equal if | v1 - v2 | <= 10 ^ -places, default places is 9.
func DeleteFloat32(slice []float32, value float32, n int, places ...int) ([]float32, int) {
	var result []float32
	var count int
	for _, v := range slice {
		if !mathx.Equal(float64(v), float64(value), places...) {
			result = append(result, v)
		} else {
			if n < 0 {
				count++
				continue
			}
			if count >= n {
				result = append(result, v)
			} else {
				count++
			}
		}
	}
	return result, count
}
