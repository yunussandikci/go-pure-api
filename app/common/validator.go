package common

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

var (
	Validate   *validator.Validate
	translator ut.Translator
)

func init() {
	en := en.New()
	tempTranslator, _ := ut.New(en, en).GetTranslator("en")
	translator = tempTranslator
	Validate = validator.New()
	en_translations.RegisterDefaultTranslations(Validate, translator)
}

func TranslateValidationErrors(err error) (errorMessages string) {
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errorMessages += fmt.Sprintf("%s. ", e.Translate(translator))
		}
	}
	return strings.TrimRight(errorMessages, " ")
}
