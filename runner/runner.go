package runner

import (
	"fmt"
	"strconv"
	"tukohama/calc"
	"tukohama/tradeapi"
)

func New(tc tradeapi.Client, cm []tradeapi.Currency) Runner {
	return Runner{tradeClient: tc, currencyMap: cm}
}

type Runner struct {
	tradeClient tradeapi.Client
	currencyMap []tradeapi.Currency
}

func (r Runner) Run() []calc.Sequence {
	rates := r.getRateOffers()
	sequences := calc.GetSequences(rates)

	fmt.Println(sequences)
	return sequences
}

func (r Runner) getRateOffers() [][]calc.Rate {
	rateOffers := make([][]calc.Rate, len(r.currencyMap))

	for i := 0; i < len(r.currencyMap); i++ {
		rateOffers[i] = make([]calc.Rate, len(r.currencyMap))

		for j := 0; j < len(r.currencyMap); j++ {
			if i == j {
				rateOffers[i][j] = calc.NewRateNoop()
				continue
			}
			offers := r.tradeClient.GetRateOffer(
				strconv.Itoa(r.currencyMap[i].Id),
				strconv.Itoa(r.currencyMap[j].Id),
			)
			avg := avgOffer(offers)
			rateOffers[i][j] = calc.NewRate(avg)
		}
	}
	return rateOffers
}

func avgOffer(arr []float64) float64 {
	count := len(arr)
	if count > 5 {
		count = 5
	}
	top := arr[0:count]
	var sum float64 = 0
	for _, v := range top {
		sum += v
	}
	return (sum / float64(len(top)))
}
