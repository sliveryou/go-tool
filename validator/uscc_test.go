package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUSCC_IsValid(t *testing.T) {
	cases := []struct {
		uscc   string
		expect bool
	}{
		{uscc: "91350100MA32C0EA03", expect: true},
		{uscc: "91350211M0000XUF46", expect: false},
		{uscc: "91350203M0001FUE2P", expect: false},
		{uscc: "913301086706280599", expect: true},
		{uscc: "", expect: false},
	}

	for _, c := range cases {
		get := NewUSCC(c.uscc)
		assert.Equal(t, c.expect, get.IsValid())
	}
}
