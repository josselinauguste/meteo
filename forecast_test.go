package main

import (
	"regexp"
	"testing"
)

func TestGetForecast(t *testing.T) {
	forecast, err := GetForecast()

	if err != nil {
		t.Errorf("error should be nil, not %v", err)
	}
	if forecast == nil {
		t.Error("forecast should not be nil")
	} else {
		if forecast.Temperatures == "" {
			t.Error("temperatures should be set")
		}
		if matched, _ := regexp.MatchString("^\\d{1,2}/\\d{1,2} °C$", forecast.Temperatures); !matched {
			t.Errorf("temperatures has an invalid format: %v", forecast.Temperatures)
		}
	}
}
