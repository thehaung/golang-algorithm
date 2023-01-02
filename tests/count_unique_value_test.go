package tests

import (
	"github.com/thehaung/golang-algorithm/solution"
	"testing"
)

func TestCountUniqueValue(t *testing.T) {
	expectResult := 2
	slice := []int{1, 2, 1, 1}

	if solution.CountUniqueValue(slice) != expectResult {
		t.Fail()
	}
}
