//From MIT lecture https://www.youtube.com/watch?v=Tw1k46ywN6E&t=0s&list=WL&index=46
//Return length of longest palindromic sub-sequence

package dynamicprogramming

var cache [][]cacheEntry

//Construct the cache and initialize it
func constructCache(length int) {

	cache = make([][]cacheEntry, length)
	for i := 0; i < length; i++ {
		cache[i] = make([]cacheEntry, length)
	}

}

type cacheEntry struct {
	lpssLength int
	lpss       []rune
}

func LongestPalindromicSubSequence(chars []rune, startIndex, endIndex int) (int, []rune) {

	//Construct the cache if first time
	if cache == nil {
		constructCache(len(chars))
	}

	//Return solution if it exists in the cache
	if cache[startIndex][endIndex].lpssLength != 0 {
		return cache[startIndex][endIndex].lpssLength, cache[startIndex][endIndex].lpss
	}

	//If start and end indices are same then longest palindromic sub-sequence is 1 long
	if startIndex == endIndex {
		cache[startIndex][endIndex].lpssLength = 1
		cache[startIndex][endIndex].lpss = []rune(string(chars[startIndex]))
		return cache[startIndex][endIndex].lpssLength, cache[startIndex][endIndex].lpss
	}

	//Both ends have same character
	if chars[startIndex] == chars[endIndex] {

		if startIndex+1 == endIndex {
			//Max length of 2
			cache[startIndex][endIndex].lpssLength = 2
			cache[startIndex][endIndex].lpss = []rune(string(string(chars[startIndex]) + string(chars[endIndex])))
		} else {
			//Recurse after taking off both ends
			length, lpss := LongestPalindromicSubSequence(chars, startIndex+1, endIndex-1)
			cache[startIndex][endIndex].lpssLength = 2 + length
			cache[startIndex][endIndex].lpss = append(append([]rune(string(chars[startIndex])), lpss...), chars[endIndex])
		}

	} else {

		//Disregard  oth ends and find longest sequence within
		lx, x := LongestPalindromicSubSequence(chars, startIndex+1, endIndex)
		ly, y := LongestPalindromicSubSequence(chars, startIndex, endIndex-1)

		if lx > ly {
			cache[startIndex][endIndex].lpssLength = lx
			cache[startIndex][endIndex].lpss = x
		} else {
			cache[startIndex][endIndex].lpssLength = ly
			cache[startIndex][endIndex].lpss = y
		}

	}

	return cache[startIndex][endIndex].lpssLength, cache[startIndex][endIndex].lpss

}
