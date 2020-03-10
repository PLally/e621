package e621

import (
	"fmt"
	"testing"
)

func TestParserTags(t *testing.T) {
	parsed, err := ParseTags("wolf fox male -female", false)
	if err != nil {
		t.Error(err)
	}
	// TODO
	fmt.Println(parsed.Normalized())
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