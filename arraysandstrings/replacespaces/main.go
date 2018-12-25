package replacespaces

func ReplaceSpaces(s string) string {

	spaceReplacement := []rune("%20")
	spaceReplacementLen := len(spaceReplacement)
	runes := []rune(s)

	for index, char := range runes {
		if char == ' ' {
			copy(runes[index+spaceReplacementLen:], runes[index+1:len(runes)-2])
			copy(runes[index:index+spaceReplacementLen], spaceReplacement)
		}
	}

	return string(runes)

}
