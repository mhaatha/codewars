package kata

// https://www.codewars.com/kata/598d91785d4ce3ec4f000018/train/go

import (
	"slices"
	"testing"
)

func TestNameValue(t *testing.T) {
	cases := []struct {
		Name  string
		Input []string
		Want  []int
	}{
		{
			Name:  "first test case",
			Input: []string{"abc", "abc", "abc", "abc"},
			Want:  []int{6, 12, 18, 24},
		},
	}

	for _, c := range cases {
		got := NameValue(c.Input)

		if !slices.Equal(got, c.Want) {
			t.Errorf("got %v, want %v", got, c.Want)
		}
	}
}
