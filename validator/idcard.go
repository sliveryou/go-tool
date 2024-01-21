package validator

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/sliveryou/go-tool/v2/timex"
)

// citizen identification number, GB 11643-1999
// verification rules: http://openstd.samr.gov.cn/bzgk/gb/newGbInfo?hcno=080D6FBF2BB468F9007657F26D60013E
// online preview: https://zh.wikisource.org/wiki/GB_11643-1999_%E5%85%AC%E6%B0%91%E8%BA%AB%E4%BB%BD%E5%8F%B7%E7%A0%81

// index - weight map
// key represents the index of id card number and value represents the weight corresponding to the index
// value = 2 ** （key） % 11
var idCardIndexWeightMap = map[int]int{
	0:  7,
	1:  9,
	2:  10,
	3:  5,
	4:  8,
	5:  4,
	6:  2,
	7:  1,
	8:  6,
	9:  3,
	10: 7,
	11: 9,
	12: 10,
	13: 5,
	14: 8,
	15: 4,
	16: 2,
}

// sum - code map
// key represents the modulus of 11 after weighted summation, and value represents the check code
var modCheckCodeMap = map[int]uint8{
	0:  '1',
	1:  '0',
	2:  'X',
	3:  '9',
	4:  '8',
	5:  '7',
	6:  '6',
	7:  '5',
	8:  '4',
	9:  '3',
	10: '2',
}

// IdCard id card validator.
type IdCard string

// NewIdCard new an id card validator.
func NewIdCard(idcard string) IdCard {
	return IdCard(idcard)
}

// IsValid checks the id card is valid.
func (ic IdCard) IsValid() bool {
	icStr := strings.ToUpper(string(ic))
	if !idcardRegex.Match([]byte(icStr)) {
		return false
	}

	sum := 0
	checkCode := icStr[17]

	for index := range icStr[:17] {
		if a, err := strconv.Atoi(string(icStr[index])); err == nil {
			// 计算加权因子
			w := idCardIndexWeightMap[index]
			// 计算加权和
			sum += a * w
		} else {
			return false
		}
	}

	// fmt.Println(string(modCheckCodeMap[sum%11]), string(checkCode))
	return modCheckCodeMap[sum%11] == checkCode
}

// GetBirthday gets the birthday recorded on id card.
func (ic IdCard) GetBirthday() (time.Time, error) {
	if !ic.IsValid() {
		return time.Time{}, errors.New("validator: invalid idcard")
	}

	icStr := string(ic)
	yearStr, monthStr, dayStr := icStr[6:10], icStr[10:12], icStr[12:14]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return time.Time{}, err
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return time.Time{}, err
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, timex.Shanghai()), nil
}

const (
	// Female gender female.
	Female = 0
	// Male gender male.
	Male = 1
)

// GetGender gets the gender recorded on id card.
func (ic IdCard) GetGender() (int, error) {
	if !ic.IsValid() {
		return 0, errors.New("validator: invalid idcard")
	}

	icStr := string(ic)
	numStr := icStr[14:17]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}

	return num % 2, nil
}

// IsMale judges whether it is male.
func (ic IdCard) IsMale() (bool, error) {
	gender, err := ic.GetGender()
	if err != nil {
		return false, err
	}

	return gender == Male, nil
}

// IsFemale judges whether it is female.
func (ic IdCard) IsFemale() (bool, error) {
	gender, err := ic.GetGender()
	if err != nil {
		return false, err
	}

	return gender == Female, nil
}
