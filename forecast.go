package main

import "net/http"

type Forecast struct {
	Temperatures string
	Summary      string
}

func GetForecast() (*Forecast, error) {
	response, err := http.Get("http://www.meteo-bordeaux.com")
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	forecast := &Forecast{}
	temperaturesNotifier := func(temperatures string) {
		forecast.Temperatures = temperatures
	}
	summaryNotifier := func(summary string) {
		forecast.Summary = summary
	}
	err = parsePage(response.Body, temperaturesNotifier, summaryNotifier)
	if err != nil {
		return nil, err
	}
	return forecast, nil
}
