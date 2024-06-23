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
	Weather     []Weather `json:"weather"`
	Temperature Main      `json:"main"`
	Units       string    `json:"units"`
	Condition   string    `json:"condition"`
	Name        string    `json:"name"`
}

type Main struct {
	Temperature    float32 `json:"temp"`
	TemperatureMax float32 `json:"temp_max"`
	TemperatureMin float32 `json:"temp_min"`
}

type Weather struct {
	Id   int    `json:"id"`
	Main string `json:"string"`
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

	baseUrl.Path += "/data/2.5/weather"

	params := url.Values{}
	params.Add("q", pincode)
	params.Add("limit", "5")
	params.Add("units", "metric")
	params.Add("appid", wc.config.apiKey)

	// Add query parameters to the base URL
	baseUrl.RawQuery = params.Encode()

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
