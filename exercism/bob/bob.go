package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {

	endsWithQuestion := strings.HasSuffix(strings.TrimSpace(remark), "?")
	yelledAt := strings.ToUpper(remark) == remark && strings.IndexFunc(remark, unicode.IsLetter) >= 0
	hasNothing := len(strings.Fields(remark)) == 0

	switch {
	case yelledAt && endsWithQuestion:
		return "Calm down, I know what I'm doing!"
	case endsWithQuestion:
		return "Sure."
	case yelledAt:
		return "Whoa, chill out!"
	case hasNothing:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

	return ""
}
