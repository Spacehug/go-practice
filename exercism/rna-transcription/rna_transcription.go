package strand

var mapping = map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}

func ToRNA(dna string) string {
	var result = []rune(dna)
	for index, r := range dna {
		result[index] = mapping[r]
	}
	return string(result)
}
