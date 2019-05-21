package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(function func(int) bool) (output Ints) {
	for _, value := range ints {
		if function(value) {
			output = append(output, value)
		}
	}
	return
}

func (ints Ints) Discard(function func(int) bool) Ints {
	return ints.Keep(func(number int) bool { return !function(number) })
}

func (lists Lists) Keep(function func([]int) bool) (output Lists) {
	for _, value := range lists {
		if function(value) {
			output = append(output, value)
		}
	}
	return
}

func (strings Strings) Keep(function func(string) bool) (output Strings) {
	for _, value := range strings {
		if function(value) {
			output = append(output, value)
		}
	}
	return
}
