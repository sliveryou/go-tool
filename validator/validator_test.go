package validator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sliveryou/go-tool/v2/pointer"
)

func TestVerify(t *testing.T) {
	user := _User{
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
		_PageInfo: _PageInfo{
			Page:     1,
			PageSize: 0,
		},
	}

	err := Verify(&user)
	require.Error(t, err)
	t.Log(err, ParseErr(err))

	err = Verify([]*_User{&user})
	require.Error(t, err)
	t.Log(err, ParseErr(err))

	err = VerifyVar("sliveryouqq.com", "email")
	require.Error(t, err)
	t.Log(err, ParseErr(err))

	err = VerifyVarWithValue("abcd", "abce", "eqcsfield")
	require.Error(t, err)
	t.Log(err, ParseErr(err))

	r := _CreateTrainsReq{Objects: []*_CreateTrainReq{
		{
			Name:          "123",
			Company:       "",
			ResponsesUser: "",
			Time:          0,
			Description:   "",
		},
	}}
	err = Verify(r)
	require.Error(t, err)
	t.Log(err, ParseErr(err))

	upe := &_UpdatePhoneEmailReq{
		Phone:       new(string),
		Email:       new(string),
		Description: new(string),
	}
	err = Verify(upe)
	require.NoError(t, err)
}

func TestNewValidator(t *testing.T) {
	v, err := NewValidator()
	require.NoError(t, err)

	err = v.V.Var("abc", "email")
	err = v.TranslateAll(err)
	require.Error(t, err)
	t.Log(err)

	upe := &_UpdatePhoneEmailReq{
		Phone:       pointer.Of("123"),
		Email:       pointer.Of("abc"),
		Description: pointer.Of("desc"),
	}
	err = v.V.Struct(upe)
	err = v.TranslateAll(err)
	require.Error(t, err)
	t.Log(err)
}

type _PageInfo struct {
	Page     int `validate:"required" label:"页数"`
	PageSize int `validate:"required" label:"每条页数"`
}

type _User struct {
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
	_PageInfo
}

type _CreateTrainReq struct {
	Name          string `json:"name" validate:"required" label:"培训名称"` // 培训名称
	Company       string `json:"company,optional"`                      // 培训单位
	ResponsesUser string `json:"responses_user,optional"`               // 负责人
	Time          int64  `json:"time" validate:"required" label:"培训日期"` // 培训日期
	Description   string `json:"description,optional"`                  // 培训备注
}

type _CreateTrainsReq struct {
	Objects []*_CreateTrainReq `json:"objects" validate:"gt=0,dive" label:"培训列表"` // 培训列表
}

type _UpdatePhoneEmailReq struct {
	Phone       *string `json:"phone,optional" validate:"omitempty,eq=|len=11,eq=|number" label:"联系电话"` // 联系电话
	Email       *string `json:"email,optional" validate:"omitempty,eq=|email" label:"邮箱"`               // 邮箱
	Description *string `json:"description,optional"`                                                   // 备注
}
