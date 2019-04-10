package twofer

import "fmt"

func ShareWith(name string) string {
	switch {
	case len(name) > 0:
		return fmt.Sprintf("One for %s, one for me.", name)
	default:
		return fmt.Sprint("One for you, one for me.")
	}
}
