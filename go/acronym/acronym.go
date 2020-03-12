// Package acronym contains a single method Abbreviate
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes a string and returns its abbreviation
func Abbreviate(s string) string {
	re, _ := regexp.Compile("[-_]")
	s = re.ReplaceAllString(s, " ")
	words := strings.Fields(s)
	abbreviation := ""
	for i := 0; i < len(words); i++ {
		abbreviation += string(words[i][0])
	}
	return strings.ToUpper(abbreviation)
}
