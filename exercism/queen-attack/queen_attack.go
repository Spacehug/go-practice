package queenattack

import (
	"errors"
)

func isValidSquare(s string) bool {
	switch {
	case len(s) != 2,
	s[0] < 'a',
	s[0] > 'h',
	s[1] < '1',
	s[1] > '8':
		return false
	default:
		return true
	}
}

func CanQueenAttack(w, b string) (attack bool, err error) {
	errW, errB := isValidSquare(w), isValidSquare(b)
	switch {
	case !errW, !errB, w == b:
		return false, errors.New("invalid inputs")
	case w[0] == b[0],
	w[1] == b[1],
	w[0] - b[0] == w[1] - b[1],
	w[0] - b[0] + w[1] - b[1] == 0:
		return true, nil
	default:
		return false, nil
	}
}