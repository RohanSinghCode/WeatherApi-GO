package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	url := fmt.Sprintf("%s/geo/1.0/direct?q=%s&limit=%s&appId=%s", wc.config.baseUrl, pincode, wc.config.apiKey)

	response, err := wc.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch by this error code: %w", err)
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
