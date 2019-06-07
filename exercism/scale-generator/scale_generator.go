package scale

import "strings"

func Scale(tonic string, interval string) (scale []string) {
	switch tonic {
	case "A", "a", "B", "b", "C", "c#", "D", "d#", "E", "e", "F#", "f#", "G", "g#":
		scale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	case "Ab", "Bb", "bb", "c", "Db", "d", "Eb", "eb", "F", "f", "Gb", "g":
		scale = []string{"Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G"}
	}
	tonic = strings.Title(tonic)
	for index, element := range scale {
		if element == tonic {
			scale = append(scale[index:], scale[:index]...)
			break
		}
	}

	if len(interval) == 0 {
		return scale
	}

	partialScale := make([]string, 0)
	offsets := map[string]int{"m": 1, "M": 2, "A": 3}
	position := 0
	for _, offset := range strings.Split(interval, "") {
		if shift, ok := offsets[offset]; ok {
			partialScale = append(partialScale, scale[position%len(scale)])
			position += shift
		}
	}

	return partialScale
}
