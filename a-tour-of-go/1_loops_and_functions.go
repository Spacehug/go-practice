package main

import (
  "fmt"
  "math"
)

func Sqrt(base, precision float64) float64 {
  guess := 1.0
  for stored := 0.0; math.Abs(guess - stored) > precision; guess -= (guess * guess - base) / (2 * guess) {
    stored = guess
  }
  return guess
}

func main() {
  fmt.Println(math.Sqrt(2))
  fmt.Println(Sqrt(2, 1e-6))
}
