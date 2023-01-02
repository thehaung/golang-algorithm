package tests

import (
	"testing"

	"github.com/thehaung/golang-algorithm/solution"
)

func TestIsAnagram(t *testing.T) {
	str1 := "qwerty"
	str2 := "qeywrt"

	if !solution.IsAnagram(str1, str2) {
		t.Fatal("IsAnagram(str1, str2) not anagram")
	}
}
