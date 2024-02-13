package pointer

import "reflect"

// Of returns the pointer to the value v.
func Of[T any](v T) *T {
	return &v
}

// Unwrap returns the value from the pointer p.
func Unwrap[T any](p *T) T {
	return *p
}

// UnwrapOr returns the value from the pointer p or fallback if p is nil.
func UnwrapOr[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}

	return *p
}

// UnwrapOrDefault returns the value from the pointer p or the default value if p is nil.
func UnwrapOrDefault[T any](p *T) T {
	var v T
	if p == nil {
		return v
	}

	return *p
}

// ExtractPointer returns the underlying value by the given any type v.
func ExtractPointer(v any) any {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer {
		return v
	}

	val = val.Elem()
	if val.IsValid() {
		return ExtractPointer(val.Interface())
	}

	return nil
}
