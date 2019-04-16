package sublist

type Relation string

func Sublist(s1, s2 []int) Relation {
	switch {
	case len(s1) < len(s2) && isSub(s1, s2):
		return "sublist"
	case len(s1) > len(s2) && isSub(s2, s1):
		return "superlist"
	case eq(s1, s2):
		return "equal"
	default:
		return "unequal"
	}
}

func eq(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func isSub(s1, s2 []int) bool {
	for i := 0; i <= len(s2) - len(s1); i++ {
		if eq(s1, s2[i:i + len(s1)]) {
			return true
		}
	}
	return false
}