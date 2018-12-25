package permutation

import "sort"

type runeSlice []rune

func (r runeSlice) Len() int           { return len(r) }
func (r runeSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runeSlice) Less(i, j int) bool { return r[i] < r[j] }

func IsPermutationOfOtherUsingSort(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	rs1 := runeSlice(s1)
	rs2 := runeSlice(s2)
	sort.Sort(rs1)
	sort.Sort(rs2)

	for index, char := range rs1 {
		if char != rs2[index] {
			return false
		}
	}

	return true
}

func IsPermutationOfOtherUsingHashing(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	rs1 := runeSlice(s1)
	rs2 := runeSlice(s2)

	charCounts1 := make(map[rune]int, len(rs1))
	charCounts2 := make(map[rune]int, len(rs2))

	for index, char := range rs1 {
		charCounts1[char] += 1
		charCounts2[rs2[index]] += 1
	}

	for char, count1 := range charCounts1 {
		if count2, ok := charCounts2[char]; ok {
			if count1 != count2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
