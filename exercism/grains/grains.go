package grains

import (
	"errors"
)

func Total() uint64 {
	return (1 << 64) - 1
}

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("integer out of bounds")

	}
	return 1 << uint64(n-1), nil
}
