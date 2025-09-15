package kata

func Solution(str string) []string {
	strContainer := ""
	result := []string{}

	if len(str)%2 != 0 {
		str += "_"
	}

	for _, s := range str {
		strContainer += string(s)

		if len(strContainer) == 2 {
			result = append(result, strContainer)
			strContainer = ""
		}
	}

	return result
}
