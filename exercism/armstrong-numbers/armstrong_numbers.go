package armstrong

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func IsNumber(number int) bool {
	feedSum, feedLength := number, number
	length, sum := 0, 0
	for feedLength > 0 {
		feedLength /= 10
		length += 1
	}
	for feedSum > 0 {
		sum += Pow(feedSum%10, length)
		feedSum /= 10
	}
	return number == sum
}
