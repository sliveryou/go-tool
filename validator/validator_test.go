package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	type PageInfo struct {
		Page     int `validate:"required" label:"页数"`
		PageSize int `validate:"required" label:"每条页数"`
	}

	type User struct {
		IdCard      string `validate:"required,min=3,idcard" label:"身份证号"`
		BankCard    string `validate:"required,oneof=0 1 2,bankcard" label:"银行卡号"`
		USCC        string `validate:"required,uscc" label:"统一社会信用代码"`
		CorpAccount string `validate:"required,corpaccount" label:"对公账户"`
		Code        *int   `validate:"omitempty,oneof=1 2" label:"代码"`
		TT          bool
		Path        string `validate:"required_if=TT true" label:"接口路径"`
		Method      string `validate:"required,httpmethod" label:"接口方法"`
		Email       string `validate:"omitempty,email" label:"电子邮箱"`
		Phone       string `validate:"omitempty,len=11" label:"手机号"`
		PageInfo
	}

	user := User{
		IdCard:      "123",
		BankCard:    "4",
		USCC:        "",
		CorpAccount: "",
		Code:        new(int),
		TT:          true,
		Path:        "",
		Method:      "SELECT",
		Email:       "sliveryou@qq.com",
		Phone:       "1",
		PageInfo: PageInfo{
			Page:     1,
			PageSize: 0,
		},
	}

	err := Verify(&user)
	if assert.Error(t, err) {
		t.Log(err, ParseErr(err))
	}

	err = Verify([]*User{&user})
	if assert.Error(t, err) {
		t.Log(err, ParseErr(err))
	}

	err = VerifyVar("sliveryouqq.com", "email")
	if assert.Error(t, err) {
		t.Log(err, ParseErr(err))
	}

	err = VerifyVarWithValue("abcd", "abce", "eqcsfield")
	if assert.Error(t, err) {
		t.Log(err, ParseErr(err))
	}
}
