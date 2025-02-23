package weather

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ctrl-vfr/json-weather-cli/internal/api"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "JSON Weather CLI is a tool to obtain the weather forecast for a geographic position",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var forcastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Obtain the weather forecast for a geographic position",
	Run: func(cmd *cobra.Command, args []string) {
		lat, _ := cmd.Flags().GetFloat64("latitude")
		lon, _ := cmd.Flags().GetFloat64("longitude")
		days, _ := cmd.Flags().GetInt("days")

		forecast, err := api.GetForecast(lat, lon, days)
		if err != nil {
			log.Fatal(err)
		}

		w, err := getForecastWeather(forecast)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(os.Stdout).Encode(w)
	},
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Obtain the current weather for a geographic position",
	Run: func(cmd *cobra.Command, args []string) {
		lat, _ := cmd.Flags().GetFloat64("latitude")
		lon, _ := cmd.Flags().GetFloat64("longitude")

		forecast, err := api.GetForecast(lat, lon, 1)
		if err != nil {
			log.Fatal(err)
		}

		w, err := getCurrentWeather(forecast)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(os.Stdout).Encode(w)
	},
}

func Execute() {
	rootCmd.AddCommand(forcastCmd)
	rootCmd.AddCommand(currentCmd)
	currentCmd.Flags().Float64("latitude", 0, "Position's latitude")
	currentCmd.Flags().Float64("longitude", 0, "Position's longitude")

	forcastCmd.Flags().Float64("latitude", 0, "Position's latitude")
	forcastCmd.Flags().Float64("longitude", 0, "Position's longitude")
	forcastCmd.Flags().Int("days", 1, "Number of days to forecast")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
