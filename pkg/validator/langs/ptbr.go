package langs

import (
	"key-haven-back/pkg/validator"

	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	pt_BR_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

func SetLanguagePTBR() ut.Translator {
	lang := pt_BR.New()
	uni := ut.New(lang, lang)
	validate := validator.GetValidator()
	trans, found := uni.GetTranslator("pt_BR")
	if !found {
		panic("translator not found")
	}
	err := pt_BR_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
	return trans
}
