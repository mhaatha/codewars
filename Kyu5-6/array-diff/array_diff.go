package kata

// Function harus menghapus semua element dari list A yang muncul pada list B
// Urutan output harus sama dengan urutan element A
// Exam: a = [1, 2] and b = [1], the result should be [2]
// Exam: a = [1, 2, 2, 2, 3] and b = [2], the result should be [1, 3]

func ArrayDiff(a, b []int) []int {
	m := make(map[int]int)
	output := []int{}

	for _, key := range b {
		m[key] = 1
	}

	for _, element := range a {
		if m[element] < 1 {
			output = append(output, element)
		}
	}

	return output
}
