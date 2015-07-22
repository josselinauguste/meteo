package main

import (
	"fmt"
	"strings"
)

var daySpliceNames = [...]string{"Matin", "Après-midi", "Soirée", "Nuit"}

func main() {
	dailyForecast, err := GetDailyForecast()
	if err != nil {
		panic(err)
	}
	for i, forecast := range dailyForecast.Forecasts {
		fmt.Printf("%v : %v, %v\n", daySpliceNames[i], forecast.Temperatures, strings.ToLower(forecast.Summary))
	}
}
