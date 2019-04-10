package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	var result strings.Builder
	newString := strings.ReplaceAll(s, "-", " ")
	for _, word := range strings.Fields(newString) {
		result.WriteRune(rune(word[0]))
	}
	return strings.ToUpper(result.String())
}
