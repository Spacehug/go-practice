package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(input string, span int) (int64, error) {
	switch {
	case len(input) < span:
		return -1, errors.New("string is shorter than span")
	case span < 0:
		return -1, errors.New("negative span")
	}
	var max int64
	for index := 0; index <= len(input)-span; index++ {
		product := int64(1)
		for position := index; position < index+span; position++ {
			if input[position] < '0' || input[position] > '9' {
				return -1, errors.New("input contains alpha-characters")
			}
			number := int64(input[position] - '0')
			if number == 0 {
				product = 0
				break
			}
			product *= number
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}
