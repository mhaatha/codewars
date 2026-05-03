package kata

// https://www.codewars.com/kata/57a6633153ba33189e000074

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	res := []Tuple{}
	charOccured := make(map[rune]int)

	for _, char := range text {
		if charOccured[char] == 0 {
			charOccured[char] = 1
		} else {
			charOccured[char]++
		}
	}

	for key, value := range charOccured {
		res = append(res, Tuple{
			Char:  key,
			Count: value,
		})
	}

	return res
}
