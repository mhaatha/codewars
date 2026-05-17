package kata

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/56f253dd75e340ff670002ac/train/go

func Compose(s1, s2 string) string {
	strng := ""

	splittedS1 := strings.Split(s1, "\n")
	splittedS2 := strings.Split(s2, "\n")

	s1Pointer := 1
	s2Pointer := len(splittedS2)

	for i := 0; i < len(splittedS1); i++ {
		currentWord := ""

		for j := 0; j < i+1; j++ {
			currentWord += string(splittedS1[i][j])
		}

		for k := len(splittedS2) - (i + 1); len(splittedS2) > (len(splittedS2) - i); k-- {
			fmt.Println(splittedS2[k])
		}

		strng += "\n"

		s1Pointer++
		s2Pointer--
	}

	return strng
}
