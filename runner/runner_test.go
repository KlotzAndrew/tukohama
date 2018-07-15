package runner

import (
	"os"
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
	a, _ := strconv.Atoi(i)
	b, _ := strconv.Atoi(j)
	values := [][][]float64{
		{{1}, {2}, {4}},
		{{0.5}, {1}, {3}},
		{{0.25}, {0.5}, {1}},
	}
	return values[a-1][b-1]
}

type mockNilClient struct{}

func (m mockNilClient) GetRateOffer(i, j string) []float64 {
	return []float64{}
}

func TestGetRateOffers(t *testing.T) {
	runner := Runner{
		tradeClient: mockClient{},
		currencyMap: mockCurrencyMap,
	}
	expected := []calc.Sequence{
		calc.Sequence{[]int{0, 1, 2, 0}, float64(1.5)},
		calc.Sequence{[]int{1, 2, 1}, float64(1.5)},
		calc.Sequence{[]int{2, 1, 2}, float64(1.5)},
	}
	sequences := runner.Run()
	assert.Equal(t, expected, sequences, "runner seqs wrong")
}

func TestGetRateOffersNoOffer(t *testing.T) {
	runner := Runner{
		tradeClient: mockNilClient{},
		currencyMap: mockCurrencyMap,
	}
	var expected []calc.Sequence

	sequences := runner.Run()
	assert.Equal(t, expected, sequences, "runner seqs wrong")
}

func TestSaveRates(t *testing.T) {
	runner := Runner{
		tradeClient: mockClient{},
		currencyMap: mockCurrencyMap,
	}

	filePath := runner.RatesToCsv("/tmp")
	contents := fileContents(t, filePath)

	pwd, err := os.Getwd()
	assert.Equal(t, nil, err, err)
	testContents := fileContents(t, pwd+"/rate_data_test.csv")

	assert.Equal(t, string(testContents), string(contents), "csv contents bad")
}

func fileContents(t *testing.T, filePath string) []byte {
	file, err := os.Open(filePath)
	assert.Equal(t, nil, err, err)

	filestat, err := file.Stat()
	assert.Equal(t, nil, err, err)
	size := filestat.Size()

	contents := make([]byte, size)
	_, err = file.Read(contents)
	assert.Equal(t, nil, err, err)

	return contents
}
