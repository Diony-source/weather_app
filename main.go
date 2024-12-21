package main

import (
	"fmt"
	"log"
	"os"
	"weather_app/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Welcome to the Weather App!")
	fmt.Println("1. CLI Mode")
	fmt.Println("2. Web Mode")
	fmt.Print("Choose a mode: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		handlers.StartCLI()
	case 2:
		fmt.Println("Starting web server on http://localhost:8080...")
		log.Fatal(handlers.StartWebServer())
	default:
		fmt.Println("Invalid choice. Exiting.")
		os.Exit(1)
	}
}
