package sliceg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	indexIntCases = []struct {
		s    []int
		v    int
		want int
	}{
		{s: []int{1, 2, 3}, v: 1, want: 0},
		{s: []int{1, 2, 3}, v: 2, want: 1},
		{s: []int{1, 2, 3}, v: 4, want: -1},
		{s: []int{}, v: 1, want: -1},
		{s: nil, v: 1, want: -1},
	}

	indexFloatCases = []struct {
		s    []float64
		v    float64
		want int
	}{
		{s: []float64{1.1, 2.2, 3.3}, v: 1.1, want: 0},
		{s: []float64{1.1, 2.2, 3.3}, v: 2.2, want: 1},
		{s: []float64{1.1, 2.2, 3.3}, v: 4.4, want: -1},
		{s: []float64{}, v: 1.1, want: -1},
		{s: nil, v: 1, want: -1},
	}

	indexStringCases = []struct {
		s    []string
		v    string
		want int
	}{
		{s: []string{"1", "2", "3"}, v: "1", want: 0},
		{s: []string{"1", "2", "3"}, v: "2", want: 1},
		{s: []string{"1", "2", "3"}, v: "4", want: -1},
		{s: []string{}, v: "1", want: -1},
		{s: nil, v: "1", want: -1},
	}
)

func TestIndex(t *testing.T) {
	for _, c := range indexIntCases {
		get := Index(c.s, c.v)
		assert.Equal(t, c.want, get)
	}
	for _, c := range indexFloatCases {
		get := Index(c.s, c.v)
		assert.Equal(t, c.want, get)
	}
	for _, c := range indexStringCases {
		get := Index(c.s, c.v)
		assert.Equal(t, c.want, get)
	}
}

func TestIndexFunc(t *testing.T) {
	for _, c := range indexIntCases {
		get := IndexFunc(c.s, func(v int) bool { return c.v == v })
		assert.Equal(t, c.want, get)
	}
	for _, c := range indexFloatCases {
		get := IndexFunc(c.s, func(v float64) bool { return c.v == v })
		assert.Equal(t, c.want, get)
	}
	for _, c := range indexStringCases {
		get := IndexFunc(c.s, func(v string) bool { return c.v == v })
		assert.Equal(t, c.want, get)
	}
}

var (
	containIntCases = []struct {
		s    []int
		v    int
		want bool
	}{
		{s: []int{1, 2, 3}, v: 1, want: true},
		{s: []int{1, 2, 3}, v: 2, want: true},
		{s: []int{1, 2, 3}, v: 4, want: false},
		{s: []int{}, v: 1, want: false},
		{s: nil, v: 1, want: false},
	}

	containFloatCases = []struct {
		s    []float64
		v    float64
		want bool
	}{
		{s: []float64{1.1, 2.2, 3.3}, v: 1.1, want: true},
		{s: []float64{1.1, 2.2, 3.3}, v: 2.2, want: true},
		{s: []float64{1.1, 2.2, 3.3}, v: 4.4, want: false},
		{s: []float64{}, v: 1.1, want: false},
		{s: nil, v: 1, want: false},
	}

	containStringCases = []struct {
		s    []string
		v    string
		want bool
	}{
		{s: []string{"1", "2", "3"}, v: "1", want: true},
		{s: []string{"1", "2", "3"}, v: "2", want: true},
		{s: []string{"1", "2", "3"}, v: "4", want: false},
		{s: []string{}, v: "1", want: false},
		{s: nil, v: "1", want: false},
	}
)

func TestContains(t *testing.T) {
	for _, c := range containIntCases {
		get := Contain(c.s, c.v)
		assert.Equal(t, c.want, get)
	}
	for _, c := range containFloatCases {
		get := Contain(c.s, c.v)
		assert.Equal(t, c.want, get)
	}
	for _, c := range containStringCases {
		get := Contain(c.s, c.v)
		assert.Equal(t, c.want, get)
	}
}

