package main

import (
	"fmt"
	"sort"
	"strconv"

	gol_banking "go-learning.com/learning/bank"
	gol_calculators "go-learning.com/learning/calculators"
	gol_datastructures "go-learning.com/learning/data_structures"
	gol_exercises "go-learning.com/learning/exercises"
	gol_io "go-learning.com/learning/io"
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
		3: {"Calculators - Calculate Prices", gol_calculators.CalculatePrices},
		4: {"Banking - GO Banking App", gol_banking.Run},
		5: {"Stucts Example", gol_exercises.RunStructsExercise},
		6: {"SLL Test", gol_datastructures.RunSLLTests},
		7: {"Notes App", gol_notes.Run},
		8: {"Todo App", gol_todo.Run},
		9: {"Fun with Slices", gol_exercises.RunSlicesExcercise},
		10: {"Fun with Maps", gol_exercises.RunMapsExercise},
	}

	sortedKeys := getSortedKeys(apps)
	minSelection, maxSelection := sortedKeys[0], sortedKeys[len(sortedKeys) - 1]
	for {
		printApps(apps, sortedKeys)
		input := gol_io.GetUserInput("Selection (q to quit): ")
		fmt.Println()
		if input == "q" {
			fmt.Println("Exiting...")
			return
		}

		option, e := strconv.Atoi(input);
		if (option < minSelection || option > maxSelection) || e != nil {
			fmt.Println("Invalid selection... Please try again.")
		} else {
			app := apps[option]
			fmt.Printf("Running - %s\n", app.title)
			app.execute()
		}

		fmt.Println()
	}
}

func getSortedKeys(apps map[int]App) []int{
	keys := make([]int, 0, len(apps))
	for k := range apps {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	return keys
}

func printApps(apps map[int]App, keys []int) {
	for _, k := range keys {
		fmt.Printf("%v) %s\n", k, apps[k].title)
	}
}
