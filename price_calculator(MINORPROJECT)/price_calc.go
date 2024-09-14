package main

import (
	"github.com/engnrutkarsh/programs/prices/price"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	for _, taxRate := range taxRates {
		priceJob := price.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}
