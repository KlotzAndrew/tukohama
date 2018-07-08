package main

import (
	"fmt"
	"tukohama/fetcher"
)

var currencyMap = map[int]string{
	1:  "alteration",
	2:  "fusing",
	3:  "alchemy",
	4:  "chaos",
	5:  "gcp",
	6:  "exalted",
	7:  "chrome",
	8:  "jeweller",
	9:  "chance",
	10: "chisel",
	11: "scouring",
	12: "blessed",
	13: "regret",
	14: "regal",
	15: "divine",
	16: "vaal",
	17: "wisdom",
	18: "portal",
	19: "armour...",
	20: "whetst...",
	21: "bauble",
	22: "transmutty",
	23: "augment...",
}

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
	fetcher.GetCurrencyMap()
}
