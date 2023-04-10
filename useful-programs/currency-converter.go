package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rates struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func convertCurrency(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	url := fmt.Sprintf("https://api.exchangeratesapi.io/latest?base=%s&symbols=%s", fromCurrency, toCurrency)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var rates Rates
	err = json.Unmarshal(body, &rates)
	if err != nil {
		return 0, err
	}

	exchangeRate := rates.Rates[toCurrency]
	convertedAmount := amount * exchangeRate

	return convertedAmount, nil

}

func main() {
	amount := 100.00
	fromCurrency := "USD"
	toCurrency := "EUR"

	convertedAmount, err := convertCurrency(amount, fromCurrency, toCurrency)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%f %s is equivalent to %f %s", amount, fromCurrency, convertedAmount, toCurrency)

}
