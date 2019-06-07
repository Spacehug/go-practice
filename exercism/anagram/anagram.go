package anagram

import (
	"strings"
	"unicode"
)

type counts [26]int

func count(s string) (c counts) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			c[unicode.ToLower(r)-'a']++
		}
	}
	return
}

func Detect(subject string, candidates []string) (sublist []string) {
	sc := count(subject)
	for _, c := range candidates {
		if len(subject) != len(c) || strings.EqualFold(subject, c) {
			continue
		}
		cc := count(c)
		if sc == cc {
			sublist = append(sublist, c)
		}
	}
	return
}
