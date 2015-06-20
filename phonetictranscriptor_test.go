package phonetictranscriptor

import (
	"log"
	"testing"
)

func TestTranspiling(t *testing.T) {
	lang, err := LoadLanguage("pl")
	if err != nil {
		t.Fail()
		return
	}

	out, err := TranscriptByLanguage(lang, "smæl")
	if out != "smal" {
		t.Fail()
	}
	out, _ = TranscriptByLanguage(lang, "fix")
	log.Println(out)
}