var (
	countIntCases = []struct {
		s    []int
		want map[int]int
	}{
		{s: []int{1, 1, 1, 2, 2, 3, 4}, want: map[int]int{1: 3, 2: 2, 3: 1, 4: 1}},
		{s: []int{1, 2, 3}, want: map[int]int{1: 1, 2: 1, 3: 1}},
		{s: []int{}, want: map[int]int{}},
		{s: nil, want: map[int]int{}},
	}

	countFloatCases = []struct {
		s    []float64
		want map[float64]int
	}{
		{s: []float64{1.1, 1.1, 1.1, 2.2, 2.2, 3.3, 4.4}, want: map[float64]int{1.1: 3, 2.2: 2, 3.3: 1, 4.4: 1}},
		{s: []float64{1.1, 2.2, 3.3}, want: map[float64]int{1.1: 1, 2.2: 1, 3.3: 1}},
		{s: []float64{}, want: map[float64]int{}},
		{s: nil, want: map[float64]int{}},
	}

	countStringCases = []struct {
		s    []string
		want map[string]int
	}{
		{s: []string{"1", "1", "1", "2", "2", "3", "4"}, want: map[string]int{"1": 3, "2": 2, "3": 1, "4": 1}},
		{s: []string{"1", "2", "3"}, want: map[string]int{"1": 1, "2": 1, "3": 1}},
		{s: []string{}, want: map[string]int{}},
		{s: nil, want: map[string]int{}},
	}
)

func TestCount(t *testing.T) {
	for _, c := range countIntCases {
		get := Count(c.s)
		assert.Equal(t, c.want, get)
	}
	for _, c := range countFloatCases {
		get := Count(c.s)
		assert.Equal(t, c.want, get)
	}
	for _, c := range countStringCases {
		get := Count(c.s)
		assert.Equal(t, c.want, get)
	}
}

var (
	deleteIntCases = []struct {
		s      []int
		v      int
		n      int
		wantDs []int
		wantDn int
	}{
		{s: []int{1, 1, 1, 2, 2, 3, 4}, v: 1, n: 2, wantDs: []int{1, 2, 2, 3, 4}, wantDn: 2},
		{s: []int{1, 1, 1, 2, 2, 3, 4}, v: 1, n: 5, wantDs: []int{2, 2, 3, 4}, wantDn: 3},
		{s: []int{1, 1, 1, 2, 2, 3, 4}, v: 1, n: 0, wantDs: []int{1, 1, 1, 2, 2, 3, 4}, wantDn: 0},
		{s: []int{1, 1, 1, 2, 2, 3, 4}, v: 5, n: 5, wantDs: []int{1, 1, 1, 2, 2, 3, 4}, wantDn: 0},
		{s: []int{1, 1, 1, 2, 2, 3, 4}, v: 1, n: -1, wantDs: []int{2, 2, 3, 4}, wantDn: 3},
		{s: []int{}, v: 1, n: 2, wantDs: nil, wantDn: 0},
		{s: nil, v: 1, n: 2, wantDs: nil, wantDn: 0},
	}

	deleteFloatCases = []struct {
		s      []float64
		v      float64
		n      int
		wantDs []float64
		wantDn int
	}{
		{s: []float64{1.1, 1.1, 1.1, 2.2, 2.2, 3.3, 4.4}, v: 1.1, n: 2, wantDs: []float64{1.1, 2.2, 2.2, 3.3, 4.4}, wantDn: 2},
		{s: []float64{1.1, 1.1, 1.1, 2.2, 2.2, 3.3, 4.4}, v: 1.1, n: 5, wantDs: []float64{2.2, 2.2, 3.3, 4.4}, wantDn: 3},
		{s: []float64{1.1, 1.1, 1.1, 2.2, 2.2, 3.3, 4.4}, v: 1.1, n: 0, wantDs: []float64{1.1, 1.1, 1.1, 2.2, 2.2, 3.3, 4.4}, wantDn: 0},
		{s: []float64{1.1, 1.1, 1.1, 2.2, 2.2, 3.3, 4.4}, v: 5.5, n: 5, wantDs: []float64{1.1, 1.1, 1.1, 2.2, 2.2, 3.3, 4.4}, wantDn: 0},
		{s: []float64{1.1, 1.1, 1.1, 2.2, 2.2, 3.3, 4.4}, v: 1.1, n: -1, wantDs: []float64{2.2, 2.2, 3.3, 4.4}, wantDn: 3},
		{s: []float64{}, v: 1.1, n: 2, wantDs: nil, wantDn: 0},
		{s: nil, v: 1.1, n: 2, wantDs: nil, wantDn: 0},
	}

	deleteStringCases = []struct {
		s      []string
		v      string
		n      int
		wantDs []string
		wantDn int
	}{
		{s: []string{"1", "1", "1", "2", "2", "3", "4"}, v: "1", n: 2, wantDs: []string{"1", "2", "2", "3", "4"}, wantDn: 2},
		{s: []string{"1", "1", "1", "2", "2", "3", "4"}, v: "1", n: 5, wantDs: []string{"2", "2", "3", "4"}, wantDn: 3},
		{s: []string{"1", "1", "1", "2", "2", "3", "4"}, v: "1", n: 0, wantDs: []string{"1", "1", "1", "2", "2", "3", "4"}, wantDn: 0},
		{s: []string{"1", "1", "1", "2", "2", "3", "4"}, v: "5", n: 5, wantDs: []string{"1", "1", "1", "2", "2", "3", "4"}, wantDn: 0},
		{s: []string{"1", "1", "1", "2", "2", "3", "4"}, v: "1", n: -1, wantDs: []string{"2", "2", "3", "4"}, wantDn: 3},
		{s: []string{}, v: "1", n: 2, wantDs: nil, wantDn: 0},
		{s: nil, v: "1", n: 2, wantDs: nil, wantDn: 0},
	}
)

