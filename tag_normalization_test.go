package e621

import (
	"testing"
)

func TestParsedTags_Normalized(t *testing.T) {
	testData := []struct{
		in string
		out string
	}{
		{"test", "test"},
		{"a c b", "a b c"},
		{"tag fox male -tag2", "-tag2 fox male tag"},
		{"wolf ", "wolf"},
	}

	for _, data := range testData {
		parsed, err := ParseTags(data.in, false)
		if err  != nil {
			t.Error(err)
		}
		if parsed.Normalized() != data.out {
			t.Errorf("Pared '%v' does not equal out '%v'", parsed.Normalized(), data.out)
		}
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