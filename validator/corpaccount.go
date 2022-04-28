package validator

import (
	"strings"
)

// CorpAccount corp account validator.
type CorpAccount string

// NewCorpAccount new a corp account validator.
func NewCorpAccount(corpaccount string) CorpAccount {
	return CorpAccount(corpaccount)
}

// IsValid checks the corp account is valid.
func (ca CorpAccount) IsValid() bool {
	caStr := strings.ToUpper(string(ca))
	return corpaccountRegex.Match([]byte(caStr))
}
