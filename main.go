package main

import (
	gol_app "go-learning.com/learning/app"
	gol_banking "go-learning.com/learning/bank"
	gol_calculators "go-learning.com/learning/calculators"
	clockwisematrixprint "go-learning.com/learning/clockwise-matrix-print"
	gol_concurrency "go-learning.com/learning/concurrency"
	gol_datastructures "go-learning.com/learning/data_structures"
	gol_eventbooking "go-learning.com/learning/event-booking"
	gol_exercises "go-learning.com/learning/exercises"
	gol_notes "go-learning.com/learning/notes"
	gol_todo "go-learning.com/learning/todo"
	gol_tourofgo "go-learning.com/learning/tour-of-go"
)

func main() {
	apps := map[int]gol_app.App{
		1: {
			Title:   "Calculators - Calculate Investment",
			Execute: gol_calculators.CalculateInvestment,
		},
		2:  {Title: "Calculators - Calculate Profit", Execute: gol_calculators.CalculateProfit},
		3:  {Title: "Calculators - Calculate Prices", Execute: gol_calculators.CalculatePrices},
		4:  {Title: "Banking - GO Banking App", Execute: gol_banking.Run},
		5:  {Title: "Stucts Example", Execute: gol_exercises.RunStructsExercise},
		6:  {Title: "SLL Test", Execute: gol_datastructures.RunSLLTests},
		7:  {Title: "Notes App", Execute: gol_notes.Run},
		8:  {Title: "Todo App", Execute: gol_todo.Run},
		9:  {Title: "Fun with Slices", Execute: gol_exercises.RunSlicesExcercise},
		10: {Title: "Fun with Maps", Execute: gol_exercises.RunMapsExercise},
		11: {Title: "Functions Deep Dive", Execute: gol_exercises.RunFunctionsExercise},
		12: {Title: "Concurrency", Execute: gol_concurrency.Run},
		13: {Title: "Spriral Print Matrix - Clockwise", Execute: clockwisematrixprint.Run},
		14: {Title: "Rest API", Execute: gol_eventbooking.Run},
		15: {Title: "Tour of GO", Execute: gol_tourofgo.RunTOG},
	}

	gol_app.RunAppWithSelections(apps, "GO Learning App")
}
