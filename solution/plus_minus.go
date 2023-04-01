package solution

import (
	"fmt"
	"strconv"
)

func PlusMinus(arr []int32) {
	n := float64(len(arr))

	negativeNums, zeroNums, postiveNums := getTypeOfNums(arr)
	fmt.Println(strconv.FormatFloat(float64(postiveNums)/n, 'f', 6, 64))
	fmt.Println(strconv.FormatFloat(float64(negativeNums)/n, 'f', 6, 64))
	fmt.Println(strconv.FormatFloat(float64(zeroNums)/n, 'f', 6, 64))
}

func getTypeOfNums(arr []int32) (int32, int32, int32) {
	var negativeNums, zeroNums, postiveNums int32
	postiveNums, negativeNums, zeroNums = 0, 0, 0

	for _, val := range arr {
		if val == 0 {
			zeroNums++
		} else if val > 0 {
			postiveNums++
		} else {
			negativeNums++
		}
	}

	return negativeNums, zeroNums, postiveNums
}
