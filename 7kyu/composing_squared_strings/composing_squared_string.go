package kata

import (
	"strings"
)

// https://www.codewars.com/kata/56f253dd75e340ff670002ac/train/go

func Compose(s1, s2 string) string {
	strng := ""

	splittedS1 := strings.Split(s1, "\n")
	splittedS2 := strings.Split(s2, "\n")

	s2Pointer := len(splittedS2) - 1

	for i := range splittedS1 {
		currentWord := ""

		for j := 0; j < i+1; j++ {
			currentWord += string(splittedS1[i][j])
		}

		for k := 0; k <= len(splittedS2)-i-1; k++ {
			currentWord += string(splittedS2[s2Pointer][k])
		}

		strng += currentWord

		if i != len(splittedS1)-1 {
			strng += "\n"
		}

		s2Pointer--
	}

	return strng
}
