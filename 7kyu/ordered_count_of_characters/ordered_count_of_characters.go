package kata

import (
	"sort"
)

// https://www.codewars.com/kata/57a6633153ba33189e000074

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	res := []Tuple{}
	charOccured := make(map[rune]int)
	charsPosition := make(map[rune]int)
	positionCounter := 0

	for _, char := range text {
		if charOccured[char] == 0 {
			positionCounter++
			charOccured[char] = 1
			charsPosition[char] = positionCounter
		} else {
			charOccured[char]++
		}
	}

	orderedCharsPosition := rankByPosition(charsPosition)

	for _, charsPosition := range orderedCharsPosition {
		for key, value := range charOccured {
			if charsPosition.Key == key {
				res = append(res, Tuple{
					Char:  key,
					Count: value,
				})
			}
		}
	}

	return res
}

func rankByPosition(charsPosition map[rune]int) PairList {
	pairList := make(PairList, len(charsPosition))
	i := 0

	for key, value := range charsPosition {
		pairList[i] = Pair{key, value}
		i++
	}

	sort.Sort(pairList)

	return pairList
}

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func (pl PairList) Len() int           { return len(pl) }
func (pl PairList) Less(i, j int) bool { return pl[i].Value < pl[j].Value }
func (pl PairList) Swap(i, j int)      { pl[i], pl[j] = pl[j], pl[i] }
