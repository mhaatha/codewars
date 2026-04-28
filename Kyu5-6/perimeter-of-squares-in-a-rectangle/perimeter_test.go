package kata

import "testing"

func TestPerimeter(t *testing.T) {
	cases := []struct {
		Name   string
		Square int
		Want   int
	}{
		{
			Name:   "should return 80",
			Square: 5,
			Want:   80,
		},
		{
			Name:   "should return 216",
			Square: 7,
			Want:   216,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Square(c.Square)

			if got != c.Want {
				t.Errorf("got %v want %v", got, c.Want)
			}
		})
	}
}
