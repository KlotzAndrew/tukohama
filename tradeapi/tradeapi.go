package tradeapi

import (
	"net/http"
	"strconv"

	"tukohama/calc"

	"golang.org/x/net/html"
)

type exchangeDiv struct {
	sellCurrency string
	buyCurrency  string
	sellValue    float64
	buyValue     float64
}

func GetRateOffers(currencies map[int]string) [][]calc.Rate {
	rateOffers := make([][]calc.Rate, len(currencies))

	for i := 0; i < len(currencies); i++ {
		rateOffers[i] = make([]calc.Rate, len(currencies))

		for j := 0; j < len(currencies); j++ {
			if i == j {
				rateOffers[i][j] = calc.NewRateNoop()
				continue
			}
			// offers := GetRateOffer(i, j)
			avg := float64(10) // avgOffer(offers)
			rateOffers[i][j] = calc.NewRate(avg)
		}
	}
	return rateOffers
}

func avgOffer(arr []float64) float64 {
	top := arr[0:5]
	var sum float64 = 0
	for _, v := range top {
		sum += v
	}
	return (sum / float64(len(top)))
}

func GetRateOffer(fromI int, toI int) []float64 {
	from := strconv.Itoa(fromI)
	to := strconv.Itoa(toI)
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
