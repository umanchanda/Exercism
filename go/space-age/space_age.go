// Package space contains a single method Age
package space

// Planet is a type that defines a Planet name
type Planet string

// Age returns the age in years given a number of seconds on a given planet
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		return seconds / (365.24 * 3600 * 24 * 0.2408467)
	case "Venus":
		return seconds / (365.24 * 3600 * 24 * 0.61519726)
	case "Earth":
		return seconds / (365.24 * 3600 * 24)
	case "Mars":
		return seconds / (365.24 * 3600 * 24 * 1.8808158)
	case "Jupiter":
		return seconds / (365.24 * 3600 * 24 * 11.862615)
	case "Saturn":
		return seconds / (365.24 * 3600 * 24 * 29.447498)
	case "Uranus":
		return seconds / (365.24 * 3600 * 24 * 84.016846)
	case "Neptune":
		return seconds / (365.24 * 3600 * 24 * 164.79132)
	default:
		return 0.0
	}
}
