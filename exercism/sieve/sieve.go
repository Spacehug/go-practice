package sieve

func Sieve(lim int) []int {
	marked := make([]bool, lim+1)
	primes := make([]int, lim/2)
	x := 0
	for index := 2; index <= lim; index++ {
		if marked[index] {
			continue
		}
		primes[x] = index
		for m := index << 1; m <= lim; m += index {
			marked[m] = true
		}
		x++
	}
	return primes[:x]
}
