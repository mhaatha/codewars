package kata

import (
	"math"
)

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	firstOp := float64(desiredHeight) / float64(upSpeed)
	secondOp := firstOp / float64(downSpeed)
	thirdOp := firstOp + secondOp

	if thirdOp < 1 {
		return int(math.Ceil(thirdOp))
	}

	frac := thirdOp - math.Floor(thirdOp)

	if frac == 0.5 {
		return int(math.Floor(thirdOp))
	}

	return int(math.Round(thirdOp))
}
