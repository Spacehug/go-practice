package main

import (
	"fmt"
	"math/big"
)

func Fibonacci() func() *big.Int {
	last, current := big.NewInt(0), big.NewInt(1)
	return func() *big.Int {
		current.Add(current, last)
		last, current = current, last
		return current
	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 1000; i++ {
		fmt.Println(i+1, f())
	}
}
