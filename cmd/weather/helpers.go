package weather

import (
	"time"
)

// Return the emoji corresponding to the weather code
func getWeatherEmoji(code int) string {
	for emoji, codes := range weatherEmojiMap {
		for _, c := range codes {
			if c == code {
				return emoji
			}
		}
	}
	return "❓"
}

// Return the emoji corresponding to the time of day and the sunrise/sunset times
func getCycleEmoji(date time.Time, isDay bool, sunrise, sunset time.Time) string {
	if !isDay {
		if date.After(sunrise.Add(-time.Hour)) && date.Before(sunset) {
			return "🌆"
		}
		if date.After(sunrise) && date.Before(sunset.Add(time.Hour)) {
			return "🌆"
		}
		return "🌃"
	}
	if date.After(sunset.Add(-time.Hour)) {
		return "🌇"
	}
	if date.Before(sunrise.Add(time.Hour)) {
		return "🌇"
	}

	return "🏙️"
}

// floor rounds a float to the nearest integer
func floor(f float64) int {
	return int(f + 0.5)
}
