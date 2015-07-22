package main

import (
	"regexp"
	"testing"
)

func TestGetSource(t *testing.T) {
	source, _ := getSource(0)

	if source != "http://www.meteo-bordeaux.com" {
		t.Error("Source is not as expected")
	}
}

func TestGetSourceForNextDay(t *testing.T) {
	source, err := getSource(1)

	if err != nil {
		t.Errorf("error should be nil, not %v", err)
	}
	matched, err := regexp.MatchString("http://www.meteo-bordeaux.com/accueil/jour_plus/\\d+", source)
	if !matched || err != nil {
		t.Errorf("Source is not as expected: %v", source)
	}
}

func TestGetForecast(t *testing.T) {
	dailyForecast, err := GetDailyForecast(0)

	if err != nil {
		t.Errorf("error should be nil, not %v", err)
	}
	assertDailyForecast(t, dailyForecast, 4)
}

func TestGetForecastForNextDay(t *testing.T) {
	dailyForecast, err := GetDailyForecast(1)

	if err != nil {
		t.Errorf("error should be nil, not %v", err)
	}
	assertDailyForecast(t, dailyForecast, 4)
}

func TestGetForecastForIn2Days(t *testing.T) {
	dailyForecast, err := GetDailyForecast(2)

	if err != nil {
		t.Errorf("error should be nil, not %v", err)
	}
	assertDailyForecast(t, dailyForecast, 2)
}

func assertDailyForecast(t *testing.T, dailyForecast *DailyForecast, forecastsExpected int) {
	if dailyForecast == nil {
		t.Error("forecast should not be nil")
	} else {
		if dailyForecast.Summary == "" {
			t.Error("daily summary should be set")
		}
		if len(dailyForecast.Forecasts) != forecastsExpected {
			t.Errorf("a day counts %v forecasts, not %v", forecastsExpected, len(dailyForecast.Forecasts))
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
