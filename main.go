package main

import (
	"fmt"
	// "tukohama/fetcher"
	// "tukohama/calc"
)

func topAvg(arr []float64) float64 {
	top := arr[0:5]
	var sum float64 = 0
	for _, v := range top {
		sum += v
	}
	return (sum / float64(len(top)))
}

func main() {
	// offers := fetcher.GetRateOffers("3", "4")
	offers := []float64{2.02, 2.02, 2, 2, 2, 2}
	fmt.Println(topAvg(offers))
	// fetcher.GetCurrencyMap()
	// calc.rates()
}
