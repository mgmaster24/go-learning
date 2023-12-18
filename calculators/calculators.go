package calculators

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func CalculateInvestment() {
	var investmentAmount, years, expectedROI float64
	const inflationRate = 2.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the years to invest: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return on investment: ")
	fmt.Scan(&expectedROI)

	futureVal := investmentAmount * math.Pow(1+expectedROI/100, years)
	futureRealVal := futureVal / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value of Investment - ", fmt.Sprintf("%.2f", futureVal))
	fmt.Println(fmt.Sprintf("Future Inflation Adjusted Value of Investment (Inflation Rate(%%): %v) -", inflationRate), fmt.Sprintf("%.2f", futureRealVal))
}

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
	fmt.Print(context)
	fmt.Scan(&validInput)
	if validInput <= 0 {
		return -1, errors.New("Values can't be 0 or negative...exiting")
	}

	return validInput, nil
}

func storeCalculatedResults(fileName string, ebt, eat, ratio float64) {
	var output string = fmt.Sprintf("EBT:%.2f\nEAT:%.2f\nRatio:%.2f", ebt, eat, ratio)
	os.WriteFile(fileName, []byte(output), 0644)
}
