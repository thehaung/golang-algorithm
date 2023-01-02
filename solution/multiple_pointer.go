package solution

import "fmt"

/*
Write a function called sumZero which accepts a sorted array of integers.
The function should find the first pair where the sum is 0.
Return an array that includes both values that sum to zero or undefined if a pair does not exist
*/

func IsSumZero(slice []int) []int {
	result := make([]int, 0)
	leftPosition := 0
	rightPosition := len(slice) - 1

	for leftPosition < rightPosition {
		sum := slice[leftPosition] + slice[rightPosition]
		if sum == 0 {
			result = append(result, slice[leftPosition])
			result = append(result, slice[rightPosition])
			break
		} else if sum > 0 { // right position is larger than left
			rightPosition--
		} else {
			leftPosition++
		}
	}
	fmt.Printf("Result: %v\n", result)
	return result
}
