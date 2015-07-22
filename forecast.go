package main

import "net/http"

type Forecast struct {
	Temperatures string
	Summary      string
}

type DailyForecast struct {
	Forecasts []Forecast
	Summary   string
}

func NewDailyForecast(forecasts []Forecast, summary string) *DailyForecast {
	return &DailyForecast{forecasts, summary}
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
	return NewDailyForecast(parser.Forecasts, parser.GlobalSummary), nil
}
