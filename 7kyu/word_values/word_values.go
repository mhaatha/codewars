package kata

// https://www.codewars.com/kata/598d91785d4ce3ec4f000018/train/go

func NameValue(my_list []string) []int {
	alphabetHashMapValue := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}

	result := make([]int, len(my_list))

	for i, word := range my_list {
		var currentWordValue int
		for j := 0; j < len(word); j++ {
			currentWordValue += alphabetHashMapValue[string(word[j])]
		}

		result[i] = currentWordValue * (i + 1)
	}

	return result
}
