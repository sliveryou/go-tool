package condition

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	t.Parallel()

	// bool
	assert.False(t, Bool(false))
	assert.True(t, Bool(true))

	// integer
	assert.False(t, Bool(0))
	assert.True(t, Bool(1))

	// float
	assert.False(t, Bool(0.0))
	assert.True(t, Bool(0.1))

	// string
	assert.False(t, Bool(""))
	assert.True(t, Bool(" "))
	assert.True(t, Bool("0"))

	// slice
	var nums [2]int
	assert.False(t, Bool(nums))
	nums = [2]int{0, 1}
	assert.True(t, Bool(nums))

	// map
	assert.False(t, Bool(map[string]string{}))
	assert.True(t, Bool(map[string]string{"a": "a"}))

	// channel
	var ch chan int
	assert.False(t, Bool(ch))
	ch = make(chan int)
	assert.True(t, Bool(ch))

	//  interface
	var err error
	assert.False(t, Bool(err))
	err = errors.New("error message")
	assert.True(t, Bool(err))

	// struct
	assert.False(t, Bool(struct{}{}))
	assert.True(t, Bool(time.Now()))

	// struct pointer
	type TestStruct struct{}
	ts := TestStruct{}
	assert.False(t, Bool(ts))
	assert.True(t, Bool(&ts))
}

func TestTernaryOperator(t *testing.T) {
	t.Parallel()

	trueValue := "1"
	falseValue := "0"

	assert.Equal(t, trueValue, TernaryOperator(true, trueValue, falseValue))
	assert.Equal(t, falseValue, TernaryOperator(false, trueValue, falseValue))
}
