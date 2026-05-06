package kata

import (
	"strconv"
)

// https://www.codewars.com/kata/683f5279c372b15f18a57d8e/train/go

func FloatToString(x float64) string {
	res := strconv.FormatFloat(x, 'f', -1, 64)

	return res
}
