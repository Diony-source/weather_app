package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"weather_app/entities"
)

const apiURL = "http://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&appid=%s"

// FetchWeather fetches the weather data from OpenWeatherMap API
func FetchWeather(location, unit string) (*entities.WeatherData, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API key is missing")
	}

	url := fmt.Sprintf(apiURL, location, unit, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("could not parse response: %w", err)
	}

	weather := &entities.WeatherData{
		City:        data["name"].(string),
		Temperature: data["main"].(map[string]interface{})["temp"].(float64),
		Humidity:    int(data["main"].(map[string]interface{})["humidity"].(float64)),
		WindSpeed:   data["wind"].(map[string]interface{})["speed"].(float64),
		Description: data["weather"].([]interface{})[0].(map[string]interface{})["description"].(string),
		Unit:        unit,
	}

	return weather, nil
}
