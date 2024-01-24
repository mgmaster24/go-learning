package gol_prices

import (
	"fmt"

	gol_conversions "go-learning.com/learning/conversions"
	gol_io "go-learning.com/learning/io"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64            `json:"taxRate"`
	Prices            []float64          `json:"prices"`
	TaxIncludedPrices map[string]float64 `json:"taxInclPrices"`
	ioHandler         gol_io.IoHandler
}

type PriceResult struct {
	Job *TaxIncludedPriceJob
	Err error
}

func NewTaxIncludedPriceJob(taxRate float64, ioHandler gol_io.IoHandler) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		ioHandler: ioHandler,
	}
}

func (tipj *TaxIncludedPriceJob) Process() error {
	err := tipj.loadData()
	if err != nil {
		return err
	}

	results := make(map[string]float64)
	for _, p := range tipj.Prices {
		ps := fmt.Sprintf("%.2f", p)
		results[ps] = gol_conversions.TruncateFloat(p * (1 + tipj.TaxRate))
	}

	tipj.TaxIncludedPrices = results

	return tipj.ioHandler.WriteResult(tipj)
}

func (tipj *TaxIncludedPriceJob) ProcessAsync(processed chan PriceResult) {
	err := tipj.loadData()
	if err != nil {
		processed <- PriceResult{
			Job: nil,
			Err: err,
		}
	}
	results := make(map[string]float64)
	for _, p := range tipj.Prices {
		ps := fmt.Sprintf("%.2f", p)
		results[ps] = gol_conversions.TruncateFloat(p * (1 + tipj.TaxRate))
	}

	tipj.TaxIncludedPrices = results

	err = tipj.ioHandler.WriteResult(tipj)
	if err != nil {
		processed <- PriceResult{
			Job: nil,
			Err: err,
		}
	}

	processed <- PriceResult{
		Job: tipj,
		Err: nil,
	}
}

func (tipj *TaxIncludedPriceJob) loadData() error {
	vals, err := tipj.ioHandler.ReadLines()
	if err != nil {
		return err
	}

	tipj.Prices = vals
	return nil
}
