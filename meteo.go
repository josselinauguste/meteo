package main

import "fmt"

func main() {
	forecast, err := GetForecast()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Temperature : %v", forecast.Temperatures)
}
