package reversestring

func reverse(s string) string {

	runes := []rune(s)
	reversedString := make([]rune, len(s))
	inputLength := len(s)

	for index, char := range runes {
		reversedString[inputLength-1-index] = char
	}

	return string(reversedString)

}

func improvedReverse(s string) string {

	runes := []rune(s)
	length := len(runes)

	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}
