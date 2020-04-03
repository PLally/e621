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
			t.Errorf("Parsed '%v' does not equal out '%v'", parsed.Normalized(), data.out)
		}
	}
}

func TestParsedTags_Matches(t *testing.T) {

	testdata := []struct{
		parsed string
		container TagContainer
		shouldMatch bool
	}{
		{"singletag", TagContainer{
			General:   []string{"test", "test2"},
			Artist:    []string{"singletag"},

		}, true},
		{"~foo ~bar", TagContainer{
			General:   []string{"test", "bar"},
			Artist:    []string{"singletag"},

		}, true},
		{"-help", TagContainer{
			General:   []string{"help", "test2"},
		}, false},
		{"-help", TagContainer{
			General:   []string{"help2", "test2"},
		}, true},
	}


	for _, data := range testdata {
		parsed, err :=ParseTags(data.parsed, false)
		if err != nil {
			t.Error(err)
		}

		matched := parsed.Matches(data.container)
		if matched != data.shouldMatch {
			t.Errorf("%v: expected %v got %v", data.parsed, data.shouldMatch, matched)
		}
	}

}