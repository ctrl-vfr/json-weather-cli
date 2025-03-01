package weather

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
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

// parseCoordinates extracts latitude and longitude flags from the command.
func parseCoordinates(cmd *cobra.Command) (float64, float64, error) {
	lat, err := cmd.Flags().GetFloat64("latitude")
	if err != nil {
		return 0, 0, fmt.Errorf("invalid latitude: %w", err)
	}
	lon, err := cmd.Flags().GetFloat64("longitude")
	if err != nil {
		return 0, 0, fmt.Errorf("invalid longitude: %w", err)
	}
	return lat, lon, nil
}

// addCommonFlags adds latitude and longitude flags to a command.
func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().Float64("latitude", 0, "Position's latitude (required)")
	cmd.Flags().Float64("longitude", 0, "Position's longitude (required)")
}
