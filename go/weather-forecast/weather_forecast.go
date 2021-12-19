// Package weather returns weather for a city.
package weather

// CurrentCondition stores the current condition.
var CurrentCondition string

// CurrentLocation stores the current location.
var CurrentLocation string

// Forecast returns the forecast given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
