package robotname

import (
	"errors"
	"math/rand"
)

const (
	namePattern = "AA###"
	ln = 26 * 26 * 10 * 10 * 10
)

var reg []string
var ns = shuffleNameList(genSequenceOfIntegers())
var nn = nextNameInt()


type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if len(reg) >= ln {return "", errors.New("namespace exhausted")}
	if len(r.name) == 0 {
		r.name = genName()
		reg = append(reg, r.name)
	}
	return r.name, nil
}

func (r *Robot) Reset() {r.name = ""}

func genSequenceOfIntegers() []int {
	result := make([]int, ln)
	for i := 0; i < len(result); i++ {result[i] = i}
	return result
}

func shuffleNameList(l []int) []int {
	rand.Shuffle(ln, func(i, j int) {l[i], l[j] = l[j], l[i]})
	return l
}

func nextNameInt() func() int {
	i := -1
	return func() int {
		i++
		return ns[i]
	}
}

func genName() string {
	p := nn()
	n := make([]byte, len(namePattern))
	for i, c := range namePattern {
		switch c {
		case 'A':
			n[i] = byte('A' + p % 26)
			p /= 26
		case '#':
			n[i] = byte('0' + p % 10)
			p /= 10
		default:
			panic("invalid name pattern")
		}
	}
	return string(n)
}
