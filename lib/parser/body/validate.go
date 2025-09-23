package bodyparser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/car-journal/api-backend/lib/httperror"
	"github.com/car-journal/api-backend/lib/httperror/const/errortype"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni *ut.UniversalTranslator
	v   *validator.Validate
)

func Validate(i interface{}, binding string) error {
	localeEN := en.New()
	uni = ut.New(localeEN, localeEN)
	trans, _ := uni.GetTranslator("en")

	v = validator.New()
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	err := v.Struct(i)
	if nil == err {
		return nil
	}
	if err, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}
	reflected := reflect.ValueOf(i)
	errs := err.(validator.ValidationErrors)
	var list []string
	for _, val := range errs {
		list = append(list, ValidationErrorToText(val, reflected, binding))
	}

	return httperror.New(errortype.InvalidInput, list)
}

func ValidationErrorToText(e validator.FieldError, reflectValue reflect.Value, b string) string {
	field, _ := reflectValue.Type().Elem().FieldByName(e.StructField())
	paramField, _ := reflectValue.Type().Elem().FieldByName(e.Param())
	var name string
	if name = field.Tag.Get(b); name == "" {
		name = strings.ToLower(e.StructField())
	}
	var param string
	if param = paramField.Tag.Get(b); param == "" {
		param = strings.ToLower(e.Param())
	}
	switch e.Tag() {
	case Required:
		return fmt.Sprintf("%s is required", name)
	case RequiredWithout:
		return fmt.Sprintf("the %s field is required when %s is not present", name, param)
	case RequiredWith:
		return fmt.Sprintf("the %s field is required when %s is present", name, param)
	case RequiredIf:
		return fmt.Sprintf("the %s field is required when %s", name, param)
	case Max:
		return fmt.Sprintf("%s cannot be longer than %s", name, param)
	case Min:
		return fmt.Sprintf("%s must be longer than %s", name, param)
	case EqField:
		return fmt.Sprintf("%s must be the same as %s", name, param)
	case Email:
		return "invalid email format"
	case Len:
		return fmt.Sprintf("%s must be %s characters long", name, param)
	}
	return fmt.Sprintf("%s is not valid", name)
}
