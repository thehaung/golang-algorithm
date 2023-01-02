package tests

import (
	"testing"

	"github.com/thehaung/golang-algorithm/solution"
)

func TestIsSumZero(t *testing.T) {
	slice := []int{-3, -2, -1, 0, 1, 1, 1}

	if len(solution.IsSumZero(slice)) == 0 {
		t.Fail()
	}
}
