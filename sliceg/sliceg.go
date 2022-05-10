package sliceg

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Index returns the index of the first occurrence of v in s,
// or -1 if v not present in s.
func Index[T comparable](s []T, v T) int {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return i
		}
	}

	return -1
}

// IndexFunc returns the index of the first index i satisfying f(s[i]),
// or -1 if v not present in s.
func IndexFunc[T any](s []T, f func(T) bool) int {
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			return i
		}
	}

	return -1
}

// Contain reports whether v is present in s.
func Contain[T comparable](s []T, v T) bool {
	return Index(s, v) >= 0
}

// Count returns value count map by s.
func Count[T comparable](s []T) map[T]int {
	m := make(map[T]int, len(s))
	for _, v := range s {
		m[v]++
	}

	return m
}

// Delete returns the slice that deletes n specified v and the number of actual deletions.
// If n < 0, it will delete all specified value.
func Delete[T comparable](s []T, v T, n int) ([]T, int) {
	var ds []T
	var c int
	for _, e := range s {
		if e != v {
			ds = append(ds, e)
		} else {
			if n < 0 {
				c++
				continue
			}
			if c >= n {
				ds = append(ds, e)
			} else {
				c++
			}
		}
	}

	return ds, c
}

// DeleteFunc returns the slice that deletes n satisfying f(s[i]) and the number of actual deletions.
// If n < 0, it will delete all satisfying f(s[i]).
func DeleteFunc[T any](s []T, f func(T) bool, n int) ([]T, int) {
	var ds []T
	var c int
	for i := range s {
		if !f(s[i]) {
			ds = append(ds, s[i])
		} else {
			if n < 0 {
				c++
				continue
			}
			if c >= n {
				ds = append(ds, s[i])
			} else {
				c++
			}
		}
	}

	return ds, c
}

// Equal reports whether s1 equals s2.
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	if (s1 == nil) != (s2 == nil) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// EqualFunc reports whether s1 equals s2 which using a comparison
// function on each pair of elements.
func EqualFunc[T1, T2 any](s1 []T1, s2 []T2, eq func(T1, T2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	if (s1 == nil) != (s2 == nil) {
		return false
	}
	for i := range s1 {
		if !eq(s1[i], s2[i]) {
			return false
		}
	}

	return true
}

// Extract returns randomly extracted n elements from s.
// It panics if n is invalid.
func Extract[T any](s []T, n int) []T {
	if n < 0 {
		panic("sliceg: n cannot be less than 0")
	}
	l := len(s)
	if l == 0 {
		return nil
	}
	if n > l {
		n = l
	}
	es := make([]T, n)
	for i, v := range rand.Perm(l) {
		if i < n {
			es[i] = s[v]
		} else {
			break
		}
	}

	return es
}

// Fill returns slice filled with v,
// where n is the number of v should be filled.
func Fill[T any](v T, n int) []T {
	if n < 0 {
		panic("sliceg: n cannot be less than 0")
	}
	fs := make([]T, n)
	for i := 0; i < n; i++ {
		fs[i] = v
	}

	return fs
}

// Reverse returns the reverse order for s.
func Reverse[T any](s []T) []T {
	l := len(s)
	rs := make([]T, l)
	i, j := 0, l-1
	for ; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = s[j], s[i]
	}
	if l > 0 && l%2 == 1 {
		rs[j] = s[j]
	}

	return rs
}

// Shuffle returns shuffled s,
// it is equivalent to Extract(s, len(s)).
func Shuffle[T any](s []T) []T {
	length := len(s)
	if length == 0 {
		return nil
	}

	return Extract(s, length)
}

// Take returns the randomly taken element from s,
// it is equivalent to Extract(s, 1)[0].
func Take[T any](s []T) T {
	if len(s) == 0 {
		var t T
		return t
	}

	return Extract(s, 1)[0]
}

// Unique returns the unique s.
func Unique[T comparable](s []T) []T {
	us := make([]T, 0, len(s))
	um := make(map[T]struct{}, len(s))
	for _, v := range s {
		if _, ok := um[v]; !ok {
			um[v] = struct{}{}
			us = append(us, v)
		}
	}

	return us
}
