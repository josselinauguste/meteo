package main

import (
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"github.com/ttacon/chalk"
)

var daySpliceNames = [...]string{"Matin", "Après-midi", "Soirée", "Nuit"}

func main() {
	dailyForecast, err := GetDailyForecast()
	if err != nil {
		panic(err)
	}
	formatDailyForecast(os.Stdout, *dailyForecast)
}

func formatDailyForecast(writer io.Writer, dailyForecast DailyForecast) {
	for i, forecast := range dailyForecast.Forecasts {
		formatForecast(writer, forecast, i)
	}
	fmt.Fprintln(writer, dailyForecast.Summary)
}

func formatForecast(writer io.Writer, forecast Forecast, i int) {
	longestSpliceNameLength := getLongestSpliceNameLength()
	headerTabulation := generateTabulation(longestSpliceNameLength - utf8.RuneCountInString(daySpliceNames[i]))
	fmt.Fprintln(writer, chalk.Green, daySpliceNames[i], headerTabulation, forecast.Temperatures)
	fmt.Fprintln(writer, chalk.White, forecast.Summary)
	fmt.Fprintln(writer)
}

func getLongestSpliceNameLength() int {
	var longestSpliceNameLength int
	for _, name := range daySpliceNames {
		length := utf8.RuneCountInString(name)
		if length > longestSpliceNameLength {
			longestSpliceNameLength = length
		}
	}
	return longestSpliceNameLength
}

func generateTabulation(length int) string {
	var tabulation string
	for k := 0; k < length; k++ {
		tabulation = tabulation + " "
	}
	return tabulation
}
