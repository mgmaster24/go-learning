package main

import (
	"fmt"
	"sort"

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
		1: {"Calculators - Calculate Investment", gol_calculators.CalculateInvestment},
		2: {"Calculators - Calculate Profit", gol_calculators.CalculateProfit},
		3: {"Banking - GO Banking App", gol_banking.Run},
		4: {"Stucts Example", gol_exercises.RunStructsExercise},
		5: {"SLL Test", gol_datastructures.RunSLLTests},
		6: {"Notes App", gol_notes.Run},
		7: {"Todo App", gol_todo.Run},
		8: {"Fun with Slices", gol_exercises.RunSlicesExcercise},
		9: {"Fun with Maps", gol_exercises.RunMapsExercise},
	}

	minSelection, maxSelection := minMaxSelections(apps)
	var option int
	validSelection := option >= minSelection && option <= maxSelection
	for validSelection == false {
		printApps(apps)
		fmt.Print("Selection: ")
		var input byte
		fmt.Scanf("%c", &input)
		fmt.Println()
		if input == 'q' {
			fmt.Println("Exiting...")
			return
		}

		option = int(input - '0')
		validSelection = option >= minSelection && option <= maxSelection
		if !validSelection {
			fmt.Println("Invalid selection... Please try again.")
			fmt.Println()
		}
	}

	apps[option].execute()
}

func minMaxSelections(apps map[int]App) (min int, max int) {
	keys := make([]int, 0, len(apps))
	for k := range apps {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	min = keys[0]
	max = keys[len(keys)-1]
	return
}

func printApps(apps map[int]App) {
	keys := make([]int, 0, len(apps))
	for k := range apps {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%v) %s\n", k, apps[k].title)
	}
}
