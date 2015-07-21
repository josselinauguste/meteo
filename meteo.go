package main

import (
	"fmt"
	"strings"
)

func main() {
	forecast, err := GetForecast()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Matin : %v, %v", forecast.Temperatures, strings.ToLower(forecast.Summary))
}
