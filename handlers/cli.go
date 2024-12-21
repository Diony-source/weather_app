package handlers

import (
	"fmt"
	"strings"
	"weather_app/services"
)

func StartCLI() {
	var city, country, unit string

	fmt.Print("Enter country code (e.g., TR for Turkey): ")
	fmt.Scanln(&country)

	fmt.Print("Enter city name: ")
	fmt.Scanln(&city)

	fmt.Print("Choose unit (C for Celsius, F for Fahrenheit): ")
	fmt.Scanln(&unit)
	unit = strings.ToLower(unit)
	if unit == "c" {
		unit = "metric"
	} else if unit == "f" {
		unit = "imperial"
	} else {
		fmt.Println("Invalid unit. Defaulting to Celsius.")
		unit = "metric"
	}

	location := fmt.Sprintf("%s,%s", city, country)
	weather, err := services.FetchWeather(location, unit)
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}

	fmt.Printf("\nWeather in %s:\n", weather.City)
	fmt.Printf("Temperature: %.2f°%s\n", weather.Temperature, strings.ToUpper(weather.Unit[:1]))
	fmt.Printf("Humidity: %d%%\n", weather.Humidity)
	fmt.Printf("Wind Speed: %.2f m/s\n", weather.WindSpeed)
	fmt.Printf("Description: %s\n", weather.Description)
}
