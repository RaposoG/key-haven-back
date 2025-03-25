package langs

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
)

func SetTranslate(t string) ut.Translator {
	lowercase := strings.ToLower(t)
	switch lowercase {
	case "pt", "pt-br", "pt_br", "br":
		return SetLanguagePTBR()
	case "en", "en-us", "en_us", "us":
		return SetLanguageEN()
	default:
		return SetLanguageEN()
	}
}
