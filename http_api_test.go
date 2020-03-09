package e621

import (
	"testing"
)

func TestSession_GetPosts(t *testing.T) {
	session := NewSession("e926.net", "e6_tests")
	resp, err := session.GetPosts("fox", 320)
	if err != nil {
		t.Error(err)
	}
	if len(resp.Posts) != 320 {
		t.Fail()
	}
}

