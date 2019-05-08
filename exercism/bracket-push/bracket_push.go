package brackets

var pair = map[rune]rune{'}': '{', ')': '(', ']': '['}

func Bracket(input string) bool {
	stack := make([]rune, 0)
	for _, r := range input {
		switch r {
		case '[', '{', '(':
			stack = append(stack, r)
		case ']', '}', ')':
			if len(stack) == 0 || stack[len(stack)-1] != pair[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
