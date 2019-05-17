package dna

import "errors"

type Histogram map[rune]int
type DNA string

func (d *DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, r := range *d {
		switch r {
		case 'A', 'C', 'G', 'T':
			h[r]++
		default:
			return h, errors.New("invalid nucleotide")
		}
	}
	return h, nil
}
