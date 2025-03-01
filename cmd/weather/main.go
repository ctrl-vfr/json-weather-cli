package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ctrl-vfr/json-weather-cli/internal/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "JSON Weather CLI - Get weather forecasts and current weather for any location.",
	Long: `JSON Weather CLI is a tool to obtain weather forecasts and current weather 
for a geographic position using latitude and longitude.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get weather forecast for a geographic position",
	Long: `Retrieve a weather forecast for the specified latitude and longitude 
for the given number of days.

Example usage:
  weather forecast --latitude 40.7128 --longitude -74.0060 --days 3`,
	Run: func(cmd *cobra.Command, args []string) {
		lat, lon, err := parseCoordinates(cmd)
		if err != nil {
			log.Fatalf("Failed to parse coordinates: %v", err)
		}

		days, err := cmd.Flags().GetInt("days")
		if err != nil {
			log.Fatalf("Failed to parse days flag: %v", err)
		}

		forecast, err := api.GetForecast(lat, lon, days)
		if err != nil {
			log.Fatalf("Failed to get forecast: %v", err)
		}

		w, err := getForecastWeather(forecast)
		if err != nil {
			log.Fatalf("Failed to parse forecast weather: %v", err)
		}

		if err := json.NewEncoder(os.Stdout).Encode(w); err != nil {
			log.Fatalf("Failed to encode forecast weather to JSON: %v", err)
		}
	},
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current weather for a geographic position",
	Long: `Retrieve the current weather for the specified latitude and longitude.

Example usage:
  weather current --latitude 40.7128 --longitude -74.0060`,
	Run: func(cmd *cobra.Command, args []string) {
		lat, lon, err := parseCoordinates(cmd)
		if err != nil {
			log.Fatalf("Failed to parse coordinates: %v", err)
		}

		forecast, err := api.GetForecast(lat, lon, 1)
		if err != nil {
			log.Fatalf("Failed to get current weather: %v", err)
		}

		w, err := getCurrentWeather(forecast)
		if err != nil {
			log.Fatalf("Failed to parse current weather: %v", err)
		}

		if err := json.NewEncoder(os.Stdout).Encode(w); err != nil {
			log.Fatalf("Failed to encode current weather to JSON: %v", err)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("JSON Weather CLI v1.0.0")
	},
}

// Execute sets up and runs the root weather CLI command.
// It registers subcommands and their associated flags.
func Execute() {
	rootCmd.AddCommand(forecastCmd)
	rootCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(versionCmd)

	addCommonFlags(currentCmd)
	addCommonFlags(forecastCmd)
	forecastCmd.Flags().Int("days", 1, "Number of days to forecast")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}
}
