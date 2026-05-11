package kata

import (
	"slices"
	"strconv"
	"strings"
)

// https://www.codewars.com/kata/56a4872cbb65f3a610000026/train/go

func MaxRot(n int) int {
	splittedN := strings.Split(strconv.Itoa(n), "")
	biggest, _ := strconv.Atoi(strings.Join(splittedN, ""))

	for i := range splittedN {
		if i == len(splittedN)-1 {
			break
		}
		splittedN = append(splittedN, splittedN[i])
		splittedN = slices.Delete(splittedN, i, i+1)
		currentValue, _ := strconv.Atoi(strings.Join(splittedN, ""))
		if biggest < currentValue {
			biggest = currentValue
		}
	}

	return biggest
}
