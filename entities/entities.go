package entities

// WeatherData represents the structure of weather information
type WeatherData struct {
	City        string
	Temperature float64
	Humidity    int
	WindSpeed   float64
	Description string
	Unit        string // "C" for Celsius, "F" for Fahrenheit
}
