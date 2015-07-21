package main

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Notifier func(string)

func parsePage(body io.Reader, temperaturesNotifier Notifier, summaryNotifier Notifier) error {
	z := html.NewTokenizer(body)
	processingTemperatures := false
	processingSummary := false
	for {
		tokenType := z.Next()
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
						summaryNotifier(strings.TrimSpace(attribute.Val))
						break
					}
				}
			}
			break
		case html.TextToken:
			if processingTemperatures {
				processingTemperatures = false
				temperaturesNotifier(strings.TrimSpace(string(z.Text())))
				return nil
			}
			break
		}
	}
}
