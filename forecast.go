package main

import "net/http"

type Forecast struct {
	Temperatures string
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
	err = parsePage(response.Body, temperaturesNotifier)
	if err != nil {
		return nil, err
	}
	return forecast, nil
}
