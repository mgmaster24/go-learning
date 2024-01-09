package gol_banking

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	gol_io "go-learning.com/learning/io"
)

func Run() {
	const fileName = "balance.txt"
	fmt.Println("Welcome to GO Bank!")
	fmt.Println(fmt.Sprintf("Contact us at %s", randomdata.PhoneNumber()))
	accountBalance, err := gol_io.GetFloatValueFromFile(fileName)
	if err != nil {
		fmt.Println("Error retrieving your balance.  Please speak to your admin.")
		fmt.Println("**Initial value with be zero for this session**")
		fmt.Println()
		//panic()
	}

	showSelections()

	for {
		option := 0

		fmt.Print("Selection: ")
		fmt.Scan(&option)
		if option == 4 {
			break
		}

		fmt.Println("Option selected: ", option)

		switch option {
		case 1:
			fmt.Println(fmt.Sprintf("Current Balance is: $%.2f", accountBalance))
		case 2:
			adjustAccountBalance(&accountBalance, true)
		case 3:
			adjustAccountBalance(&accountBalance, false)
		}

		fmt.Println()
	}

	gol_io.WriteFloatValToFile(fileName, accountBalance)
	fmt.Println("Exiting GO Bank App!")
}

func adjustAccountBalance(accountBalance *float64, deposit bool) {
	var adjustment = 0.0
	var action string = "deposit"
	if !deposit {
		action = "withdraw"
	}

	for !validateAdjusment(*accountBalance, action, &adjustment, deposit) {
	}

	if deposit {
		*accountBalance += adjustment
	} else {
		*accountBalance -= adjustment
	}

	var result string = "Deposited:"
	if !deposit {
		result = "Withdrew:"
	}

	fmt.Println(fmt.Sprintf("%v $%.2f", result, adjustment))
	fmt.Println(fmt.Sprintf("New Balance is: $%.2f", *accountBalance))
}

func validateAdjusment(
	currentBalance float64,
	action string,
	adjustment *float64,
	isDeposit bool) bool {
	fmt.Printf("Please enter the amount to %v: ", action)
	fmt.Scan(adjustment)

	if *adjustment <= 0.0 {
		fmt.Println("Please enter a valid amount. ")
		return false
	}

	if !isDeposit && *adjustment > currentBalance {
		fmt.Println("Attempting to withdraw more than the current balance.")
		fmt.Printf("Please enter a value less than or equal to $%.2f\n", currentBalance)
		return false
	}

	return true
}
