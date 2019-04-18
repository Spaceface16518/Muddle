package translation

import (
	"log"
	"math/rand"

	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
)

func translateTo(languageTranslator *languagetranslatorv3.LanguageTranslatorV3, text *[]string, source *string, target *string) []string {
	resp, err := languageTranslator.Translate(
		&languagetranslatorv3.TranslateOptions{
			Text:   *text,
			Source: source,
			Target: target,
		},
	)

	if err != nil {
		panic(err)
	}

	r := languageTranslator.GetTranslateResult(resp).Translations
	output := make([]string, len(r))

	for i := 0; i < len(r); i++ {
		output[i] = *r[i].TranslationOutput
	}

	return output
}

// Translate takes a pointer to a LanguageTranslatorV3 service and translates some text (specified by pointer) into a random language, given the source language of the text.
func Translate(languageTranslator *languagetranslatorv3.LanguageTranslatorV3, text *[]string, source *string) (string, []string) {

	models := findModels(languageTranslator, source, nil)

	lang := models[rand.Int()%len(models)]
	log.Printf("Translating from %s to %s\n", *source, *lang.Target)
	// TODO: remove this if it's redundant
	if *lang.Target == *source {
		log.Printf("Languages are the same. Skipping translation\n")
		return *lang.Target, *text
	}

	return *lang.Target, translateTo(languageTranslator, text, source, lang.Target)
}
