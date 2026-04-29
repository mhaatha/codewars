// https://www.codewars.com/kata/56dbe0e313c2f63be4000b25

package kata

import (
	"testing"
)

func TestVertMirror(t *testing.T) {
	cases := []struct {
		Name        string
		Input, Want string
	}{
		{
			Name:  "first vertical mirror test case",
			Input: "abcd\nefgh\nijkl\nmnop",
			Want:  "dcba\nhgfe\nlkji\nponm",
		},
	}

	for _, cc := range cases {
		t.Run(cc.Name, func(t *testing.T) {
			got := VertMirror(cc.Input)

			if got != cc.Want {
				t.Errorf("got %s want %s", got, cc.Want)
			}
		})
	}
}
