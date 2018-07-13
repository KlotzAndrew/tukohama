package runner

import (
	"fmt"
	"strconv"
	"sync"
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
	fmt.Println(rates)
	sequences := calc.GetSequences(rates)

	fmt.Println(sequences)
	return sequences
}

func (r Runner) getRateOffers() [][]calc.Rate {
	size := len(r.currencyMap)
	rateOffers := make([][]calc.Rate, size)
	var wg sync.WaitGroup
	results := make(chan rateRes, size*size)

	for i := 0; i < size; i++ {
		rateOffers[i] = make([]calc.Rate, size)

		for j := 0; j < size; j++ {
			wg.Add(1)
			go fetchOffers(i, j, r, results, &wg)
		}
	}
	wg.Wait()
	close(results)

	for c := range results {
		rateOffers[c.x][c.y] = c.rate
	}

	return rateOffers
}

type rateRes struct {
	x, y int
	rate calc.Rate
}

func fetchOffers(i, j int, r Runner, results chan rateRes, wg *sync.WaitGroup) {
	defer wg.Done()

	if i != j {
		offers := r.tradeClient.GetRateOffer(
			strconv.Itoa(r.currencyMap[i].Id),
			strconv.Itoa(r.currencyMap[j].Id),
		)
		avg := avgOffer(offers)
		results <- rateRes{x: i, y: j, rate: calc.NewRate(avg)}
	} else {
		results <- rateRes{x: i, y: j, rate: calc.NewRateNoop()}
	}
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
