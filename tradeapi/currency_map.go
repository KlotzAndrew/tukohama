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

var StaticCurrencyMap = []Currency{
	Currency{Id: 1, Name: "alteration"},
	Currency{Id: 2, Name: "fusing"},
	Currency{Id: 3, Name: "alchemy"},
	Currency{Id: 4, Name: "chaos"},
	Currency{Id: 5, Name: "gcp"},
	Currency{Id: 6, Name: "exalted"},
	Currency{Id: 7, Name: "chrome"},
	Currency{Id: 8, Name: "jeweller"},
	Currency{Id: 9, Name: "chance"},
	Currency{Id: 10, Name: "chisel"},
	Currency{Id: 11, Name: "scouring"},
	Currency{Id: 12, Name: "blessed"},
	Currency{Id: 13, Name: "regret"},
	Currency{Id: 14, Name: "regal"},
	Currency{Id: 15, Name: "divine"},
	Currency{Id: 16, Name: "vaal"},
	Currency{Id: 17, Name: "wisdom"},
	Currency{Id: 18, Name: "portal"},
	Currency{Id: 19, Name: "armour..."},
	Currency{Id: 20, Name: "whetst..."},
	Currency{Id: 21, Name: "bauble"},
	Currency{Id: 22, Name: "transmutty"},
	Currency{Id: 23, Name: "augment..."},
}

func SeqToNames(indexes []int) []string {
	result := []string{}
	for _, index := range indexes {
		name := StaticCurrencyMap[index].Name
		if name == "" {
			panic("name not found!")
		}
		result = append(result, name)
	}
	return result
}

func getNumberName(i int) (string, bool) {
	url := tradeUrlInt(i)
	response, _ := http.Get(url)
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
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

type Currency struct {
	Id   int
	Name string
}

type byId []Currency

func (s byId) Len() int           { return len(s) }
func (s byId) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byId) Less(i, j int) bool { return s[i].Id < s[j].Id }

func goGetName(i int, results chan Currency, wg *sync.WaitGroup) {
	defer wg.Done()

	name, isFound := getNumberName(i)
	if isFound == true && name != "chaos" {
		results <- Currency{Id: i, Name: name}
	}
}

func GetCurrencyMap() []Currency {
	var all []Currency
	all = append(all, Currency{Id: 4, Name: "chaos"})
	var wg sync.WaitGroup

	results := make(chan Currency, 100)

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
		fmt.Printf("%d: \"%s\",\n", v.Id, v.Name)
	}

	return all
}
