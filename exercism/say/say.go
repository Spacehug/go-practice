package say

import (
	"errors"
	"strings"
)

var verbatim = []string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

var scales = []struct {
	number int64
	name   string
}{
	{1e18, "quintillion"},
	{1e15, "quadrillion"},
	{1e12, "trillion"},
	{1e9, "billion"},
	{1e6, "million"},
	{1e3, "thousand"},
	{1e2, "hundred"},
}

func Say(number int64) (string, error) {
	switch {
	case number < 0, number >= 1e12:
		return "", errors.New("number out of bounds")
	case number < int64(len(verbatim)) && verbatim[number] != "":
		return verbatim[number], nil
	case number < 100:
		return verbatim[number-number%10] + "-" + verbatim[number%10], nil
	}
	parts := make([]string, 0)
	for _, scale := range scales {
		if number < scale.number {
			continue
		}
		said, _ := Say(number / scale.number)
		parts = append(parts, said, scale.name)
		number %= scale.number
	}
	if number != 0 {
		said, _ := Say(number)
		parts = append(parts, said)
	}
	return strings.Join(parts, " "), nil
}
