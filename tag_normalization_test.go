package e621

import (
	"testing"
)

func TestParserTags(t *testing.T) {
	_, err := ParseTags("wolf fox male -female", false)
	if err != nil {
		t.Error(err)
	}

	_, err = ParseTags("wolf ", false)
	if err != nil {
		t.Error(err)
	}
}

func TestParsedTags_Matches(t *testing.T) {
	parsed, err := ParseTags("singletag", false)
	if err != nil {
		t.Error(err)
	}

	if !parsed.Matches(TagContainer{
		General:   []string{"test", "test2"},
		Artist:    []string{"singletag"},
	}) {
		t.Fail()
	}
}