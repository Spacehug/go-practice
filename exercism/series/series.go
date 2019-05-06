package series

func All(n int, s string) []string {
	var res = make([]string, len(s) - n + 1)
	for i := 0; i <= len(s) - n; i++ {res[i] = s[i:i + n]}
	return res
}

func UnsafeFirst(n int, s string) string {
	res, ok := First(n, s)
	switch ok {
	case true:
		return res
	default:
		panic(nil)
	}
}

func First(n int, s string) (first string, ok bool) {
	switch {
	case n > len(s), n < 1:
		return s, false
	default:
		return s[0:n], true
	}
}
