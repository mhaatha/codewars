package kata

// https://www.codewars.com/kata/58941fec8afa3618c9000184

import "testing"

func TestGrowingPlant(t *testing.T) {
	cases := []struct {
		Name                                    string
		UpSpeed, DownSpeed, DesiredHeight, Want int
	}{
		{
			Name:          "it should return 10",
			UpSpeed:       100,
			DownSpeed:     10,
			DesiredHeight: 910,
			Want:          10,
		},
		{
			Name:          "it should return 1",
			UpSpeed:       10,
			DownSpeed:     9,
			DesiredHeight: 4,
			Want:          1,
		},
		{
			Name:          "it should return 1",
			UpSpeed:       5,
			DownSpeed:     2,
			DesiredHeight: 5,
			Want:          1,
		},
		{
			Name:          "it should return 2",
			UpSpeed:       5,
			DownSpeed:     2,
			DesiredHeight: 6,
			Want:          2,
		},
		{
			Name:          "it should return 397",
			UpSpeed:       8,
			DownSpeed:     6,
			DesiredHeight: 800,
			Want:          397,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := GrowingPlant(c.UpSpeed, c.DownSpeed, c.DesiredHeight)

			if got != c.Want {
				t.Errorf("got %d want %d", got, c.Want)
			}
		})
	}
}
