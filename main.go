package main

import (
	"flag"
	"fmt"
	"os"

	"tukohama/runner"
	"tukohama/tradeapi"
)

func main() {
	cmd := flag.String("cmd", "", "command to run")
	flag.Parse()

	switch *cmd {
	case "full":
		fullRun()
	case "save-rates":
		saveRates()
	default:
		fmt.Println("command required!")
		os.Exit(1)
	}
}

func fullRun() {
	runner := runner.New(
		tradeapi.ConcreteClient{},
		tradeapi.StaticCurrencyMap,
	)
	runner.Run()
}

func saveRates() {
	// runner := runner.New(
	// 	tradeapi.ConcreteClient{},
	// 	tradeapi.StaticCurrencyMap,
	// )
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// runner.RatesToCsv(pwd + "/data")
}
