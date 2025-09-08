package kata

import (
	"testing"
)

// Function harus menghapus semua element dari list A yang muncul pada list B
// Urutan output harus sama dengan urutan element A
// Exam: a = [1, 2] and b = [1], the result should be [2]
// Exam: a = [1, 2, 2, 2, 3] and b = [2], the result should be [1, 3]

func TestArrDiff(t *testing.T) {
	cases := []struct {
		Name        string
		FirstSlice  []int
		SecondSlice []int
		Want        []int
	}{
		{
			Name:        "the result should be [2]",
			FirstSlice:  []int{1, 2},
			SecondSlice: []int{1},
			Want:        []int{2},
		},
		{
			Name:        "the result should be [1, 3]",
			FirstSlice:  []int{1, 2, 2, 2, 3},
			SecondSlice: []int{2},
			Want:        []int{1, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := ArrayDiff(c.FirstSlice, c.SecondSlice)

			if !compareSlice(got, c.Want) {
				t.Errorf("got %v want %v", got, c.Want)
			}
		})
	}
}

func compareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
