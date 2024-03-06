package validator

import (
	"errors"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	zht "github.com/go-playground/validator/v10/translations/zh"
)

// ErrUnexpected unexpected error.
var ErrUnexpected = errors.New("参数校验出错")

const (
	bankcardRegexString    = "^[0-9]{15,19}$"
	corpaccountRegexString = "^[0-9]{9,25}$"
	idcardRegexString      = "^[0-9]{17}[0-9X]$"
	usccRegexString        = "^[A-Z0-9]{18}$"
)

var (
	v = MustNewValidator()

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
		"skip_unless",
		"eq_ignore_case",
		"ne_ignore_case",
		"fieldcontains",
		"fieldexcludes",
		"boolean",
		"e164",
		"http_url",
		"urn_rfc2141",
		"file",
		"filepath",
		"base64url",
		"base64rawurl",
		"startsnotwith",
		"endsnotwith",
		"eth_addr",
		"eth_addr_checksum",
		"btc_addr",
		"btc_addr_bech32",
		"uuid_rfc4122",
		"uuid3_rfc4122",
		"uuid4_rfc4122",
		"uuid5_rfc4122",
		"md4",
		"md5",
		"sha256",
		"sha384",
		"sha512",
		"ripemd128",
		"ripemd160",
		"tiger128",
		"tiger160",
		"tiger192",
		"hostname",
		"hostname_rfc1123",
		"fqdn",
		"unique",
		"html",
		"html_encoded",
		"url_encoded",
		"dir",
		"dirpath",
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
		"semver",
		"dns_rfc1035_label",
		"credit_card",
		"cve",
		"luhn_checksum",
		"mongodb",
		"cron",
		"spicedb",
	}
)

// Validator represents the validator structure.
type Validator struct {
	V *validator.Validate
	T ut.Translator
}

// NewValidator new a validator.
func NewValidator() (*Validator, error) {
	vi := validator.New()
	vi.SetTagName("validate")
	vi.RegisterTagNameFunc(getLabelTagName)

	zhi := zh.New()
	uti := ut.New(zhi)
	ti, _ := uti.GetTranslator("zh")

	if err := zht.RegisterDefaultTranslations(vi, ti); err != nil {
		return nil, err
	}

	for tag, fn := range validatorFuncMap {
		if err := vi.RegisterValidation(tag, fn); err != nil {
			return nil, err
		}
		if err := registerTranslation(tag, vi, ti, true); err != nil {
			return nil, err
		}
	}

	for _, defaultTag := range defaultTags {
		if err := registerTranslation(defaultTag, vi, ti, false); err != nil {
			return nil, err
		}
	}

	return &Validator{
		V: vi,
		T: ti,
	}, nil
}

// MustNewValidator must new a validator.
func MustNewValidator() *Validator {
	v, err := NewValidator()
	if err != nil {
		panic(err)
	}

	return v
}

// Validate validates the request and parsed data.
func (v *Validator) Validate(r *http.Request, data any) error {
	return v.Translate(v.V.Struct(data))
}

// Translate translates the validation errors.
func (v *Validator) Translate(err error) error {
	if err != nil {
		var (
			es  validateErrors
			ves validator.ValidationErrors
		)

		if errors.As(err, &ves) {
			for _, ve := range ves {
				es = append(es, ve.Translate(v.T))
			}
		}

		return es
	}

	return nil
}

// TranslateAll translates all validation errors.
func (v *Validator) TranslateAll(err error) error {
	if err != nil {
		var (
			es  []string
			ves validator.ValidationErrors
		)

		if errors.As(err, &ves) {
			for _, ve := range ves {
				es = append(es, ve.Translate(v.T))
			}
		}

		if len(es) == 0 {
			return ErrUnexpected
		}

		return errors.New(strings.Join(es, ","))
	}

	return nil
}

// validateErrors represents the validation error strings.
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
	if err == nil {
		return ""
	}

	var ves validateErrors
	ok := errors.As(err, &ves)
	if ok && len(ves) > 0 {
		return strings.Join(ves, ",")
	}

	return err.Error()
}

// Verify checks the data validity of the exportable field of the structure according to the validate tag.
func Verify(obj interface{}) error {
	return v.Translate(v.V.Struct(obj))
}

// VerifyVar checks the data validity of the field according to the validate tag.
func VerifyVar(field interface{}, tag string) error {
	return v.Translate(v.V.Var(field, tag))
}

// VerifyVarWithValue checks the data validity of the field against another field according to the validate tag.
func VerifyVarWithValue(field, other interface{}, tag string) error {
	return v.Translate(v.V.VarWithValue(field, other, tag))
}

// getLabelTagName gets the label name.
func getLabelTagName(sf reflect.StructField) string {
	name := sf.Name
	if labelTag := strings.SplitN(sf.Tag.Get("label"), ",", 2)[0]; labelTag != "" {
		if name = labelTag; name == "-" {
			name = ""
		}
	} else if jsonTag := strings.SplitN(sf.Tag.Get("json"), ",", 2)[0]; jsonTag != "" {
		if name = jsonTag; name == "-" {
			name = ""
		}
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

// idcard represents the id card validator.
func idcard(fl validator.FieldLevel) bool {
	return NewIdCard(fl.Field().String()).IsValid()
}

// bankcard represents the bank card validator.
func bankcard(fl validator.FieldLevel) bool {
	return NewBankCard(fl.Field().String()).IsValid()
}

// uscc represents the uscc validator.
func uscc(fl validator.FieldLevel) bool {
	return NewUSCC(fl.Field().String()).IsValid()
}

// corpaccount represents the corp account validator.
func corpaccount(fl validator.FieldLevel) bool {
	return NewCorpAccount(fl.Field().String()).IsValid()
}

// httpmethod represents the http method validator.
func httpmethod(fl validator.FieldLevel) bool {
	_, ok := httpMethodMap[strings.ToUpper(fl.Field().String())]
	return ok
}
