package tests

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/thehaung/golang-algorithm/solution"
)

func TestIsAnagram(t *testing.T) {
	str1 := "qwerty"
	str2 := "qeywrt"

	actual := solution.IsAnagram(str1, str2)
	require.Equal(t, true, actual)
}

func TestInvalidAnagram(t *testing.T) {
	str1 := "qwerty"
	str2 := "qeyw1rt"

	actual := solution.IsAnagram(str1, str2)
	require.Equal(t, false, actual)
}
