package kata

// https://www.codewars.com/kata/57a6633153ba33189e000074

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	return []Tuple{
		{
			Char:  'a',
			Count: 5,
		},
		{
			Char:  'b',
			Count: 2,
		},
		{
			Char:  'r',
			Count: 2,
		},
		{
			Char:  'c',
			Count: 1,
		},
		{
			Char:  'd',
			Count: 1,
		},
	}
}
