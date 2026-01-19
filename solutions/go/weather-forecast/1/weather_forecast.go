// Package weather provides basic utilities for managing and reporting weather conditions for specific geographic locations. It allows for the tracking of a global state regarding the current location and its associated weather status.
package weather

var (
    // CurrentCondition variable that stores the most recently updated weather state (e.g., "Sunny", "Rainy").
	CurrentCondition string

    // CurrentLocation variable that stores the name of the city or region currently being tracked.
	CurrentLocation  string
)

// Forecast function updates the package's global state and returns a formatted summary of the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
