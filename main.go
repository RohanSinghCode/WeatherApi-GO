package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvFile()
	fmt.Println("Hello Welcome To Weather Console App!!!!")
	for {
		menu()
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter the pincode:")
			var pincode string
			fmt.Scan(&pincode)
			client := NewWeatherClient(os.Getenv("OPENWEATHER_BASE_URL"), os.Getenv("OPENWEATHER_APP_ID"))
			response, err := client.GetWeatherByPincode(pincode)
			if err != nil {
				fmt.Printf("failed to fetch for this pin code: %s", err)
			}
			temperatureReport(*response)
		case 2:
			fmt.Print("Exiting Weather App !")
			return
		default:
			fmt.Println("Invalid option. Please enter 1 or 2.")
		}

	}
}

func loadEnvFile() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading .env file")
	}
}

func menu() {
	fmt.Println("1. Get Temperature By Pincode")
	fmt.Println("2. Exit")
}

func temperatureReport(weatherResponse WeatherResponse) {
	fmt.Printf("Temperature Report for %s \n", weatherResponse.Name)
	fmt.Printf("Current Temperature(Celsius): %f \n", weatherResponse.Temperature.Temperature)
	fmt.Printf("Max Temperature(Celsius): %f \n", weatherResponse.Temperature.TemperatureMax)
	fmt.Printf("Min Temperature(Celsius): %f \n", weatherResponse.Temperature.TemperatureMin)
	fmt.Print("\n")
}
