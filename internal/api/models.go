package api

type WeatherResponse struct {
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	CurrentWeatherUnits struct {
		Temperature string `json:"temperature"`
		IsDay       string `json:"is_day"`
		Weathercode string `json:"weathercode"`
	} `json:"current_weather_units"`
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		IsDay       int     `json:"is_day"`
		Weathercode int     `json:"weathercode"`
	} `json:"current_weather"`
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature2M []float64 `json:"temperature_2m"`
		WeatherCode   []int     `json:"weather_code"`
		IsDay         []int     `json:"is_day"`
	} `json:"hourly"`
	Daily struct {
		Time    []string `json:"time"`
		Sunrise []string `json:"sunrise"`
		Sunset  []string `json:"sunset"`
	} `json:"daily"`
}
