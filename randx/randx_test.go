package randx

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdSource(t *testing.T) {
	const expect = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	assert.Equal(t, expect, StdSource())
}

func TestStdNumberSource(t *testing.T) {
	const expect = "0123456789"
	assert.Equal(t, expect, StdNumberSource())
}

func TestNewInCurrency(t *testing.T) {
	c := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			c <- NewString(10)
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	r := make([]string, 0, 100)
	for s := range c {
		r = append(r, s)
	}

	uniqueMap := make(map[string]struct{}, len(r))
	for _, s := range r {
		uniqueMap[s] = struct{}{}
	}
	isUnique := len(uniqueMap) == len(r)
	t.Logf("isUnique:%v", isUnique)
}

func TestNewString(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			t.Log(NewString(10))
		}()
	}

	wg.Wait()
}

func TestNewNumber(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			t.Log(NewNumber(10))
		}()
	}

	wg.Wait()
}

func TestNewWithSource(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			t.Log(NewWithSource(10, "ABCDEFGH"))
		}()
	}

	wg.Wait()

	assert.PanicsWithValue(t, "randx: wrong source length", func() {
		NewWithSource(10, NewString(257))
	})
	assert.PanicsWithValue(t, "randx: wrong source length", func() {
		NewWithSource(10, "")
	})
}

func BenchmarkNewString(b *testing.B) {
	benchFunc := func(b *testing.B, length int) {
		for i := 0; i < b.N; i++ {
			NewString(length)
		}
	}

	lengths := []int{6, 10, 20, 50, 100, 1000}
	for _, length := range lengths {
		b.Run(fmt.Sprintf("length=%d", length), func(b *testing.B) {
			benchFunc(b, length)
		})
	}
}

func BenchmarkNewNumber(b *testing.B) {
	benchFunc := func(b *testing.B, length int) {
		for i := 0; i < b.N; i++ {
			NewNumber(length)
		}
	}

	lengths := []int{6, 10, 20, 50, 100, 1000}
	for _, length := range lengths {
		b.Run(fmt.Sprintf("length=%d", length), func(b *testing.B) {
			benchFunc(b, length)
		})
	}
}

func BenchmarkNewWithSource(b *testing.B) {
	benchFunc := func(b *testing.B, length int, source string) {
		for i := 0; i < b.N; i++ {
			NewWithSource(length, source)
		}
	}

	lengths := []int{6, 10, 20, 50, 100, 1000}
	sources := []string{
		"0123456789",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
	for i := 0; i < len(lengths); i++ {
		for j := 0; j < len(sources); j++ {
			b.Run(fmt.Sprintf("length=%d sosuce=%s", lengths[i], sources[j]), func(b *testing.B) {
				benchFunc(b, lengths[i], sources[j])
			})
		}
	}
}
