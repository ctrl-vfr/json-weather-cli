package weather

// weatherEmojiMap maps weather codes to emojis
var weatherEmojiMap = map[string][]int{
	"☀️": {0},                      // Clear sky
	"⛅":  {1, 2, 3},                // Mainly clear, partly cloudy, overcast
	"🌫️": {45, 48},                 // Fog and depositing rime fog
	"🌧️": {51, 53, 55, 61, 63, 65}, // Drizzle light, moderate, r ain slight, moderate, heavy
	"🌨️": {56, 57, 66, 67, 77},     // Freezing drizzle light and dense, Freezing rain light, heavy
	"❄️": {71, 73, 75, 85, 86},     // Snow fall & snow showers slight, moderate, heavy, snow grains
	"🌦️": {80, 81, 82},             // Rain showers slight, moderate, violent
	"⛈️": {95, 96, 99},             // Thunderstorm slight/moderate, with hail
}
