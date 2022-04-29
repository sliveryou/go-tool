package validator

import (
	"errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zht "github.com/go-playground/validator/v10/translations/zh"
)

const (
	bankcardRegexString    = "^[0-9]{15,19}$"
	corpaccountRegexString = "^[0-9]{9,25}$"
	idcardRegexString      = "^[0-9]{17}[0-9X]$"
	usccRegexString        = "^[A-Z0-9]{18}$"
)

var (
	// ErrUnexpected unexpected error.
	ErrUnexpected = errors.New("参数校验出错")

	ti ut.Translator
	vi *validator.Validate

	bankcardRegex    = regexp.MustCompile(bankcardRegexString)
	corpaccountRegex = regexp.MustCompile(corpaccountRegexString)
	idcardRegex      = regexp.MustCompile(idcardRegexString)
	usccRegex        = regexp.MustCompile(usccRegexString)

	httpMethodMap = map[string]struct{}{
		"GET":     {},
		"POST":    {},
		"PUT":     {},
		"DELETE":  {},
		"PATCH":   {},
		"OPTIONS": {},
	}
	validatorFuncMap = map[string]validator.Func{
		"idcard":      idcard,
		"bankcard":    bankcard,
		"uscc":        uscc,
		"corpaccount": corpaccount,
		"httpmethod":  httpmethod,
	}
	defaultTags = []string{
		"required_if",
		"required_unless",
		"required_with",
		"required_with_all",
		"required_without",
		"required_without_all",
		"excluded_with",
		"excluded_with_all",
		"excluded_without",
		"excluded_without_all",
		"isdefault",
		"fieldcontains",
		"fieldexcludes",
		"boolean",
		"e164",
		"urn_rfc2141",
		"file",
		"base64url",
		"startsnotwith",
		"endsnotwith",
		"eth_addr",
		"btc_addr",
		"btc_addr_bech32",
		"uuid_rfc4122",
		"uuid3_rfc4122",
		"uuid4_rfc4122",
		"uuid5_rfc4122",
		"hostname",
		"hostname_rfc1123",
		"fqdn",
		"unique",
		"html",
		"html_encoded",
		"url_encoded",
		"dir",
		"jwt",
		"hostname_port",
		"timezone",
		"iso3166_1_alpha2",
		"iso3166_1_alpha3",
		"iso3166_1_alpha_numeric",
		"iso3166_2",
		"iso4217",
		"iso4217_numeric",
		"bcp47_language_tag",
		"postcode_iso3166_alpha2",
		"postcode_iso3166_alpha2_field",
		"bic",
	}
)

type validateErrors []string

// Error returns the validation error string and implement the error interface,
// by default, only the validation error of the first field is returned.
func (ves validateErrors) Error() string {
	if len(ves) > 0 {
		return ves[0]
	}

	return ErrUnexpected.Error()
}

// ParseErr parses the content of validation error.
func ParseErr(err error) string {
	ves, ok := err.(validateErrors)
	if ok && len(ves) > 0 {
		return strings.Join(ves, ",")
	}

	return err.Error()
}

func init() {
	var err error
	vi = validator.New()
	vi.SetTagName("validate")
	vi.RegisterTagNameFunc(getLabelTagName)

	zhi := zh.New()
	uti := ut.New(zhi)
	ti, _ = uti.GetTranslator("zh")

	err = zht.RegisterDefaultTranslations(vi, ti)
	checkErr(err)

	for tag, fn := range validatorFuncMap {
		err = vi.RegisterValidation(tag, fn)
		checkErr(err)
		err = registerTranslation(tag, vi, ti, true)
		checkErr(err)
	}
	for _, defaultTag := range defaultTags {
		_ = registerTranslation(defaultTag, vi, ti, false)
	}
}

// Verify checks the data validity of the exportable field of the structure according to the validate tag.
func Verify(obj interface{}) error {
	return convertErr(vi.Struct(obj))
}

// VerifyVar checks the data validity of the field according to the validate tag.
func VerifyVar(field interface{}, tag string) error {
	return convertErr(vi.Var(field, tag))
}

// VerifyVarWithValue checks the data validity of the field against another field according to the validate tag.
func VerifyVarWithValue(field, other interface{}, tag string) error {
	return convertErr(vi.VarWithValue(field, other, tag))
}

// convertErr converts err to validation error.
func convertErr(err error) error {
	if err != nil {
		var ves validateErrors
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return ves
		}
		for _, err := range err.(validator.ValidationErrors) {
			ves = append(ves, err.Translate(ti))
		}

		return ves
	}

	return nil
}

// getLabelTagName gets the label name.
func getLabelTagName(sf reflect.StructField) string {
	name := strings.SplitN(sf.Tag.Get("label"), ",", 2)[0]
	if name == "-" {
		return ""
	} else if name == "" {
		return sf.Name
	}

	return name
}

// registerTranslation registers translator.
func registerTranslation(tag string, v *validator.Validate, t ut.Translator, override bool) error {
	return v.RegisterTranslation(tag, t, registerTranslationsFunc(tag, override), translationFunc(tag))
}

// registerTranslationsFunc register translation function.
func registerTranslationsFunc(key string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		return ut.Add(key, "{0}校验失败", override)
	}
}

// translationFunc translation function.
func translationFunc(key string) validator.TranslationFunc {
	return func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(key, fe.Field())
		return t
	}
}

// idcard id card validator.
func idcard(fl validator.FieldLevel) bool {
	return NewIdCard(fl.Field().String()).IsValid()
}

// bankcard bank card validator.
func bankcard(fl validator.FieldLevel) bool {
	return NewBankCard(fl.Field().String()).IsValid()
}

// uscc uscc validator.
func uscc(fl validator.FieldLevel) bool {
	return NewUSCC(fl.Field().String()).IsValid()
}

// corpaccount corp account validator.
func corpaccount(fl validator.FieldLevel) bool {
	return NewCorpAccount(fl.Field().String()).IsValid()
}

// httpmethod http method validator.
func httpmethod(fl validator.FieldLevel) bool {
	_, ok := httpMethodMap[strings.ToUpper(fl.Field().String())]
	return ok
}

// checkErr checks the err.
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
