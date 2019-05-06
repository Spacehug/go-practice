package cryptosquare

import (
	"math"
	"unicode"
)

func Encode(input string) string {
	var norm = make([]rune, len(input))
	for _, r := range input {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			norm = append(norm, unicode.ToLower(r))
		}
	}
	r, c := sizeRowCol(len(norm))
	if c < 1 {return ""}
	ret := make([]rune, c * r + c - 1)
	for i := 0; i < c; i++ {
		for j := 0; j < r + 1; j++ {
			ix := j * c + i
			jx := i * r + j + i
			if len(ret) <= jx {continue}
			if len(norm) > ix {
				ret[jx] = norm[ix]
				continue
			}
			ret[jx] = ' '
		}
	}
	return string(ret)
}
func sizeRowCol(n int) (int, int) {
	col := int(math.Sqrt(float64(n)))
	row := col
	if col * row < n {
		col++
	}
	if col * row < n {
		row++
	}
	return row, col
}