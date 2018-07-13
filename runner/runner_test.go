package runner

import (
	"strconv"
	"testing"
	"tukohama/calc"
	"tukohama/tradeapi"

	"github.com/stretchr/testify/assert"
)

var mockCurrencyMap = []tradeapi.Currency{
	tradeapi.Currency{Id: 1, Name: "alteration"},
	tradeapi.Currency{Id: 2, Name: "fusing"},
	tradeapi.Currency{Id: 3, Name: "alchemy"},
}

type mockClient struct{}

func (m mockClient) GetRateOffer(i, j string) []float64 {
	a, _ := strconv.ParseFloat(i, 64)
	b, _ := strconv.ParseFloat(j, 64)
	return []float64{a, b}
}

func TestGetRateOffers(t *testing.T) {
	runner := Runner{
		tradeClient: mockClient{},
		currencyMap: mockCurrencyMap,
	}
	expected := []calc.Sequence{
		calc.Sequence{[]int{0, 1, 2, 0}, float64(7.5)},
		calc.Sequence{[]int{1, 2, 1}, float64(6.25)},
		calc.Sequence{[]int{2, 1, 2}, float64(6.25)},
	}
	sequences := runner.Run()
	assert.Equal(t, expected, sequences, "runner seqs wrong")
}
