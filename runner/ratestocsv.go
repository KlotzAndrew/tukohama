package runner

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"tukohama/calc"
	"tukohama/tradeapi"
)

func (r Runner) ratesToCsv(rates [][]calc.Rate, dir string) string {
	file := newCsvFile(dir)
	defer file.Close()
	w := csv.NewWriter(file)

	writeCsvRowHeader(w, r.currencyMap)
	writeCsvRow(w, r.currencyMap, rates)

	w.Flush()
	if err := w.Error(); err != nil {
		panic(err)
	}
	fmt.Println(file.Name())

	return file.Name()
}

func writeCsvRowHeader(w *csv.Writer, curs []tradeapi.Currency) {
	row := []string{"rateMatrix"}
	for _, cur := range curs {
		row = append(row, cur.Name)
	}
	if err := w.Write(row); err != nil {
		panic(err)
	}
}

func writeCsvRow(w *csv.Writer, curs []tradeapi.Currency, rates [][]calc.Rate) {
	for i, rateRow := range rates {
		row := []string{curs[i].Name}
		for _, r := range rateRow {
			s := strconv.FormatFloat(r.Value, 'f', -1, 64)
			row = append(row, s)
		}
		if err := w.Write(row); err != nil {
			panic(err)
		}
	}
}

func newCsvFile(dir string) *os.File {
	t := time.Now()
	fileName := dir + "/data/rates-" + t.Format("20060102150405") + ".csv"

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return file
}
