package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankCard_IsValid(t *testing.T) {
	cases := []struct {
		bankcard string
		expect   bool
	}{
		{bankcard: "6222081203009584273", expect: true},
		{bankcard: "6221081204209584174", expect: false},
		{bankcard: "6225760008219524", expect: true},
		{bankcard: "0202020202", expect: false},
		{bankcard: "6259650871772098", expect: true},
		{bankcard: "", expect: false},
	}

	for _, c := range cases {
		get := NewBankCard(c.bankcard)
		assert.Equal(t, c.expect, get.IsValid())
	}
}
