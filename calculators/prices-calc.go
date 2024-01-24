package gol_calculators

import (
	"fmt"

	gol_io "go-learning.com/learning/io"
	gol_prices "go-learning.com/learning/prices"
)

func CalculatePrices() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan gol_prices.PriceResult, len(taxRates))

	for idx, tr := range taxRates {
		doneChans[idx] = make(chan gol_prices.PriceResult)
		jobName := fmt.Sprintf("jobResults_%.0f.json", tr*100)
		fileIoHandler := gol_io.NewFileHandler("prices.txt", jobName)
		//cmdIoHandler := gol_io.CMDHandler{}
		newTaxIncludedPriceJob := gol_prices.NewTaxIncludedPriceJob(tr, fileIoHandler)
		// var err error = nil
		go newTaxIncludedPriceJob.ProcessAsync(doneChans[idx])
	}

	for _, doneChan := range doneChans {
		result := <-doneChan
		OutputJobOrErr(&result)
	}
}

func OutputJobOrErr(result *gol_prices.PriceResult) {
	if result.Err != nil {
		fmt.Printf("An error occurred while processing the job.\n%v\nExiting...", result.Err)
	} else {
		for k, v := range result.Job.TaxIncludedPrices {
			fmt.Printf("Price: %s, Tax Rate: %.2f, Adjusted: %.2f\n", k, result.Job.TaxRate, v)
		}
	}
}
