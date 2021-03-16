package e926

import (
	"fmt"
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

func TestSession_FindTag(t *testing.T) {
	session := NewSession("e926.net", "e6_tests")
	tags, err := session.FindTag("letodoesart")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tags)
}
