package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"weather_app/services"
)

func StartWebServer() error {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/weather", weatherHandler)
	return http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1>Welcome to the Weather App</h1>`)
	fmt.Fprintln(w, `<form action="/weather" method="get">`)
	fmt.Fprintln(w, `Country Code: <input type="text" name="country" placeholder="e.g., TR"><br>`)
	fmt.Fprintln(w, `City: <input type="text" name="city" placeholder="e.g., Istanbul"><br>`)
	fmt.Fprintln(w, `Unit (C/F): <input type="text" name="unit" placeholder="e.g., C"><br>`)
	fmt.Fprintln(w, `<button type="submit">Get Weather</button>`)
	fmt.Fprintln(w, `</form>`)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	city := r.URL.Query().Get("city")
	unit := r.URL.Query().Get("unit")
	unit = strings.ToLower(unit)
	if unit == "c" {
		unit = "metric"
	} else if unit == "f" {
		unit = "imperial"
	}

	location := fmt.Sprintf("%s,%s", city, country)
	weather, err := services.FetchWeather(location, unit)
	if err != nil {
		http.Error(w, "Could not fetch weather: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
