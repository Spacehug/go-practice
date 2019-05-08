package wordcount

import "strings"
import "unicode"

type Frequency map[string]int

func WordCount(input string) Frequency {
	var res = Frequency{}
	for _, s := range strings.FieldsFunc(strings.ToLower(input), func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	}) {
		res[strings.Trim(s, "'")]++
	}
	return res
}
