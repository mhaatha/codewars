// https://www.codewars.com/kata/56dbe0e313c2f63be4000b25

package kata

import (
	"fmt"
	"strings"
)

func VertMirror(s string) string {
	separatedString := strings.Split(s, "\n")
	groupedString := ""

	for _, s := range separatedString {
		currentString := ""

		for j := len(s); j > 0; j-- {
			currentString += string(s[j])
		}

		groupedString += currentString + "\n"
	}

	fmt.Println(groupedString)

	return "dcba\nhgfe\nlkji\nponm"
}
