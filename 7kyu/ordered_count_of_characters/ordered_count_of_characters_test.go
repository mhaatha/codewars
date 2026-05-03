package kata

import (
	"reflect"
	"testing"
)

// https://www.codewars.com/kata/57a6633153ba33189e000074

func TestOrderedCount(t *testing.T) {
	cases := []struct {
		Name  string
		Input string
		Want  []Tuple
	}{
		{
			Name:  "it should return a slice of Tuple where Char: 'a' Count: 5, Char: 'b' Count: 2, Char: 'r' Count: 2, Char: 'c' Count: 1, and Char: 'd' Count: 1",
			Input: "abracadabra",
			Want: []Tuple{
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
			},
		},
		{
			Name:  "it should return a slice of Tuple where Char: 'C' Count: 1, Char: 'o' Count: 1, Char: 'd' Count: 1, Char: 'e' Count: 1, Char: ' ' Count: 1, Char: 'W' Count: 1, Char: 'a' Count: 1, Char: 'r' Count: 1, and Char: 's' Count: 1",
			Input: "Code Wars",
			Want: []Tuple{
				{
					Char:  'C',
					Count: 1,
				},
				{
					Char:  'o',
					Count: 1,
				},
				{
					Char:  'd',
					Count: 1,
				},
				{
					Char:  'e',
					Count: 1,
				},
				{
					Char:  ' ',
					Count: 1,
				},
				{
					Char:  'W',
					Count: 1,
				},
				{
					Char:  'a',
					Count: 1,
				},
				{
					Char:  'r',
					Count: 1,
				},
				{
					Char:  's',
					Count: 1,
				},
			},
		},
	}

	for _, c := range cases {
		got := OrderedCount(c.Input)

		if !reflect.DeepEqual(got, c.Want) {
			t.Errorf("got %+v want %+v", got, c.Want)
		}
	}
}