func TestDelete(t *testing.T) {
	for _, c := range deleteIntCases {
		getDs, getDn := Delete(c.s, c.v, c.n)
		assert.Equal(t, c.wantDs, getDs)
		assert.Equal(t, c.wantDn, getDn)
	}
	for _, c := range deleteFloatCases {
		getDs, getDn := Delete(c.s, c.v, c.n)
		assert.Equal(t, c.wantDs, getDs)
		assert.Equal(t, c.wantDn, getDn)
	}
	for _, c := range deleteStringCases {
		getDs, getDn := Delete(c.s, c.v, c.n)
		assert.Equal(t, c.wantDs, getDs)
		assert.Equal(t, c.wantDn, getDn)
	}
}

func TestDeleteFunc(t *testing.T) {
	for _, c := range deleteIntCases {
		getDs, getDn := DeleteFunc(c.s, func(v int) bool { return c.v == v }, c.n)
		assert.Equal(t, c.wantDs, getDs)
		assert.Equal(t, c.wantDn, getDn)
	}
	for _, c := range deleteFloatCases {
		getDs, getDn := DeleteFunc(c.s, func(v float64) bool { return c.v == v }, c.n)
		assert.Equal(t, c.wantDs, getDs)
		assert.Equal(t, c.wantDn, getDn)
	}
	for _, c := range deleteStringCases {
		getDs, getDn := DeleteFunc(c.s, func(v string) bool { return c.v == v }, c.n)
		assert.Equal(t, c.wantDs, getDs)
		assert.Equal(t, c.wantDn, getDn)
	}
}

