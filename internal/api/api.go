package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var ENDPOINT = "https://api.open-meteo.com/v1/"

func GetForecast(lat, lon float64, days int) (*WeatherResponse, error) {
	url := fmt.Sprintf("%s/forecast?latitude=%f&longitude=%f&hourly=temperature_2m,weather_code,is_day&current_weather=true&daily=sunrise,sunset&forecast_days=%d", ENDPOINT, lat, lon, days)
	var response WeatherResponse
	err := get(url, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func get(url string, target interface{}) error {
	http.DefaultClient.Timeout = time.Second
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
