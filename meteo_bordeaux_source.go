// var date = $("#date_accueil").text();
// 		$.ajax(
// 		{
// 			url: '/accueil/jour_plus/' + date,
// 			async: false,
// 			type: "POST",
// 			data: "",
// 			dataType: "",
// 			success: function(data)
// 			{
// 				$('#prevision').html(data);
// 			}
// 		});

package main

import "net/http"

func GetDailyForecast(day int) (*DailyForecast, error) {
	source, err := getSource(day)
	if err != nil {
		return nil, err
	}
	parser, err := parseSource(source)
	if err != nil {
		return nil, err
	}
	return NewDailyForecast(parser.Forecasts, parser.GlobalSummary), nil
}

func getSource(day int) (string, error) {
	currentSource := "http://www.meteo-bordeaux.com"
	for i := 0; i < day; i++ {
		parser, err := parseSource(currentSource)
		if err != nil {
			return "", err
		}
		currentSource = "http://www.meteo-bordeaux.com/accueil/jour_plus/" + parser.today
	}
	return currentSource, nil
}

func parseSource(source string) (*ForecastsParser, error) {
	response, err := http.Get(source)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	parser := NewForecastsParser()
	err = parser.ParsePage(response.Body)
	if err != nil {
		return nil, err
	}
	return parser, nil
}
