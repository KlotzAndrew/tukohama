package main

import (
	"tukohama/runner"
	"tukohama/tradeapi"
)

func main() {
	runner := runner.New(
		tradeapi.ConcreteClient{},
		tradeapi.StaticCurrencyMap,
	)
	runner.Run()
}
