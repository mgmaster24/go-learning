package gol_calculators

import (
	"errors"
	"fmt"

	gol_io "go-learning.com/learning/io"
)

func CalculateProfit() {
	revenue, err := getProfitVar("Enter the expected revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getProfitVar("Enter the expenses accrued: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getProfitVar("Enter the current tax rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, eat, ratioEbtToEat := calculateProfitResults(revenue, expenses, taxRate)
	storeCalculatedResults("profits.txt", ebt, eat, ratioEbtToEat)
	fmt.Println("Earning Before Tax -", fmt.Sprintf("%.2f", ebt))
	fmt.Println("Earnings After Tax -", fmt.Sprintf("%.2f", eat))
	fmt.Println(fmt.Sprintf("Ratio EBT/EAT - %.2f", ratioEbtToEat))
}

func getProfitVar(text string) (profitVar float64, err error) {
	return validateInput(text)
}

func calculateProfitResults(revenue, expenses, taxRate float64) (ebt float64, eat float64, ratioEbtToEat float64) {
	ebt = revenue - expenses
	eat = ebt * (1 - taxRate/100)
	ratioEbtToEat = ebt / eat
	return
}

func validateInput(context string) (validInput float64, err error) {
	gol_io.ScanForInput(context, &validInput)
	if validInput <= 0 {
		return -1, errors.New("Values can't be 0 or negative...exiting")
	}

	return validInput, nil
}

func storeCalculatedResults(fileName string, ebt, eat, ratio float64) {
	var output string = fmt.Sprintf("EBT:%.2f\nEAT:%.2f\nRatio:%.2f", ebt, eat, ratio)
	gol_io.WriteStringToFile(fileName, output)
}
