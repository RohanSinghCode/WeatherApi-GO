package main

import (
	"fmt"
	"net/http"
)

func main() {
	baseUrl := "http://api.openweathermap.org/"
	geoCodeEndpoint := fmt.Sprintf("%d/geo/1.0/direct", baseUrl)
	response, err := http.Get(fmt.Sprintf("%d"))
}
