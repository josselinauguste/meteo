package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/ttacon/chalk"
)

var daySpliceNames = [...]string{"Matin", "Après-midi", "Soirée", "Nuit"}

func main() {
	day := getDay()
	dailyForecast, err := GetDailyForecast(day)
	if err != nil {
		panic(err)
	}
	formatDailyForecast(os.Stdout, day, *dailyForecast)
}

func getDay() int {
	if len(os.Args) > 1 {
		day, err := strconv.Atoi(os.Args[1])
		if err == nil {
			return day
		}
	}
	return 0
}

func formatDailyForecast(writer io.Writer, day int, dailyForecast DailyForecast) {
	formatHeader(writer, day)
	for i, forecast := range dailyForecast.Forecasts {
		formatForecast(writer, forecast, i)
	}
	fmt.Fprintln(writer, dailyForecast.Summary)
}

func formatHeader(writer io.Writer, day int) {
	var message string
	switch day {
	case 0:
		message = "Aujourd'hui"
	case 1:
		message = "Demain"
	default:
		message = "Dans " + strconv.Itoa(day) + " jours"
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, " -- ", message, " --")
	fmt.Fprintln(writer)
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
