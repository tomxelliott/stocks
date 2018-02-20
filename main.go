package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"time"
)

type QuoteResponse struct {
	Status string
	Name string
	LastPrice float32
	Change float32
	ChangePercent float32
	TimeStamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32
}

var stocks = []string{
	"googl",
	"msft",
	"nflx",
	"tsla",
	"aapl",
	"bbry",
	"hpq",
	"vz",
	"t",
	"tmus",
	"s",
}

func main() {
	numComplete := 0
	for _, symbol := range stocks {
		go func(symbol string) {
			resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	// Ensure computation is carried out before main goroutine terminates
	for numComplete < len(stocks) {
		time.Sleep(10 * time.Millisecond)
	}
}
