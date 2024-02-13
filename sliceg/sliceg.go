package sliceg

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"

	"github.com/sliveryou/go-tool/v2/sliceg/internal/slices"
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
// or -1 if none do.
func IndexFunc[T any](s []T, f func(v T) bool) int {
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

// ContainFunc reports whether at least one
// element v of s satisfies f(v).
func ContainFunc[T any](s []T, f func(v T) bool) bool {
	return IndexFunc(s, f) >= 0
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
func DeleteFunc[T any](s []T, f func(v T) bool, n int) ([]T, int) {
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
func EqualFunc[T1, T2 any](s1 []T1, s2 []T2, eq func(v1 T1, v2 T2) bool) bool {
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
		return []T{}
	}
	if n > l {
		n = l
	}
	es := make([]T, n)
	for i, v := range rand.Perm(l) {
		if i >= n {
			break
		}
		es[i] = s[v]
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

// ReverseSelf reverses the elements of the slice in place.
func ReverseSelf[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Shuffle returns shuffled s,
// it is equivalent to Extract(s, len(s)).
func Shuffle[T any](s []T) []T {
	return Extract(s, len(s))
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

// UniqueFunc returns the unique s.
func UniqueFunc[T any, U comparable](s []T, f func(v T) U) []T {
	us := make([]T, 0, len(s))
	um := make(map[U]struct{}, len(s))
	for _, v := range s {
		fv := f(v)
		if _, ok := um[fv]; !ok {
			um[fv] = struct{}{}
			us = append(us, v)
		}
	}

	return us
}

// Subset reports whether the specified slice contains all
// elements given in the specified subset.
func Subset[T comparable](slice, subset []T) bool {
	for _, v := range subset {
		if !Contain(slice, v) {
			return false
		}
	}

	return true
}

// SubsetFunc reports whether the specified slice contains all
// elements given in the specified subset using the cmp function.
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b.
func SubsetFunc[T any](slice, subset []T, cmp func(a, b T) int) bool {
	if len(subset) == 0 {
		return true
	}

	clone := Clone(slice)
	SortFunc(clone, cmp)

	for _, a := range subset {
		if _, find := BinarySearchFunc(clone, a, cmp); !find {
			return false
		}
	}

	return true
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[T any](s []T, needInit ...bool) []T {
	if s == nil {
		if len(needInit) > 0 && needInit[0] {
			return []T{}
		}
		return nil
	}

	return append([]T{}, s...)
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func Clip[T any](s []T) []T {
	return s[:len(s):len(s)]
}

// Sort sorts a slice of any ordered type in ascending order.
// When sorting floating-point numbers, NaNs are ordered before other values.
func Sort[T constraints.Ordered](x []T) {
	slices.Sort(x)
}

// SortFunc sorts the slice x in ascending order as determined by the cmp
// function. This sort is not guaranteed to be stable.
// cmp(a, b) should return a negative number when a < b, a positive number when
// a > b and zero when a == b.
//
// SortFunc requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func SortFunc[T any](x []T, cmp func(a, b T) int) {
	slices.SortFunc(x, cmp)
}

// SortStableFunc sorts the slice x while keeping the original order of equal
// elements, using cmp to compare elements in the same way as [SortFunc].
func SortStableFunc[T any](x []T, cmp func(a, b T) int) {
	slices.SortStableFunc(x, cmp)
}

// IsSorted reports whether x is sorted in ascending order.
func IsSorted[T constraints.Ordered](x []T) bool {
	return slices.IsSorted(x)
}

// IsSortedFunc reports whether x is sorted in ascending order, with cmp as the
// comparison function as defined by [SortFunc].
func IsSortedFunc[T any](x []T, cmp func(a, b T) int) bool {
	return slices.IsSortedFunc(x, cmp)
}

// Min returns the minimal value in x. It panics if x is empty.
// For floating-point numbers, Min propagates NaNs (any NaN value in x
// forces the output to be NaN).
func Min[T constraints.Ordered](x []T) T {
	return slices.Min(x)
}

// MinFunc returns the minimal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one minimal element
// according to the cmp function, MinFunc returns the first one.
func MinFunc[T any](x []T, cmp func(a, b T) int) T {
	return slices.MinFunc(x, cmp)
}

// Max returns the maximal value in x. It panics if x is empty.
// For floating-point E, Max propagates NaNs (any NaN value in x
// forces the output to be NaN).
func Max[T constraints.Ordered](x []T) T {
	return slices.Max(x)
}

// MaxFunc returns the maximal value in x, using cmp to compare elements.
// It panics if x is empty. If there is more than one maximal element
// according to the cmp function, MaxFunc returns the first one.
func MaxFunc[T any](x []T, cmp func(a, b T) int) T {
	return slices.MaxFunc(x, cmp)
}

// BinarySearch searches for target in a sorted slice and returns the position
// where target is found, or the position where target would appear in the
// sort order; it also returns a bool saying whether the target is really found
// in the slice. The slice must be sorted in increasing order.
func BinarySearch[T constraints.Ordered](x []T, target T) (int, bool) {
	return slices.BinarySearch(x, target)
}

// BinarySearchFunc works like [BinarySearch], but uses a custom comparison
// function. The slice must be sorted in increasing order, where "increasing"
// is defined by cmp. cmp should return 0 if the slice element matches
// the target, a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target.
// cmp must implement the same ordering as the slice, such that if
// cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.
func BinarySearchFunc[E, T any](x []E, target T, cmp func(a E, b T) int) (int, bool) {
	return slices.BinarySearchFunc(x, target, cmp)
}