var (
	equalIntCases = []struct {
		s1   []int
		s2   []int
		want bool
	}{
		{s1: []int{1, 2, 3, 4}, s2: []int{1, 2, 3, 4}, want: true},
		{s1: []int{1, 2, 3, 4}, s2: []int{1, 2, 3}, want: false},
		{s1: []int{1, 2, 3, 4}, s2: []int{1, 2, 3, 4, 5}, want: false},
		{s1: []int{}, s2: []int{}, want: true},
		{s1: nil, s2: []int{}, want: false},
		{s1: nil, s2: nil, want: true},
	}

	equalFloatCases = []struct {
		s1   []float64
		s2   []float64
		want bool
	}{
		{s1: []float64{1.1, 2.2, 3.3, 4.4}, s2: []float64{1.1, 2.2, 3.3, 4.4}, want: true},
		{s1: []float64{1.1, 2.2, 3.3, 4.4}, s2: []float64{1.1, 2.2, 3.3}, want: false},
		{s1: []float64{1.1, 2.2, 3.3, 4.4}, s2: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, want: false},
		{s1: []float64{}, s2: []float64{}, want: true},
		{s1: nil, s2: []float64{}, want: false},
		{s1: nil, s2: nil, want: true},
	}

	equalStringCases = []struct {
		s1   []string
		s2   []string
		want bool
	}{
		{s1: []string{"1", "2", "3", "4"}, s2: []string{"1", "2", "3", "4"}, want: true},
		{s1: []string{"1", "2", "3", "4"}, s2: []string{"1", "2", "3"}, want: false},
		{s1: []string{"1", "2", "3", "4"}, s2: []string{"1", "2", "3", "4", "5"}, want: false},
		{s1: []string{"1", "2", "3", "4"}, s2: []string{}, want: false},
		{s1: nil, s2: []string{}, want: false},
		{s1: nil, s2: nil, want: true},
	}
)

func TestEqual(t *testing.T) {
	for _, c := range equalIntCases {
		get := Equal(c.s1, c.s2)
		assert.Equal(t, c.want, get)
	}
	for _, c := range equalFloatCases {
		get := Equal(c.s1, c.s2)
		assert.Equal(t, c.want, get)
	}
	for _, c := range equalStringCases {
		get := Equal(c.s1, c.s2)
		assert.Equal(t, c.want, get)
	}
}

func TestEqualFunc(t *testing.T) {
	for _, c := range equalIntCases {
		get := EqualFunc(c.s1, c.s2, func(v1, v2 int) bool { return v1 == v2 })
		assert.Equal(t, c.want, get)
	}
	for _, c := range equalFloatCases {
		get := EqualFunc(c.s1, c.s2, func(v1, v2 float64) bool { return v1 == v2 })
		assert.Equal(t, c.want, get)
	}
	for _, c := range equalStringCases {
		get := EqualFunc(c.s1, c.s2, func(v1, v2 string) bool { return v1 == v2 })
		assert.Equal(t, c.want, get)
	}
}

var (
	fillIntCases = []struct {
		v    int
		n    int
		want []int
	}{
		{v: 1, n: 5, want: []int{1, 1, 1, 1, 1}},
		{v: 1, n: 3, want: []int{1, 1, 1}},
		{v: 1, n: 0, want: []int{}},
	}

	fillFloatCases = []struct {
		v    float64
		n    int
		want []float64
	}{
		{v: 1.1, n: 5, want: []float64{1.1, 1.1, 1.1, 1.1, 1.1}},
		{v: 1.1, n: 3, want: []float64{1.1, 1.1, 1.1}},
		{v: 1.1, n: 0, want: []float64{}},
	}

	fillStringCases = []struct {
		v    string
		n    int
		want []string
	}{
		{v: "1", n: 5, want: []string{"1", "1", "1", "1", "1"}},
		{v: "1", n: 3, want: []string{"1", "1", "1"}},
		{v: "1", n: 0, want: []string{}},
	}
)

var (
	extractIntCases = []struct {
		s []int
		n int
	}{
		{s: []int{1, 2, 3, 4, 5}, n: 2},
		{s: []int{1, 3, 5}, n: 5},
		{s: []int{}},
		{s: nil},
	}

	extractFloatCases = []struct {
		s []float64
		n int
	}{
		{s: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, n: 2},
		{s: []float64{1.1, 3.3, 5.5}, n: 5},
		{s: []float64{}},
		{s: nil},
	}

	extractStringCases = []struct {
		s []string
		n int
	}{
		{s: []string{"1", "2", "3", "4", "5"}, n: 2},
		{s: []string{"1", "3", "5"}, n: 5},
		{s: []string{}},
		{s: nil},
	}
)

