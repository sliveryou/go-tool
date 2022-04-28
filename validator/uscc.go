package validator

import (
	"strings"
)

// the coding rule of the unified social credit identifier for legal entities and other organizations, GB 32100-2015
// verification rules: http://openstd.samr.gov.cn/bzgk/gb/newGbInfo?hcno=24691C25985C1073D3A7C85629378AC0
// online preview: https://zh.wikisource.org/wiki/GB_32100-2015_%E6%B3%95%E4%BA%BA%E5%92%8C%E5%85%B6%E4%BB%96%E7%BB%84%E7%BB%87%E7%BB%9F%E4%B8%80%E7%A4%BE%E4%BC%9A%E4%BF%A1%E7%94%A8%E4%BB%A3%E7%A0%81%E7%BC%96%E7%A0%81%E8%A7%84%E5%88%99

// char - num map
// key represents the code character of uscc, and value represents the corresponding value of the code character
var charValueMap = map[uint8]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
	'G': 16,
	'H': 17,
	'J': 18,
	'K': 19,
	'L': 20,
	'M': 21,
	'N': 22,
	'P': 23,
	'Q': 24,
	'R': 25,
	'T': 26,
	'U': 27,
	'W': 28,
	'X': 29,
	'Y': 30,
}

// index - weight map
// key represents the index of uscc and value represents the weight corresponding to the index
// value = 3 ** （key） % 31
var usccIndexWeightMap = map[int]int{
	0:  1,
	1:  3,
	2:  9,
	3:  27,
	4:  19,
	5:  26,
	6:  16,
	7:  17,
	8:  20,
	9:  29,
	10: 25,
	11: 13,
	12: 8,
	13: 24,
	14: 10,
	15: 30,
	16: 28,
}

// USCC uscc validator.
type USCC string

// NewUSCC new an uscc validator.
func NewUSCC(uscc string) USCC {
	return USCC(uscc)
}

// IsValid checks the uscc is valid.
func (uscc USCC) IsValid() bool {
	usccStr := strings.ToUpper(string(uscc))
	if !usccRegex.Match([]byte(usccStr)) {
		return false
	}

	sum := 0
	checkCode := usccStr[17]
	for index := range usccStr[:17] {
		value := charValueMap[usccStr[index]]
		// 计算加权因子
		weight := usccIndexWeightMap[index]
		// 计算加权和
		sum += value * weight
	}

	c := 31 - sum%31
	if c == 31 {
		c = 0
	}
	cChar := uint8(0)
	for char, value := range charValueMap {
		if value == c {
			cChar = char
			break
		}
	}

	// fmt.Println(string(cChar), string(checkCode))
	return cChar == checkCode
}
