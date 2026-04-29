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
		{
			Name:  "second vertical mirror test case",
			Input: "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu",
			Want:  "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw",
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

func TestHorMirror(t *testing.T) {
	cases := []struct {
		Name        string
		Input, Want string
	}{
		{
			Name:  "first horizontal mirror test case",
			Input: "abcd\nefgh\nijkl\nmnop",
			Want:  "mnop\nijkl\nefgh\nabcd",
		},
	}

	for _, cc := range cases {
		t.Run(cc.Name, func(t *testing.T) {
			got := HorMirror(cc.Input)

			if got != cc.Want {
				t.Errorf("got %s want %s", got, cc.Want)
			}
		})
	}
}
