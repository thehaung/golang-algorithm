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

func MaxSubArraySum(slice []int, nums int) int {
	maxSum, tempSum := 0, 0
	if len(slice) < nums {
		return 0
	}

	for i := 0; i < nums; i++ {
		maxSum += slice[i]
	}

	tempSum = maxSum
	for i := nums; i < len(slice); i++ {
		tempSum = tempSum - slice[i-nums] + slice[i]
		maxSum = maxInt(maxSum, tempSum)
	}
	fmt.Printf("maxSum: %v\n", maxSum)
	return maxSum
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
