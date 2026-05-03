package kata

// https://www.codewars.com/kata/58941fec8afa3618c9000184

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	currentHeight := 0
	days := 0

	for currentHeight < desiredHeight {
		days++
		currentHeight += upSpeed

		if currentHeight >= desiredHeight {
			break
		}
		currentHeight -= downSpeed
	}

	return days
}
