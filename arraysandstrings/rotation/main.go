package rotation

import "strings"

func IsSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}

func IsRotationOf(s, rotatedString string) bool {

	b := strings.Builder{}
	b.WriteString(s)
	b.WriteString(s)
	return IsSubstring(b.String(), rotatedString)

}
