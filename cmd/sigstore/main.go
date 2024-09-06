package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const apiKey = "45a77d7eb900420cab83790005d08553"
const apiURL = "https://openexchangerates.org/api/latest.json?app_id=" + apiKey

// Struct to hold API response
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

func getExchangeRates() (ExchangeRates, error) {
	// Make the API request
	resp, err := http.Get(apiURL)
	if err != nil {
		return ExchangeRates{}, fmt.Errorf("error fetching exchange rates: %v", err)
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var rates ExchangeRates
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return ExchangeRates{}, fmt.Errorf("error parsing exchange rates: %v", err)
	}

	return rates, nil
}

func convertCurrency(amount float64, fromCurrency, toCurrency string, rates ExchangeRates) (float64, error) {
	fromRate, exists := rates.Rates[fromCurrency]
	if !exists {
		return 0, fmt.Errorf("unsupported currency: %s", fromCurrency)
	}

	toRate, exists := rates.Rates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("unsupported currency: %s", toCurrency)
	}

	// Convert the amount
	convertedAmount := (amount / fromRate) * toRate
	return convertedAmount, nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run currency_converter.go <amount> <from_currency> <to_currency>")
		return
	}

	// Parse user input
	amountStr, fromCurrency, toCurrency := os.Args[1], os.Args[2], os.Args[3]
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Printf("Invalid amount: %s\n", amountStr)
		return
	}

	// Fetch exchange rates
	rates, err := getExchangeRates()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert currency
	convertedAmount, err := convertCurrency(amount, fromCurrency, toCurrency, rates)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Output result
	fmt.Printf("%.2f %s is %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
