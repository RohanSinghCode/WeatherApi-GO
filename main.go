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
	fmt.Println("Enter the pincode:")
	var pincode string
	fmt.Scan(&pincode)
	client := NewWeatherClient(os.Getenv("OPENWEATHER_BASE_URL"), os.Getenv("OPENWEAHTER_APP_ID"))
	response, err := client.GetWeatherByPincode(pincode)
	if err != nil {
		fmt.Printf("failed to fetch for this pin code: %s", err)
	}

	fmt.Print(response)
}

func loadEnvFile() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error Loading .env file")
	}
}
