package main

import (
	"fmt"
	"math"
)

func Sqrt(base, precision float64) float64 {
	guess := 1.0
	for previous_guess := 0.0; math.Abs(guess - previous_guess) > precision; guess -= (guess * guess - base) / (2 * guess) {
		previous_guess = guess
	}
	return guess
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2, 1e-6))
}
