package accumulate

func Accumulate(inp []string, con func(string) string) []string {
	var res = make([]string, len(inp))
	for i, f := range inp {
		res[i] = con(f)
	}
	return res
}
