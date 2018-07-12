package tradeapi

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var StaticCurrencyMap = map[int]string{
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

func getNumberName(i int) (string, bool) {
	url := tradeUrlInt(i)
	response, _ := http.Get(url)
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		// log.Fatal(err)
	}

	selectors := doc.Find(".currency-name")
	for i := range selectors.Nodes {
		currencyName := selectors.Eq(i).Text()
		if currencyName != "chaos" {
			currencyName = strings.TrimSpace(currencyName)
			return currencyName, true
		}
	}

	return "", false
}

func tradeUrlInt(from int) string {
	s := strconv.Itoa(from)
	return "http://currency.poe.trade/search?league=Incursion&online=x&want=" + s + "&have=" + "4"
}

type currency struct {
	id   int
	name string
}

type byId []currency

func (s byId) Len() int           { return len(s) }
func (s byId) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byId) Less(i, j int) bool { return s[i].id < s[j].id }

func goGetName(i int, results chan currency, wg *sync.WaitGroup) {
	defer wg.Done()

	name, isFound := getNumberName(i)
	if isFound == true && name != "chaos" {
		results <- currency{i, name}
	}
}

func GetCurrencyMap() []currency {
	var all []currency
	all = append(all, currency{4, "chaos"})
	var wg sync.WaitGroup

	results := make(chan currency, 100)

	for i := 1; i < 24; i++ {
		wg.Add(1)
		go goGetName(i, results, &wg)
	}
	wg.Wait()
	close(results)

	for c := range results {
		all = append(all, c)
	}

	sort.Sort(byId(all))

	for _, v := range all {
		fmt.Printf("%d: \"%s\",\n", v.id, v.name)
	}

	return all
}
