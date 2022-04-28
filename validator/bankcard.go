package validator

import (
	"strconv"
	"strings"
)

// payment card number (PAN), ISO/IEC 7812
// more details: https://zh.wikipedia.org/wiki/%E5%8F%91%E5%8D%A1%E8%A1%8C%E8%AF%86%E5%88%AB%E7%A0%81
// verification rules: https://zh.wikipedia.org/wiki/Luhn%E7%AE%97%E6%B3%95
// reference standard: https://zh.wikipedia.org/wiki/ISO/IEC_7812

// key represents the value of the bank card number, and value represents the value after the value * 2 is added.
var valueDoubleSumMap = map[int]int{
	0: 0, // 0 * 2 = 0, 0 = 0
	1: 2, // 1 * 2 = 2, 2 = 2
	2: 4, // 2 * 2 = 4, 4 = 4
	3: 6, // 3 * 2 = 6, 6 = 6
	4: 8, // 4 * 2 = 8, 8 = 8
	5: 1, // 5 * 2 = 10, 1 + 0 = 1
	6: 3, // 6 * 2 = 12, 1 + 2 = 3
	7: 5, // 7 * 2 = 14, 1 + 4 = 5
	8: 7, // 8 * 2 = 16, 1 + 6 = 7
	9: 9, // 9 * 2 = 18, 1+ 8 = 9
}

// BankCard bank card validator.
type BankCard string

// NewBankCard new a bank card validator.
func NewBankCard(bankcard string) BankCard {
	return BankCard(bankcard)
}

// IsValid checks the bank card is valid.
func (bc BankCard) IsValid() bool {
	bcStr := strings.ToUpper(string(bc))
	if !bankcardRegex.Match([]byte(bcStr)) {
		return false
	}

	sum := 0
	length := len(bcStr)

	for index := length - 1; index >= 0; index-- {
		if value, err := strconv.Atoi(string(bcStr[index])); err == nil {
			if (length-index)%2 == 0 {
				sum += valueDoubleSumMap[value]
			} else {
				sum += value
			}
		} else {
			return false
		}
	}

	return sum%10 == 0
}
