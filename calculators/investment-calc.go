package gol_calculators

import (
	"fmt"
	"math"

	gol_io "go-learning.com/learning/io"
)

func CalculateInvestment() {
	var investmentAmount, years, expectedROI float64
	const inflationRate = 2.5

	gol_io.ScanForInput("Enter the investment amount: ", &investmentAmount)
	gol_io.ScanForInput("Enter the years to invest: ", &years)
	gol_io.ScanForInput("Enter the expected return on investment: ", &expectedROI)

	futureVal := investmentAmount * math.Pow(1+expectedROI/100, years)
	futureRealVal := futureVal / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future Value of Investment - ", fmt.Sprintf("%.2f", futureVal))
	fmt.Println(fmt.Sprintf("Future Inflation Adjusted Value of Investment (Inflation Rate(%%): %v) -", inflationRate), fmt.Sprintf("%.2f", futureRealVal))
}