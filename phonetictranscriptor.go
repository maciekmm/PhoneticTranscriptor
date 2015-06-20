package phonetictranscriptor

import (
	"errors"
	"strings"
)

//TranscriptByLanguage transcripts phonetic word to language specific phones
func TranscriptByLanguage(lang Language, phonetic string) (string, error) {
	outcome := phonetic
	//TODO: Replace with faster method
	for _, val := range lang.Rules {
		outcome = strings.Replace(outcome, val.From, val.To, -1)
	}
	return outcome, nil
}

//TranscriptByCode transripts phonetic word to language specific phones
func TranscriptByCode(lang string, phonetic string) (string, error) {
	la, ok := languages[lang]
	if !ok {
		var err error
		la, err = LoadLanguage(lang)
		if err != nil {
			return phonetic, errors.New("Could not find language")
		}
	}
	return TranscriptByLanguage(la, phonetic)
}
