package solution

func IsSame(slice1, slice2 []int) bool {

	mapSlice1 := make(map[int]int)
	mapSlice2 := make(map[int]int)

	if len(slice1) != len(slice2) {
		return false
	}

	increaseMapSlice(slice1, mapSlice1)
	increaseMapSlice(slice2, mapSlice2)

	for key, val := range mapSlice1 {
		quantity, found := mapSlice2[key*key]
		if !found {
			return false
		}

		if quantity != val {
			return false
		}
	}

	return true
}

func increaseMapSlice(sliece []int, mapSlice map[int]int) {
	for _, key := range sliece {
		if _, found := mapSlice[key]; found {
			mapSlice[key] += 1
		} else {
			mapSlice[key] = 0
		}
	}
}
