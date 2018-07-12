package runner

import (
	"fmt"
	"tukohama/calc"
	"tukohama/tradeapi"
)

func Run() {
	currencyMap := tradeapi.StaticCurrencyMap
	rates := tradeapi.GetRateOffers(currencyMap)
	sequences := calc.GetSequences(rates)

	fmt.Println(sequences)
}
