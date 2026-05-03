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
	}

	for _, c := range cases {
		got := OrderedCount(c.Input)

		if !reflect.DeepEqual(got, c.Want) {
			t.Errorf("got %+v want %+v", got, c.Want)
		}
	}
}
