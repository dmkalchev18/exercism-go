// Package weather forecasts the current weather condition of various cities in Goblinocus.
package weather

var (
// CurrentCondition represents the current condition of the weather.
	CurrentCondition string
// CurrentLocation represents the current location of where the weather will be measured.
	CurrentLocation  string
)

// Forecast returns the data of what the current weather condition is by provided location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
