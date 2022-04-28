package uuid

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextV1(t *testing.T) {
	c := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			c <- NextV1()
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

	t.Log(r)

	uniqueMap := make(map[string]struct{}, len(r))
	for _, s := range r {
		uniqueMap[s] = struct{}{}
	}
	isUnique := len(uniqueMap) == len(r)
	assert.True(t, isUnique)
}

func TestNextV4(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			t.Log(NextV4())
		}()
	}

	wg.Wait()
}

func TestDecompose(t *testing.T) {
	expect := "936dbe97-ec4e-4ded-b459-ef676b566485"

	u1, err := Parse("936dbe97-ec4e-4ded-b459-ef676b566485")
	assert.NoError(t, err)
	assert.Equal(t, expect, u1.String())

	u2, err := Parse("936dbe97ec4e4dedb459ef676b566485")
	assert.NoError(t, err)
	assert.Equal(t, expect, u2.String())

	_, err = Parse("err uuid")
	assert.Error(t, err)
}

func BenchmarkNextV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextV1()
	}
}

func BenchmarkNextV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextV1()
	}
}
