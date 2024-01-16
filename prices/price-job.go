package gol_prices

import (
	"fmt"

	gol_conversions "go-learning.com/learning/conversions"
	gol_io "go-learning.com/learning/io"
)

type TaxIncludedPriceJob struct {
	TaxRate float64 											`json:"taxRate"`
	Prices []float64 											`json:"prices"`
	TaxIncludedPrices map[string]float64 	`json:"taxInclPrices"`
	fileManager gol_io.FileManager
}

func NewTaxIncludedPriceJob(taxRate float64, fileManager gol_io.FileManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		fileManager: fileManager,
	}
}

func (tipj *TaxIncludedPriceJob) Process() {
	tipj.loadData()
	results := make(map[string]float64)
	for _, p := range tipj.Prices {
		ps := fmt.Sprintf("%.2f", p)
		results[ps] = gol_conversions.TruncateFloat(p * (1 + tipj.TaxRate))
	}

	tipj.TaxIncludedPrices = results

	tipj.fileManager.WriteResult(tipj)
}


func (tipj* TaxIncludedPriceJob) loadData() {
	vals, err := tipj.fileManager.ReadLines()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	tipj.Prices = vals
}

func (tipj *TaxIncludedPriceJob) Output() {
	for k,v := range tipj.TaxIncludedPrices {
		fmt.Printf("Price: %s, Tax Rate: %.2f, Adjusted: %.2f\n", k, tipj.TaxRate, v)
	}
}