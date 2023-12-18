package main

import (
	"fmt"

	banking "go-learning.com/learning/bank"
	"go-learning.com/learning/calculators"
)

func main() {

	progOpt := 0
	fmt.Println("Choose what program to run.")
	fmt.Println("1.) Calculators - Calculate Investment")
	fmt.Println("2.) Calculators - Calculate Profit")
	fmt.Println("3.) Banking - GO Banking App")
	fmt.Print("Selection: ")
	fmt.Scan(&progOpt)

	switch progOpt {
	case 1:
		calculators.CalculateInvestment()
	case 2:
		calculators.CalculateProfit()
	case 3:
		banking.Run()
	}
}
