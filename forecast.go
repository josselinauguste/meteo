package main

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
