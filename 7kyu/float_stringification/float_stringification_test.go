package kata

import "testing"

// https://www.codewars.com/kata/683f5279c372b15f18a57d8e/train/go

func TestFloatToString(t *testing.T) {
	cases := []struct {
		Name  string
		Input float64
		Want  string
	}{
		{
			Name:  "first test case",
			Input: 123.456,
			Want:  "123.456",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := FloatToString(c.Input)

			if got != c.Want {
				t.Errorf("got %s want %s", got, c.Want)
			}
		})
	}
}
