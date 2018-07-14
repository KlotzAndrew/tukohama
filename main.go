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
