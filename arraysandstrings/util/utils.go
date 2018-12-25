package util

import (
	"math/rand"
	"time"
)

const Length = 10000

var runes = []rune("一二三四五六七八九十1234567890")

func GenerateRandomRune(n int) string {
	randRune := make([]rune, n)

	for i := range randRune {
		// without this, the final value will be same all the time.
		rand.Seed(time.Now().UnixNano())

		randRune[i] = runes[rand.Intn(len(runes))]
	}
	return string(randRune)
}
