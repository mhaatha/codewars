package kata

import "testing"

// Function harus memisahkan string menjadi 2 pasang karakter
// Kalau string ganjil, pasangan terakhir harus diganti dengan underscore ('_').
// Exam: 'abc' => ['ab', 'c_']
// Exam: 'abcdef' => ['ab', 'cd', 'ef']

func TestSolution(t *testing.T) {
	cases := []struct {
		Name string
		Str  string
		Want []string
	}{
		{
			Name: "the output should be ['ab', 'c_']",
			Str:  "abc",
			Want: []string{"ab", "c_"},
		},
		{
			Name: "the output should be ['ab', 'cd', 'ef']",
			Str:  "abcdef",
			Want: []string{"ab", "cd", "ef"},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := Solution(c.Str)

			if !compareSlice(got, c.Want) {
				t.Errorf("got %v want %v", got, c.Want)
			}
		})
	}
}

func compareSlice(a, b []string) bool {
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
