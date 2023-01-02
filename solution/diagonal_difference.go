package solution

func DiagonalDifference(arr [][]int) int {
	cursor, leftSum, rightSum := 0, 0, 0

	for i, row := range arr {
		if cursor >= len(row) {
			break
		}
		leftSum += arr[i][cursor]
		rightSum += arr[i][len(row)-cursor-1] // out of index
		cursor++
	}
	return absInt(leftSum, rightSum)
}

func absInt(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}
