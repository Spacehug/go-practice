package summultiples

func SumMultiples(lim int, divisors ... int) (sum int) {
	for i := 0; i < lim; i++ {
		for _, number := range divisors {
			if number != 0 && i % number == 0 {
				sum += i
				break
			}
		}
	}
	return
}