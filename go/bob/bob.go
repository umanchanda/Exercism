// Package bob does stuff
package bob

import "unicode"

// Hey is a function that
func Hey(remark string) string {
	if string(remark[len(remark)-1]) == "?" {
		if IsUpper(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if IsUpper(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

// IsUpper returns true if a given string is all uppercase
func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}