package railfence

func Encode(input string, rails int) string {
	return Swap(input, rails, 1)
}

func Decode(input string, rails int) string {
	return Swap(input, rails, -1)
}

func Swap(input string, rails int, action int) string {
	var insert, index, delta int
	var rPos, sPos *int
	switch action {
	case 1:
		rPos = &index
		sPos = &insert
	default:
		rPos = &insert
		sPos = &index
	}
	cycle := (rails - 1) * 2
	result := []byte(input)
	for rail := 0; rail < rails; rail++ {
		delta = rail * 2
		for insert = rail; insert < len(input); insert += delta {
			result[*rPos] = input[*sPos]
			index++
			if delta == cycle {
				continue
			}
			delta = cycle - delta
		}
	}
	return string(result)
}
