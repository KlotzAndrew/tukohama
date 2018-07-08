package main

import (
	"fmt"

	"tukohama/fetcher"
)

func main() {
	offers := fetcher.GetRateOffers("3", "4")
	for _, o := range offers {
		fmt.Println(o)
	}
}
