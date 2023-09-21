package weatherSubscription

import (
	"fmt"
	"testing"
)

func TestForecast(t *testing.T) {
	var weatherApi WeatherApi
	weatherApi = &MockWeatherApi{}
	forecast, err := weatherApi.GetForecast(123, 456)
	if err == nil {
		fmt.Println(forecast)
	}
}

func TestCreateUser(t *testing.T) {

}
