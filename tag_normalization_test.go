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
	fmt.Println(parsed.Normalized())
}