package weatherSubscription

import "fmt"

type MockWeatherApi struct {
}

func (weatherApi *MockWeatherApi) GetForecast(lat, lon int) (string, error) {
	forecast := fmt.Sprintf("Forecast is very well\nLat: [%d] Lon: [%d]\n", lat, lon)
	return forecast, nil
}
