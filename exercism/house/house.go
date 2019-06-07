package house

import (
	"fmt"
	"strings"
)

var verses = []struct {
	verb string
	noun string
}{
	{"lay in", "the house that Jack built"},
	{"ate", "the malt"},
	{"killed", "the rat"},
	{"worried", "the cat"},
	{"tossed", "the dog"},
	{"milked", "the cow with the crumpled horn"},
	{"kissed", "the maiden all forlorn"},
	{"married", "the man all tattered and torn"},
	{"woke", "the priest all shaven and shorn"},
	{"kept", "the rooster that crowed in the morn"},
	{"belonged to", "the farmer sowing his corn"},
	{"", "the horse and the hound and the horn"},
}

// Verse function. Given a number (1-12) returns a verse.
func Verse(n int) string {
	if n > len(verses) {
		return ""
	}
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("This is %s", verses[n-1].noun))
	for i := n - 2; i >= 0; i-- {
		b.WriteString(fmt.Sprintf("\nthat %s %s", verses[i].verb, verses[i].noun))
	}
	b.WriteString(".")
	return b.String()
}

// Song function. Returns all verses of the Song.
func Song() string {
	b := strings.Builder{}
	b.WriteString(Verse(1))
	for i := 2; i <= len(verses); i++ {
		b.WriteString("\n\n")
		b.WriteString(Verse(i))
	}
	return b.String()
}
