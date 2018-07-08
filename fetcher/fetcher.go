package fetcher

import (
	"golang.org/x/net/html"
	"net/http"
	"strconv"
)

type exchangeDiv struct {
	sellCurrency string
	buyCurrency  string
	sellValue    float64
	buyValue     float64
}

func GetRateOffers(from string, to string) []float64 {
	url := tradeUrl(from, to)
	response, _ := http.Get(url)
	defer response.Body.Close()

	tokens := html.NewTokenizer(response.Body)

	var offers []float64
	for {
		tt := tokens.Next()

		switch {
		case tt == html.ErrorToken:
			return offers
		case tt == html.StartTagToken:
			t := tokens.Token()
			offer, isOffer := findOfferDiv(t, from, to)
			if isOffer == true {
				offers = append(offers, offer)
			}
		}
	}
}

func tradeUrl(from string, to string) string {
	return "http://currency.poe.trade/search?league=Incursion&online=x&want=" + from + "&have=" + to
}

func findOfferDiv(t html.Token, from string, to string) (float64, bool) {
	// <div class="displayoffer " data-username="Thyronis" data-sellcurrency="3" data-sellvalue="20.0" data-buycurrency="4" data-buyvalue="10.0" data-ign="Gottamakethatmoney" >
	if t.Data == "div" {
		var div exchangeDiv
		for _, a := range t.Attr {
			if a.Key == "data-sellcurrency" && a.Val == from {
				div.sellCurrency = a.Val
			}
			if a.Key == "data-buycurrency" && a.Val == to {
				div.buyCurrency = a.Val
			}
			if a.Key == "data-sellvalue" {
				f, _ := strconv.ParseFloat(a.Val, 64)
				div.sellValue = f
			}
			if a.Key == "data-buyvalue" {
				f, _ := strconv.ParseFloat(a.Val, 64)
				div.buyValue = f
			}
		}
		if div.sellCurrency == from && div.buyCurrency == to {
			return (div.sellValue / div.buyValue), true
		}
	}
	return 0, false
}
