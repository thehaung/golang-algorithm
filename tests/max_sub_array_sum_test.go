package tests

import (
	"github.com/thehaung/golang-algorithm/solution"
	"testing"
)

func TestMaxSubArraySum(t *testing.T) {
	nums, slice, expectResult := 2, []int{1, 2, 3, 4, 5, 6}, 11

	if solution.MaxSubArraySum(slice, nums) != expectResult {
		t.Fail()
	}
}
