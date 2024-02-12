package condition

import (
	"reflect"
)

// Bool returns the bool value of anything.
func Bool[T any](value T) bool {
	switch m := any(value).(type) {
	case interface{ Bool() bool }:
		return m.Bool()
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}

	return reflectValue(&value)
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() != 0
	default:
		return !rv.IsZero()
	}
}

// TernaryOperator checks the value of isTrue, if true return ifValue else return elseValue.
func TernaryOperator[T, U any](isTrue T, ifValue, elseValue U) U {
	if Bool(isTrue) {
		return ifValue
	}

	return elseValue
}
