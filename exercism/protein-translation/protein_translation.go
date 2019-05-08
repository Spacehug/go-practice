package protein

var ErrStop error
var ErrInvalidBase error

func FromCodon(input string) (string, error) {
	switch input {
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG", "CUU", "CUC", "CUA", "CUG":
		return "Leucine", nil
	case "AUU", "AUC", "AUA":
		return "Isoleucine", nil
	case "AUG":
		return "Methionine", nil
	case "GUU", "GUC", "GUA", "GUG":
		return "Valine", nil
	case "CCU", "CCC", "CCA", "CCG":
		return "Proline", nil
	case "ACU", "ACC", "ACA", "ACG":
		return "Threonine", nil
	case "GCU", "GCC", "GCA", "GCG":
		return "Alanine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "CAU", "CAC":
		return "Histidine", nil
	case "CAA", "CAG":
		return "Glutamine", nil
	case "AAU", "AAC":
		return "Asparagine", nil
	case "AAA", "AAG":
		return "Lysine", nil
	case "GAU", "GAC":
		return "Aspartic acid", nil
	case "GAA", "GAG":
		return "Glutamic acid", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "CGU", "CGC", "CGA", "CGG", "AGA", "AGG":
		return "Arginine", nil
	case "UCU", "UCC", "UCA", "UCG", "AGU", "AGC":
		return "Serine", nil
	case "GGU", "GGC", "GGA", "GGG":
		return "Glycin", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(input string) ([]string, error) {
	res := make([]string, len(input)/3)
	for sl, lim := input, 0; len(sl) > 0; sl = sl[3:] {
		if len(sl) < 3 {
			return res[:lim], ErrInvalidBase
		}
		aminoAcid, err := FromCodon(sl[:3])
		switch {
		case aminoAcid != "":
			res[lim] = aminoAcid
			lim++
		case err == ErrStop:
			return res[:lim], nil
		case err == ErrInvalidBase:
			return res[:lim], ErrInvalidBase
		}
	}
	return res, nil
}
