package main

import (
	"fmt"

	banking "go-learning.com/learning/bank"
	"go-learning.com/learning/calculators"
	gol_datastructures "go-learning.com/learning/data_structures"
	gol_io "go-learning.com/learning/io"
	gol_notes "go-learning.com/learning/notes"
	gol_todo "go-learning.com/learning/todo"
	gol_user "go-learning.com/learning/user"
)

func main() {
	progOpt := 0
	fmt.Println("Choose what program to run.")
	fmt.Println("1.) Calculators - Calculate Investment")
	fmt.Println("2.) Calculators - Calculate Profit")
	fmt.Println("3.) Banking - GO Banking App")
	fmt.Println("4.) Stucts Example")
	fmt.Println("5.) SLL Test")
	fmt.Println("6.) Notes App")
	fmt.Println("7.) Todo App")
	fmt.Print("Selection: ")
	fmt.Scan(&progOpt)
	fmt.Println()

	switch progOpt {
	case 1:
		calculators.CalculateInvestment()
	case 2:
		calculators.CalculateProfit()
	case 3:
		banking.Run()
	case 4:
		structExample()
	case 5:
		sllTests()
	case 6:
		notesApp()
	case 7:
		todoApp()
	}
}

func notesApp() {
	fmt.Println("Welcome to the GOL Notes App!")
	fmt.Println("Loading existing notes...")
	notes, err := gol_notes.Load("notes.json")
	if err != nil {
		fmt.Println("Unable to read notes from file...")
		fmt.Printf("Error is %v\n", err)
		fmt.Println("Will start application with empty notes collection")
	} else {
		fmt.Printf("Loaded %v notes\n", len(notes))
		notes.Print()
	}

	shouldContinue := "y"
	fmt.Print("Would you like to add a new note? (y/n)")
	fmt.Scanln(&shouldContinue)
	if shouldContinue == "y" {
		getNotes(&notes)
	}

	outPutAndSave(notes, "notes.json")

	fmt.Println("Goodbye!")
}

func getNotes(notes *gol_notes.Notes) {
	for {
		fmt.Println("Please add a note...")
		title := getUserInput("Note Title:")
		content := getUserInput("Note Content:")

		note, err := gol_notes.New(title, content)
		if err != nil {
			fmt.Println(err)
			return
		}

		notes.Add(note)

		shouldContinue := "y"
		fmt.Print("Would you like to add another? (y/n)")
		fmt.Scanln(&shouldContinue)
		if shouldContinue == "n" {
			break
		}
	}
}

func todoApp() {
	todoText := getUserInput("Todo text: ")
	todo, err := gol_todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	outPutAndSave(todo, "todo.json")
}

func outPutAndSave(noteOperator gol_notes.NoteOps, fileName string) {
	noteOperator.Print()
	fmt.Println("Saving...")
	err := noteOperator.Save(fileName)
	if err != nil {
		fmt.Println(err)
	}
}

func getUserInput(prompt string) (val string) {
	fmt.Printf("%s ", prompt)

	val, err := gol_io.GetLongStringFromStdIn()
	if err != nil {
		fmt.Println(err)
	}

	return
}

func structExample() {
	firstName := getUserData("Please enter your firstName: ")
	lastName := getUserData("Please enter your lastName: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user, err := gol_user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := gol_user.NewAdmin("admin@golearning.com", "ASDSAFASDDFASD")

	user.OutputUserData()
	user.ClearUserName()
	user.OutputUserData()
	admin.OutputUserData()
}

func getUserData(context string) (val string) {
	fmt.Print(context)
	fmt.Scanln(&val)
	return
}

func sllTests() {
	var sll gol_datastructures.SinglyLinkedList[string]
	sll.PushBack("This")
	sll.PushBack("is")
	sll.PushBack("a test")
	sll.Print()
	fmt.Println()

	sll.PushFront("front")
	sll.PushFront("to")
	sll.PushFront("Nodes")
	sll.PushFront("Adding")
	sll.Print()
	fmt.Println()

	sll.InsertBefore("is", "has/")
	sll.Print()
	fmt.Println()

	sll.InsertAfter("a test", "OMG")
	sll.Print()
	fmt.Println()

	sll.Delete("Adding")
	sll.Delete("This")
	sll.Print()
	fmt.Println()
}
