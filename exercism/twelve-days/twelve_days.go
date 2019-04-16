package twelve

import "fmt"

var dth = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var g = []string{"a Partridge in a Pear Tree", "two Turtle Doves, ", "three French Hens, ", "four Calling Birds, ", "five Gold Rings, ", "six Geese-a-Laying, ", "seven Swans-a-Swimming, ", "eight Maids-a-Milking, ", "nine Ladies Dancing, ", "ten Lords-a-Leaping, ", "eleven Pipers Piping, ", "twelve Drummers Drumming, "}

func Song() string {
	var song string
	for d := 1; d <= 12; d++ {
		song += Verse(d) + "\n"
	}
	return song
}

func Verse(d int) string {
	s := "On the %s day of Christmas my true love gave to me: %s."
	return fmt.Sprintf(s, dth[d - 1], getGifts(d - 1, ""))
}

func getGifts(d int, c string) string {
	switch {
	case d > 0:
		return getGifts(d - 1, c + g[d])
	case len(c) > 0:
		return c + "and " + g[d]
	default:
		return g[d]
	}
}
