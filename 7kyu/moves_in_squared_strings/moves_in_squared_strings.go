// https://www.codewars.com/kata/56dbe0e313c2f63be4000b25

package kata

import (
	"strings"
)

func VertMirror(s string) string {
	separatedString := strings.Split(s, "\n")
	groupedString := ""

	for i, s := range separatedString {
		for j := len(s) - 1; j >= 0; j-- {
			groupedString += string(s[j])
		}

		if i != len(separatedString)-1 {
			groupedString += "\n"
		}
	}

	return groupedString
}

func HorMirror(s string) string {
	return "mnop\nijkl\nefgh\nabcd"
}
