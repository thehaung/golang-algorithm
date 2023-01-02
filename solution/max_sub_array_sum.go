package solution

import (
	"fmt"
	"math"
)

/*
Write a function called maxSubarraySum which accepts an array of integers and a number called n.
The function should calculate the maximum sum of n consecutive elements in the array.
*/

func MaxSubArraySumNaive(slice []int, nums int) int {
	sum, startPosition, maxSum := 0, 0, math.MinInt

	for len(slice) >= startPosition+nums {
		for i := startPosition; i < startPosition+nums; i++ {
			sum += slice[i]
		}

		if maxSum < sum {
			maxSum = sum
		}
		sum = 0
		startPosition++
	}
	fmt.Printf("maxSum: %v\n", maxSum)
	return maxSum
}
