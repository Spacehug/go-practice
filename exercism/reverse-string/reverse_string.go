package reverse

import "bytes"

func String(s string) string {
	var b bytes.Buffer
	rs := []rune(s)
	for i := len(rs) - 1; i >= 0; i-- {b.WriteRune(rs[i])}
	return b.String()
}