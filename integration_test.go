package xstr

import (
	"strings"
	"testing"
)

func TestHasPrefix(t *testing.T) {
	prefix := "www"
	str1 := "wwwg-tool-xstr"
	str2 := "g-tool-xstr"
	if !HasPrefix(str1, prefix) {
		t.Fatalf("%s should has %s as prefix", str1, prefix)
	}
	if HasPrefix(str2, prefix) {
		t.Fatalf("%s should not has %s as prefix", str2, prefix)
	}
	if HasPrefix(str1, prefix) != strings.HasPrefix(str1, prefix) || HasPrefix(str2, prefix) != strings.HasPrefix(str2, prefix) {
		t.Fatal("xstr.HasPrefix is not equal to strings.HasPrefix")
	}

}

// Testing just one function is enough
