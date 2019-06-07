package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	for i, r := range s {
		if unicode.IsLetter(r) && strings.ContainsRune(s[i+1:], r) {
			return false
		}
	}
	return true
}
