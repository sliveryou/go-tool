package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorpAccount_IsValid(t *testing.T) {
	cases := []struct {
		corpaccount string
		expect      bool
	}{
		{corpaccount: "33050161963500000428", expect: true},
		{corpaccount: "3305016196350000042X", expect: false},
		{corpaccount: "", expect: false},
	}

	for _, c := range cases {
		get := NewCorpAccount(c.corpaccount)
		assert.Equal(t, c.expect, get.IsValid())
	}
}
