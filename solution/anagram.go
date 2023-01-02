package solution

/*
Given two strings, write a function to determine if the second string is an anagram of the first.
An anagram is a word, phrase, or name formed by rearranging the letters of another,
such as cinema, formed from iceman.
*/

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) || str1 == "" || str2 == "" {
		return false
	}

	mapStr1 := make(map[byte]int)
	mapStr2 := make(map[byte]int)

	initAnagram(str1, mapStr1)
	initAnagram(str2, mapStr2)

	for key, val := range mapStr1 {
		if _, found := mapStr2[key]; !found {
			return false
		}

		if mapStr2[key] != val {
			return false
		}
	}
	return true
}

func initAnagram(str string, capture map[byte]int) {
	for i := 0; i < len(str); i++ {
		if _, found := capture[str[i]]; !found {
			capture[str[i]] = 0
		} else {
			capture[str[i]] = capture[str[i]] + 1
		}
	}
}
