package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.0f", e)
}

func SqrtWithErrorControl(base, precision float64) (float64, error) {
	if base < 0 {
		return 0, ErrNegativeSqrt(base)
	}
	guess := 1.0
	for stored := 0.0; math.Abs(guess-stored) > precision; guess -= (guess*guess - base) / (2 * guess) {
		stored = guess
	}
	return guess, nil
}

func main() {
	fmt.Println(SqrtWithErrorControl(2, 1e-6))
	fmt.Println(SqrtWithErrorControl(-2, 1e-6))
}
