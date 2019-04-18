package translation

import "github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"

func findModels(languageTranslator *languagetranslatorv3.LanguageTranslatorV3, source *string, target *string) []languagetranslatorv3.TranslationModel {
	response, err := languageTranslator.ListModels(
		&languagetranslatorv3.ListModelsOptions{
			Source: source,
			Target: target,
		},
	)
	if err != nil {
		panic(err)
	}
	return languageTranslator.GetListModelsResult(response).Models
}
