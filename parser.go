package main

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type ForecastsParser struct {
	buildedForecast *Forecast
	Forecasts       []Forecast
	GlobalSummary   string
}

func NewForecastsParser() *ForecastsParser {
	forecastsParser := &ForecastsParser{}
	forecastsParser.Forecasts = make([]Forecast, 0)
	return forecastsParser
}

func (parser *ForecastsParser) ParsePage(body io.Reader) error {
	z := html.NewTokenizer(body)
	defer parser.parsingFinished()
	processingTemperatures := false
	processingSummary := false
	globalSummaryDeepness := 0
	for {
		tokenType := z.Next()
		if globalSummaryDeepness > 0 {
			globalSummaryDeepness++
		}
		switch tokenType {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				return nil
			}
			return z.Err()
		case html.StartTagToken:
			token := z.Token()
			if token.Data == "div" {
				for _, attribute := range token.Attr {
					if attribute.Key == "class" && attribute.Val == "ac_picto_ensemble" {
						parser.foundForecast()
						break
					}
					if attribute.Key == "class" && attribute.Val == "ac_com" {
						globalSummaryDeepness = 1
						break
					}
					if attribute.Key == "class" && attribute.Val == "ac_temp" {
						processingTemperatures = true
						break
					}
					if attribute.Key == "class" && attribute.Val == "ac_picto" {
						processingSummary = true
						break
					}
				}
			}
			break
		case html.SelfClosingTagToken:
			if processingSummary {
				token := z.Token()
				for _, attribute := range token.Attr {
					if attribute.Key == "title" {
						processingSummary = false
						parser.summaryFound(attribute.Val)
						break
					}
				}
			}
			break
		case html.TextToken:
			if processingTemperatures {
				processingTemperatures = false
				parser.temperaturesFound(string(z.Text()))
			} else if globalSummaryDeepness >= 4 {
				globalSummaryDeepness = 0
				parser.globalSummaryFound(string(z.Text()))
			}
			break
		}
	}
}

func (parser *ForecastsParser) foundForecast() {
	parser.addForecast(parser.buildedForecast)
	parser.buildedForecast = &Forecast{}
}

func (parser *ForecastsParser) parsingFinished() {
	parser.addForecast(parser.buildedForecast)
}

func (parser *ForecastsParser) addForecast(forecast *Forecast) {
	if forecast != nil {
		parser.Forecasts = append(parser.Forecasts, *forecast)
	}
}

func (parser *ForecastsParser) temperaturesFound(value string) {
	parser.buildedForecast.Temperatures = strings.TrimSpace(value)
}

func (parser *ForecastsParser) summaryFound(value string) {
	parser.buildedForecast.Summary = strings.TrimSpace(value)
}

func (parser *ForecastsParser) globalSummaryFound(value string) {
	parser.GlobalSummary = strings.TrimSpace(value)
}
