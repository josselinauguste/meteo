package main

import "testing"

func TestNewDailyForecast(t *testing.T) {
	forecasts := make([]Forecast, 0)
	summary := "lala"

	dailyForecast := NewDailyForecast(forecasts, summary)

	if dailyForecast.Forecasts == nil {
		t.Error("expected to be equal")
	}

	if dailyForecast.Summary != summary {
		t.Error("expected to be equal")
	}
}
