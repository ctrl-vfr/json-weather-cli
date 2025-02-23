# JSON Weather CLI

JSON Weather CLI is a small CLI tool that retrieves current weather conditions and forecasts for a specified location using the Open-Meteo API.

## Installation

1. **Clone the repository:**  
   ```sh
   git clone https://github.com/ctrl-vfr/json-weather-cli.git
   cd json-weather-cli
   ```

2. **Install dependencies:**  
   ```sh
   go mod tidy
   ```

3. **Build the project:**  
   ```sh
   go build -o json-weather-cli
   ```

Alternatively, you can install it directly using `go install`:
```sh
go install github.com/ctrl-vfr/json-weather-cli@latest
```
Make sure your `$GOPATH/bin` is in your system's `PATH` to run `weather` from anywhere.

## Usage

### Get the current weather
```sh
json-weather-cli current --latitude 48.8566 --longitude 2.3522
```
Example of prettified response:
```json
{
    "date": "2025-02-23T14:00",
    "temperature": 12,
    "weather": "☀️",
    "cycle": "🏙️"
}
```

### Get the weather forecast
```sh
json-weather-cli forecast --latitude 48.8566 --longitude 2.3522 --days 3 | jq
```
Example of prettified response:
```json
[
  {
    "date": "2025-02-23T00:00",
    "temperature": 7,
    "weather": "☀️",
    "cycle": "🌃"
  },
  {
    "date": "2025-02-23T01:00",
    "temperature": 7,
    "weather": "🌫️",
    "cycle": "🌃"
  },
  {
    "date": "2025-02-23T02:00",
    "temperature": 6,
    "weather": "🌫️",
    "cycle": "🌃"
  },
  ...
]
```

## License

This project is licensed under the MIT License.
