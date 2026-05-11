package kata

import "testing"

// https://www.codewars.com/kata/56a4872cbb65f3a610000026/train/go

func TestMaxRot(t *testing.T) {
	cases := []struct {
		Name        string
		Input, Want int
	}{
		{
			Name:  "it should return 68957",
			Input: 56789,
			Want:  68957,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := MaxRot(c.Input)

			if got != c.Want {
				t.Errorf("got %d want %d", got, c.Want)
			}
		})
	}
}
