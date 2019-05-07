package palindrome

import (
	"fmt"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fMin, fMax int) (pMin, pMax Product, err error) {
	var product int
	if fMin > fMax {
		err = fmt.Errorf("fmin > fmax (%d, %d)", fMin, fMax)
		return
	}
	for number1 := fMin; number1 <= fMax; number1++ {
		for number2 := number1; number2 <= fMax; number2++ {
			product = number1 * number2
			if pMin.Product != 0 && product > pMin.Product || !isPalindrome(product) {
				break
			}
			if product == pMin.Product {
				pMin.Factorizations = append(pMin.Factorizations, [2]int{number1, number2})
			}
			if pMin.Product == 0 || product < pMin.Product {
				pMin = Product{product, [][2]int{{number1, number2}}}
			}
		}
	}
	for number1 := fMax; number1 >= fMin; number1-- {
		for number2 := number1; number2 >= fMin; number2-- {
			product = number1 * number2
			if product < pMax.Product {
				break
			}
			if isPalindrome(product) {
				if product == pMax.Product {
					pMax.Factorizations = append(pMax.Factorizations, [2]int{number2, number1})
				}
				if product > pMax.Product {
					pMax = Product{product, [][2]int{{number2, number1}}}
				}
			}
		}
	}
	if pMin.Product == 0 {
		err = fmt.Errorf("no palindromes")
		return
	}
	return
}

func isPalindrome(product int) bool {
	var reverse int
	for number := product; number != 0; number /= 10 {
		reverse = reverse*10 + number%10
	}
	return reverse == product
}
