package main

import (
	"fmt"
	"slices"
	"strconv"

	gol_banking "go-learning.com/learning/bank"
	gol_calculators "go-learning.com/learning/calculators"
	gol_datastructures "go-learning.com/learning/data_structures"
	gol_exercises "go-learning.com/learning/exercises"
	gol_notes "go-learning.com/learning/notes"
	gol_todo "go-learning.com/learning/todo"
)

type App struct {
	title   string
	execute func()
}

func main() {
	apps := map[int]App{
		1:  {"Calculators - Calculate Investment", gol_calculators.CalculateInvestment},
		2:  {"Calculators - Calculate Profit", gol_calculators.CalculateProfit},
		3:  {"Calculators - Calculate Prices", gol_calculators.CalculatePrices},
		4:  {"Banking - GO Banking App", gol_banking.Run},
		5:  {"Stucts Example", gol_exercises.RunStructsExercise},
		6:  {"SLL Test", gol_datastructures.RunSLLTests},
		7:  {"Notes App", gol_notes.Run},
		8:  {"Todo App", gol_todo.Run},
		9:  {"Fun with Slices", gol_exercises.RunSlicesExcercise},
		10: {"Fun with Maps", gol_exercises.RunMapsExercise},
		11: {"Functions Deep Dive", gol_exercises.RunFunctionsExercise},
	}

	var option int
	for {
		printApps(apps)
		fmt.Print("Select and option from above or q to quit: ")
		var input string
		fmt.Scan(&input)
		fmt.Println()
		if input == "q" {
			fmt.Println("Exiting...")
			return
		}

		option, _ = strconv.Atoi(input)

		if validateSelection(apps, option) {
			apps[option].execute()
		} else {
			fmt.Println("Invalid selection... Please try again.")
		}

		fmt.Println()
	}
}

func getSortedKeys(apps map[int]App) []int {
	keys := make([]int, 0, len(apps))
	for k := range apps {
		keys = append(keys, k)
	}

	slices.Sort(keys)
	return keys
}

func validateSelection(apps map[int]App, selection int) bool {
	keys := getSortedKeys(apps)
	min := keys[0]
	max := keys[len(keys)-1]
	return selection >= min && selection <= max
}

func printApps(apps map[int]App) {
	keys := getSortedKeys(apps)

	for _, k := range keys {
		fmt.Printf("%v) %s\n", k, apps[k].title)
	}
}