func TestExtract(t *testing.T) {
	for _, c := range extractIntCases {
		get := Extract(c.s, c.n)
		if len(c.s) == 0 {
			assert.Zero(t, get)
		} else {
			t.Log(c.s, get)
			assert.True(t, c.n >= len(get))
			assert.Subset(t, c.s, get)
		}
	}
	for _, c := range extractFloatCases {
		get := Extract(c.s, c.n)
		if len(c.s) == 0 {
			assert.Zero(t, get)
		} else {
			t.Log(c.s, get)
			assert.True(t, c.n >= len(get))
			assert.Subset(t, c.s, get)
		}
	}
	for _, c := range extractStringCases {
		get := Extract(c.s, c.n)
		if len(c.s) == 0 {
			assert.Zero(t, get)
		} else {
			t.Log(c.s, get)
			assert.True(t, c.n >= len(get))
			assert.Subset(t, c.s, get)
		}
	}
}

func TestFill(t *testing.T) {
	for _, c := range fillIntCases {
		get := Fill(c.v, c.n)
		assert.Equal(t, c.want, get)
	}
	for _, c := range fillFloatCases {
		get := Fill(c.v, c.n)
		assert.Equal(t, c.want, get)
	}
	for _, c := range fillStringCases {
		get := Fill(c.v, c.n)
		assert.Equal(t, c.want, get)
	}
}

