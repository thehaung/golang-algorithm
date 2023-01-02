package solution

type SetInteger map[int]bool

func CountUniqueValue(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	uniqueMap := make(SetInteger)
	for _, val := range slice {
		if found, _ := uniqueMap[val]; !found {
			uniqueMap[val] = true
		}
	}
	return len(uniqueMap)
}
