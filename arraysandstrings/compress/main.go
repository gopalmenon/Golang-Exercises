package compress

import (
	"bytes"
	"strconv"
	"strings"
)

func compress(s string) string {

	runes := []rune(s)

	var previousChar = rune(' ')
	charCount, prevIndex := 1, 0
	var buffer bytes.Buffer
	for index, char := range runes {
		if index != 0 {
			if char == previousChar {
				charCount += 1
			} else {
				if charCount > 1 {
					buffer.WriteString(strings.Join([]string{string(runes[prevIndex]), strconv.Itoa(charCount)}, ""))
				} else {
					buffer.WriteString(string(runes[prevIndex]))
				}
				previousChar = char
				prevIndex = index
				charCount = 1
			}
		}
		previousChar = char
	}

	if charCount > 1 {
		buffer.WriteString(strings.Join([]string{string(runes[prevIndex]), strconv.Itoa(charCount)}, ""))
	} else {
		buffer.WriteString(string(runes[prevIndex]))
	}

	return buffer.String()

}
