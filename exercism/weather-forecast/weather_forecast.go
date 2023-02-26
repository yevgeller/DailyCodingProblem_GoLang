// Package weather provides weather forecast.
package weather

// CurrentCondition is a variable that defines the current condition.
var CurrentCondition string

// CurrentLocation is a variable that defines the current location.
var CurrentLocation string

//
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
