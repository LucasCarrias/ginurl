package shortener

import (
	"regexp"
	"testing"
)

func TestBuildShortenedUrl(t *testing.T) {
	want := regexp.MustCompile(`[a-zA-Z0-9]{5}`)
	result := CreateCode()
	
	if !want.MatchString(result) {
		t.Fatalf(`CreateCode("") = %q, didn't macth the pattern %s`, result, `[a-zA-Z0-9]{5}`)
	}
}