package tests

import (
	"testing"

	"github.com/thehaung/golang-algorithm/solution"
)

func TestSame(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 4, 9}

	if !solution.IsSame(slice1, slice2) {
		t.Fatal("Same(slice, slice) not same")
	}
}
