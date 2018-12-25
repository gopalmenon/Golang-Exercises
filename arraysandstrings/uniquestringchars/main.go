//Two functions to check if a string has unique characters
package uniquestringchars

import (
	"sort"
)

type runeSlice []rune

func (r runeSlice) Len() int           { return len(r) }
func (r runeSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runeSlice) Less(i, j int) bool { return r[i] < r[j] }

//Sort the input string and then scan it to see if any characters are repeated
func HasUniqueCharsBySorting(s string) bool {

	runes := runeSlice(s)
	sort.Sort(runes)

	var previousChar rune
	for index, char := range runes {
		if index != 0 {
			if previousChar == char {
				return false
			}
		}
		previousChar = char
	}
	return true
}

//Use a hash table to identify duplicates in a string
func HasUniqueCharsByHashing(s string) bool {

	runesMap := make(map[rune]bool)
	runesSlice := []rune(s)

	for _, char := range runesSlice {

		if _, ok := runesMap[char]; ok {
			return false
		}

		runesMap[char] = true

	}

	return true

}
