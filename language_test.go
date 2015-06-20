package phonetictranscriptor

import "testing"

func TestLanguageLoad(t *testing.T) {
	lang, err := LoadLanguage("pl")
	if err != nil {
		t.Fail()
		return
	}
	// Test sorting
	leng := lang.Rules.Len()
	if len(lang.Rules[0].From) < len(lang.Rules[leng-1].From) {
		t.Fail()
	}
}
