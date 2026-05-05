package kata

// https://www.codewars.com/kata/57a6633153ba33189e000074

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	res := []Tuple{}
	charOccured := make(map[rune]int)
	charPosition := make(map[rune]int)
	positionCounter := 0

	for _, char := range text {
		if charOccured[char] == 0 {
			positionCounter++
			charOccured[char] = 1
			charPosition[char] = positionCounter
		} else {
			charOccured[char]++
		}
	}

	for key, index := range charPosition {
		for char, value := range charOccured {
			if key == char {
				res = append(res, Tuple{})
				copy(res[index+1:], res[index:])
				res[index] = Tuple{
					Char:  char,
					Count: value,
				}
			}
		}
	}

	return res
}
