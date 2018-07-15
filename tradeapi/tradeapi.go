package tradeapi

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/html"
)

type exchangeDiv struct {
	sellCurrency string
	buyCurrency  string
	sellValue    float64
	buyValue     float64
}

type Client interface {
	GetRateOffer(i, j string) []float64
}
type ConcreteClient struct{}

func (c ConcreteClient) GetRateOffer(from, to string) []float64 {
	url := tradeUrl(from, to)
	response := get(url)
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

func get(url string) *http.Response {
	retries := 0

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	for response.StatusCode != 200 {
		if retries > 5 {
			panic("http request error!")
		}

		jitter := 1000 + rand.Intn(20)*50*(retries+1)
		httpBackoff := time.Duration(jitter) * time.Millisecond
		fmt.Printf("http backoff, retries: %d; time:%s\n", retries, httpBackoff.String())
		time.Sleep(httpBackoff)

		response, err = http.Get(url)
		if err != nil {
			panic(err)
		}

		retries += 1
	}

	return response
}

func tradeUrl(from, to string) string {
	return "http://currency.poe.trade/search?league=Incursion&online=x&want=" + to + "&have=" + from
}

func findOfferDiv(t html.Token, from, to string) (float64, bool) {
	// <div class="displayoffer " data-username="Thyronis" data-sellcurrency="3" data-sellvalue="20.0" data-buycurrency="4" data-buyvalue="10.0" data-ign="Gottamakethatmoney" >
	if t.Data == "div" {
		var div exchangeDiv
		for _, a := range t.Attr {
			if a.Key == "data-sellcurrency" && a.Val == to {
				div.sellCurrency = a.Val
			}
			if a.Key == "data-buycurrency" && a.Val == from {
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
		if div.sellCurrency == to && div.buyCurrency == from {
			return (div.sellValue / div.buyValue), true
		}
	}
	return 0, false
}
