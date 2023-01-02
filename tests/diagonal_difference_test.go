package tests

import (
	"github.com/thehaung/golang-algorithm/solution"
	"testing"
)

func TestDiagonalDifference(t *testing.T) {
	arr := [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	expectResult := 15

	if solution.DiagonalDifference(arr) != expectResult {
		t.Fail()
	}
}
