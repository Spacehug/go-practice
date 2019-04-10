package hamming

import (
	"errors"
)

func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("strains are of different length")
	}
	for index := 0; len(a) > index; index++ {
		if a[index] != b[index] {
			distance++
		}
	}
	return
}
