package weather

import (
	"time"

	"github.com/ctrl-vfr/json-weather-cli/internal/api"
)

// Return the Weather data for the current time
func getCurrentWeather(forecast *api.WeatherResponse) (Response, error) {
	date := time.Now().UTC()
	sunrise, sunset, err := getDailySunRiseSunset(time.Now().UTC(), forecast)
	if err != nil {
		return Response{}, err
	}
	isDay := forecast.CurrentWeather.IsDay == 1

	currentWeather := Response{
		Date:        date.Format("2006-01-02T15:04"),
		Temperature: floor(forecast.CurrentWeather.Temperature),
		Weather:     getWeatherEmoji(forecast.CurrentWeather.Weathercode),
		Cycle:       getCycleEmoji(date, isDay, sunrise, sunset),
	}

	return currentWeather, nil
}

// Return the Weather data for the forecasted time
func getForecastWeather(forecast *api.WeatherResponse) ([]Response, error) {
	var weathers []Response
	for i, t := range forecast.Hourly.Time {
		date, err := time.Parse("2006-01-02T15:04", t)
		if err != nil {
			return nil, err
		}

		isDay := forecast.Hourly.IsDay[i] == 1
		sunrise, sunset, err := getDailySunRiseSunset(date, forecast)
		if err != nil {
			return nil, err
		}

		weathers = append(weathers, Response{
			Date:        date.Format("2006-01-02T15:04"),
			Temperature: floor(forecast.Hourly.Temperature2M[i]),
			Weather:     getWeatherEmoji(forecast.Hourly.WeatherCode[i]),
			Cycle:       getCycleEmoji(date, isDay, sunrise, sunset),
		})
	}
	return weathers, nil
}

// Return the sunrise and sunset times for a given date
func getDailySunRiseSunset(date time.Time, forecast *api.WeatherResponse) (time.Time, time.Time, error) {
	for i, t := range forecast.Daily.Time {
		d, err := time.Parse("2006-01-02", t)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		if d.Format("2006-01-02") == date.Format("2006-01-02") {
			sunrise, err := time.Parse("2006-01-02T15:04", forecast.Daily.Sunrise[i])
			if err != nil {
				return time.Time{}, time.Time{}, err
			}
			sunset, err := time.Parse("2006-01-02T15:04", forecast.Daily.Sunset[i])
			if err != nil {
				return time.Time{}, time.Time{}, err
			}
			return sunrise, sunset, nil
		}
	}
	return time.Time{}, time.Time{}, nil
}
