package kata

import (
	"testing"
)

// https://www.codewars.com/kata/56f253dd75e340ff670002ac/train/go

func TestCompose(t *testing.T) {
	cases := []struct {
		Name, FirstString, SecondString, Want string
	}{
		{
			Name:         "first test case",
			FirstString:  "abcd\nefgh\nijkl\nmnop",
			SecondString: "qrst\nuvwx\nyz12\n3456",
			Want:         "a3456\nefyz1\nijkuv\nmnopq",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Compose(c.FirstString, c.SecondString)

			if got != c.Want {
				t.Errorf("got %s want %s", got, c.Want)
			}
		})
	}
}
