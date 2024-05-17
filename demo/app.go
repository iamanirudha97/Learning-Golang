package main

import (
	"example.com/app/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxInclusicePriceJob(taxRate)
		priceJob.Process()
	}
}
