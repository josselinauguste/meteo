package main

import "net/http"

type Forecast struct {
	Temperatures string
	Summary      string
}

type DailyForecast struct {
	Forecasts []Forecast
}

func NewDailyForecast(forecasts []Forecast) *DailyForecast {
	return &DailyForecast{forecasts}
}

func GetDailyForecast() (*DailyForecast, error) {
	response, err := http.Get("http://www.meteo-bordeaux.com")
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	parser := NewForecastsParser()
	err = parser.ParsePage(response.Body)
	if err != nil {
		return nil, err
	}
	return NewDailyForecast(parser.Forecasts), nil
}
