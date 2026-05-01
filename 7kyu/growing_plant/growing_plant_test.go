package kata

import "testing"

func TestGrowingPlant(t *testing.T) {
	cases := []struct {
		Name                                    string
		UpSpeed, DownSpeed, DesiredHeight, Want int
	}{
		{
			Name:          "it should return 10",
			UpSpeed:       10,
			DownSpeed:     1,
			DesiredHeight: 910,
			Want:          10,
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
