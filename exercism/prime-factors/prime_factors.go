package prime

func Factors(number int64) []int64 {
	var primeFactors = make([]int64, 0)
	for number%2 == 0 {
		primeFactors = append(primeFactors, 2)
		number = number / 2
	}
	for i := int64(3); i*i <= number; i = i + 2 {
		for number%i == 0 {
			primeFactors = append(primeFactors, i)
			number = number / i
		}
	}
	if number > 2 {
		primeFactors = append(primeFactors, number)
	}
	return primeFactors
}
