// Package twelve does
package twelve

var gifts = [12]string{
	"and a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

var days = [12]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

// Song returns the entire song
func Song() string {
	song := "On the " + days[0] + " day of Christmas my true love gave to me: a Partridge in a Pear Tree.\n"
	for i := 2; i <=11; i++ {
		song += "On the " + days[i-1] + " day of Christmas my true love gave to me: " + Verse(i) + ".\n"
	}
	song += "On the " + days[11] + " day of Christmas my true love gave to me: " + Verse(12) + "."
	return song
}

// Verse generates the presents for a given day
func Verse(n int) string {
	if n == 1 {
		return gifts[0]
	}
	return gifts[n-1] + ", " + Verse(n-1)
}