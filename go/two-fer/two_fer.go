// Package twofer contains a single method ShareWith
package twofer

// ShareWith returns a message in the format of one for name, one for you
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