var (
	reverseIntCases = []struct {
		s    []int
		want []int
	}{
		{s: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{s: []int{1, 2, 3, 4}, want: []int{4, 3, 2, 1}},
		{s: []int{1, 2}, want: []int{2, 1}},
		{s: []int{1}, want: []int{1}},
		{s: []int{}, want: []int{}},
		{s: nil, want: []int{}},
	}

	reverseFloatCases = []struct {
		s    []float64
		want []float64
	}{
		{s: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, want: []float64{5.5, 4.4, 3.3, 2.2, 1.1}},
		{s: []float64{1.1, 2.2, 3.3, 4.4}, want: []float64{4.4, 3.3, 2.2, 1.1}},
		{s: []float64{1.1, 2.2}, want: []float64{2.2, 1.1}},
		{s: []float64{1.1}, want: []float64{1.1}},
		{s: []float64{}, want: []float64{}},
		{s: nil, want: []float64{}},
	}

	reverseStringCases = []struct {
		s    []string
		want []string
	}{
		{s: []string{"1", "2", "3", "4", "5"}, want: []string{"5", "4", "3", "2", "1"}},
		{s: []string{"1", "2", "3", "4"}, want: []string{"4", "3", "2", "1"}},
		{s: []string{"1", "2"}, want: []string{"2", "1"}},
		{s: []string{"1"}, want: []string{"1"}},
		{s: []string{}, want: []string{}},
		{s: nil, want: []string{}},
	}
)

func TestReverse(t *testing.T) {
	for _, c := range reverseIntCases {
		get := Reverse(c.s)
		assert.Equal(t, c.want, get)
	}
	for _, c := range reverseFloatCases {
		get := Reverse(c.s)
		assert.Equal(t, c.want, get)
	}
	for _, c := range reverseStringCases {
		get := Reverse(c.s)
		assert.Equal(t, c.want, get)
	}
}

var (
	shuffleIntCases = []struct {
		s []int
	}{
		{s: []int{1, 2, 3, 4, 5}},
		{s: []int{1}},
		{s: []int{}},
		{s: nil},
	}

	shuffleFloatCases = []struct {
		s []float64
	}{
		{s: []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{s: []float64{1.1}},
		{s: []float64{}},
		{s: nil},
	}

	shuffleStringCases = []struct {
		s []string
	}{
		{s: []string{"1", "2", "3", "4", "5"}},
		{s: []string{"1"}},
		{s: []string{}},
		{s: nil},
	}
)

func TestShuffle(t *testing.T) {
	for _, c := range shuffleIntCases {
		get := Shuffle(c.s)
		if len(c.s) == 0 {
			assert.Nil(t, get)
		} else {
			t.Log(c.s, get)
			assert.Subset(t, c.s, get)
		}
	}
	for _, c := range shuffleFloatCases {
		get := Shuffle(c.s)
		if len(c.s) == 0 {
			assert.Nil(t, get)
		} else {
			t.Log(c.s, get)
			assert.Subset(t, c.s, get)
		}
	}
	for _, c := range shuffleStringCases {
		get := Shuffle(c.s)
		if len(c.s) == 0 {
			assert.Nil(t, get)
		} else {
			t.Log(c.s, get)
			assert.Subset(t, c.s, get)
		}
	}
}

var (
	takeIntCases = []struct {
		s []int
	}{
		{s: []int{1, 2, 3, 4, 5}},
		{s: []int{1}},
		{s: []int{}},
		{s: nil},
	}

	takeFloatCases = []struct {
		s []float64
	}{
		{s: []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{s: []float64{1.1}},
		{s: []float64{}},
		{s: nil},
	}

	takeStringCases = []struct {
		s []string
	}{
		{s: []string{"1", "2", "3", "4", "5"}},
		{s: []string{"1"}},
		{s: []string{}},
		{s: nil},
	}
)

func TestTake(t *testing.T) {
	for _, c := range takeIntCases {
		get := Take(c.s)
		if len(c.s) == 0 {
			assert.Zero(t, get)
		} else {
			t.Log(c.s, get)
			assert.Contains(t, c.s, get)
		}
	}
	for _, c := range takeFloatCases {
		get := Take(c.s)
		if len(c.s) == 0 {
			assert.Zero(t, get)
		} else {
			t.Log(c.s, get)
			assert.Contains(t, c.s, get)
		}
	}
	for _, c := range takeStringCases {
		get := Take(c.s)
		if len(c.s) == 0 {
			assert.Zero(t, get)
		} else {
			t.Log(c.s, get)
			assert.Contains(t, c.s, get)
		}
	}
}

var (
	uniqueIntCases = []struct {
		s    []int
		want []int
	}{
		{s: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{s: []int{1, 1, 2, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{s: []int{1, 2, 1, 1, 2, 2}, want: []int{1, 2}},
		{s: []int{1}, want: []int{1}},
		{s: []int{}, want: []int{}},
		{s: nil, want: []int{}},
	}

	uniqueFloatCases = []struct {
		s    []float64
		want []float64
	}{
		{s: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, want: []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{s: []float64{1.1, 1.1, 2.2, 2.2, 3.3, 4.4, 5.5}, want: []float64{1.1, 2.2, 3.3, 4.4, 5.5}},
		{s: []float64{1.1, 2.2, 1.1, 1.1, 2.2, 2.2}, want: []float64{1.1, 2.2}},
		{s: []float64{1.1}, want: []float64{1.1}},
		{s: []float64{}, want: []float64{}},
		{s: nil, want: []float64{}},
	}

	uniqueStringCases = []struct {
		s    []string
		want []string
	}{
		{s: []string{"1", "2", "3", "4", "5"}, want: []string{"1", "2", "3", "4", "5"}},
		{s: []string{"1", "1", "2", "2", "3", "4", "5"}, want: []string{"1", "2", "3", "4", "5"}},
		{s: []string{"1", "2", "1", "1", "2", "2"}, want: []string{"1", "2"}},
		{s: []string{"1"}, want: []string{"1"}},
		{s: []string{}, want: []string{}},
		{s: nil, want: []string{}},
	}
)

func TestUnique(t *testing.T) {
	for _, c := range uniqueIntCases {
		get := Unique(c.s)
		assert.Equal(t, c.want, get)
	}
	for _, c := range uniqueFloatCases {
		get := Unique(c.s)
		assert.Equal(t, c.want, get)
	}
	for _, c := range uniqueStringCases {
		get := Unique(c.s)
		assert.Equal(t, c.want, get)
	}
}
