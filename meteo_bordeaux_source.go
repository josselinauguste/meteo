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

func GetDailyForecast() (*DailyForecast, error) {
	source := getSource()
	parser, err := parseSource(source)
	if err != nil {
		return nil, err
	}
	return NewDailyForecast(parser.Forecasts, parser.GlobalSummary), nil
}

func getSource() string {
	return "http://www.meteo-bordeaux.com"
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
