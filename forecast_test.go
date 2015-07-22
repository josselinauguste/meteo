package main

import (
	"regexp"
	"testing"
)

func TestGetForecast(t *testing.T) {
	dailyForecast, err := GetDailyForecast()

	if err != nil {
		t.Errorf("error should be nil, not %v", err)
	}
	if dailyForecast == nil {
		t.Error("forecast should not be nil")
	} else {
		if len(dailyForecast.Forecasts) != 4 {
			t.Errorf("a day counts 4 forecasts, not %v", len(dailyForecast.Forecasts))
		} else {
			forecast := dailyForecast.Forecasts[0]
			if forecast.Temperatures == "" {
				t.Error("temperatures should be set")
			}
			if matched, _ := regexp.MatchString("^\\d{1,2}/\\d{1,2} °C$", forecast.Temperatures); !matched {
				t.Errorf("temperatures has an invalid format: %v", forecast.Temperatures)
			}
			if forecast.Summary == "" {
				t.Error("summary should be set")
			}
		}
	}
}
