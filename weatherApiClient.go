package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type ApiConfg struct {
	apiKey  string
	baseUrl string
}

type WeatherClient struct {
	config ApiConfg
	client *http.Client
}

type WeatherResponse struct {
	Location    string `json:"location"`
	Temperature string `json:"temperature"`
	Unit        string `json:"unit"`
	Condition   string `json:"condition"`
}

type GeoCodingResponse struct {
	Zip     string `json:"zip"`
	Name    string `json:"name"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Country string `json:"country"`
}

func NewWeatherClient(baseUrl, apiKey string) *WeatherClient {
	return &WeatherClient{
		config: ApiConfg{
			apiKey:  apiKey,
			baseUrl: baseUrl,
		},
		client: &http.Client{},
	}
}

func (wc *WeatherClient) GetWeatherByPincode(pincode string) (*WeatherResponse, error) {
	baseUrl, err := url.Parse(wc.config.baseUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing url : %s", err)
	}
	geoCoding, err := wc.GetLatLongByZipCode(pincode)
	if err != nil {
		return nil, fmt.Errorf("failed fetching lat long by pincode: %s", err)
	}

	baseUrl.Path += "/geo/1.0/reverse"

	params := url.Values{}
	params.Add("lat", geoCoding.Lat)
	params.Add("lon", geoCoding.Lon)
	params.Add("limit", "5")
	params.Add("appId", wc.config.apiKey)

	response, err := wc.client.Get(baseUrl.String())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather by lat long: %w", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json: %s", err)
	}
	return &weatherResponse, nil
}

func (wc *WeatherClient) GetLatLongByZipCode(pincode string) (*GeoCodingResponse, error) {
	baseUrl, err := url.Parse(wc.config.baseUrl)
	if err != nil {
		return nil, fmt.Errorf("error parsing url : %s", err)
	}
	baseUrl.Path += "/geo/1.0/zip"

	// Add query parameters.
	params := url.Values{}
	params.Add("zip", pincode)
	params.Add("limit", "5")
	params.Add("appId", wc.config.apiKey)
	baseUrl.RawQuery = params.Encode()
	response, err := wc.client.Get(baseUrl.String())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch by this error code: %w", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var getCodingRespond GeoCodingResponse
	err = json.Unmarshal(body, &getCodingRespond)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json: %s", err)
	}

	return &getCodingRespond, nil
}
