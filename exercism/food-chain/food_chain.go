package foodchain

const TestVersion = 1

type lyric struct {
	name    string
	comment string
}

var song = []lyric{
	{"", ""},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her.\n"},
	{"bird", "How absurd to swallow a bird!\n"},
	{"cat", "Imagine that, to swallow a cat!\n"},
	{"dog", "What a hog, to swallow a dog!\n"},
	{"goat", "Just opened her throat and swallowed a goat!\n"},
	{"cow", "I don't know how she swallowed a cow!\n"},
	{"horse", "She's dead, of course!"},
}

func Verse(v int) string {
	if v < 1 || v > 8 {
		return "No Verse"
	}

	verse := "I know an old lady who swallowed a " + song[v].name + ".\n" + song[v].comment

	if v == 1 || v == 8 {
		return verse
	}

	for ; v > 1; v-- {
		verse += "She swallowed the " + song[v].name + " to catch the " + song[v-1].name
		if v == 3 {
			verse += " that wriggled and jiggled and tickled inside her"
		}
		verse += ".\n"
	}
	verse += song[1].comment
	return verse
}

func Verses(start int, end int) string {
	if start < 1 || start > end || end > 8 {
		return Verse(0)
	}
	verses := Verse(start)
	for start < end {
		start++
		verses += "\n\n" + Verse(start)
	}
	return verses
}

func Song() string {
	return Verses(1, 8)
}
