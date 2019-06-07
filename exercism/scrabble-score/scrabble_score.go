package scrabble

func Score(word string) (score int) {
	for _, c := range word {
		switch c {
		case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u', 'L', 'l', 'N', 'n', 'R', 'r', 'S', 's', 'T', 't':
			score += 1
		case 'D', 'd', 'G', 'g':
			score += 2
		case 'B', 'b', 'C', 'c', 'M', 'm', 'P', 'p':
			score += 3
		case 'F', 'f', 'H', 'h', 'V', 'v', 'W', 'w', 'Y', 'y':
			score += 4
		case 'K', 'k':
			score += 5
		case 'J', 'j', 'X', 'x':
			score += 8
		case 'Q', 'q', 'Z', 'z':
			score += 10
		default:
			score += 0
		}
	}
	return
}
