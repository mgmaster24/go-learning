package gol_calculators

import (
	"fmt"

	gol_io "go-learning.com/learning/io"
	gol_prices "go-learning.com/learning/prices"
)

func CalculatePrices() {
	taxRates := []float64 {0, 0.7, 0.1, 0.15}

	for _, tr := range taxRates {
		jobName := fmt.Sprintf("jobResults_%.0f.json", tr*100)
		newTaxIncludedPriceJob := gol_prices.NewTaxIncludedPriceJob(tr,  gol_io.NewFileManager("prices.txt", jobName))
		newTaxIncludedPriceJob.Process()
		newTaxIncludedPriceJob.Output()
	}
}