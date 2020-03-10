// Package bob does stuff
package bob

import "strings"

// Hey is a function that
func Hey(remark string) string {
	if strings.Index(remark, "?") == (len(remark) - 1) {
		return "Sure."
	}
	return "Whatever."
}
