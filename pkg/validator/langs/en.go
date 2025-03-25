package langs

import (
	"key-haven-back/pkg/validator"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func SetLanguageEN() ut.Translator {
	lang := en.New()
	uni := ut.New(lang, lang)
	validate := validator.GetValidator()
	trans, found := uni.GetTranslator("en")
	if !found {
		panic("translator not found")
	}
	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
	return trans
}
