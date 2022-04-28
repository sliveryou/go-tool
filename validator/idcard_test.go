package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdCard_IsValid(t *testing.T) {
	cases := []struct {
		idcard string
		expect bool
	}{
		{idcard: "330333199001053317", expect: true},
		{idcard: "330333200801052846", expect: false},
		{idcard: "", expect: false},
	}

	for _, c := range cases {
		get := NewIdCard(c.idcard)
		assert.Equal(t, c.expect, get.IsValid())
	}
}

func TestIdCard_GetBirthday(t *testing.T) {
	cases := []struct {
		idcard         string
		expectBirthday string
		expectErr      bool
	}{
		{idcard: "330333199001053317", expectBirthday: "1990-01-05", expectErr: false},
		{idcard: "330333200801052846", expectBirthday: "none", expectErr: true},
		{idcard: "", expectBirthday: "none", expectErr: true},
	}

	for _, c := range cases {
		get := NewIdCard(c.idcard)
		birthday, err := get.GetBirthday()
		if c.expectErr {
			assert.EqualError(t, err, "validator: invalid idcard")
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expectBirthday, birthday.Format("2006-01-02"))
		}
	}
}

func TestIdCard_GetGender(t *testing.T) {
	cases := []struct {
		idcard       string
		expectGender int
		expectErr    bool
	}{
		{idcard: "330333199001053317", expectGender: 1, expectErr: false},
		{idcard: "330333200801052846", expectGender: 0, expectErr: true},
		{idcard: "", expectGender: 0, expectErr: true},
	}

	for _, c := range cases {
		get := NewIdCard(c.idcard)
		gender, err := get.GetGender()
		if c.expectErr {
			assert.EqualError(t, err, "validator: invalid idcard")
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expectGender, gender)
		}
	}
}

func TestIdCard_IsMale(t *testing.T) {
	cases := []struct {
		idcard     string
		expectMale bool
		expectErr  bool
	}{
		{idcard: "330333199001053317", expectMale: true, expectErr: false},
		{idcard: "330333200801052846", expectMale: false, expectErr: true},
		{idcard: "", expectMale: false, expectErr: true},
	}

	for _, c := range cases {
		get := NewIdCard(c.idcard)
		isMale, err := get.IsMale()
		if c.expectErr {
			assert.EqualError(t, err, "validator: invalid idcard")
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expectMale, isMale)
		}
	}
}

func TestIdCard_IsFemale(t *testing.T) {
	cases := []struct {
		idcard       string
		expectFemale bool
		expectErr    bool
	}{
		{idcard: "330333199001053317", expectFemale: false, expectErr: false},
		{idcard: "330333200801052846", expectFemale: false, expectErr: true},
		{idcard: "", expectFemale: false, expectErr: true},
	}

	for _, c := range cases {
		get := NewIdCard(c.idcard)
		isFemale, err := get.IsFemale()
		if c.expectErr {
			assert.EqualError(t, err, "validator: invalid idcard")
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expectFemale, isFemale)
		}
	}
}
