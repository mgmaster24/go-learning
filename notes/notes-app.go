package gol_notes

import (
	"fmt"

	gol_io "go-learning.com/learning/io"
)

func Run() {
	fmt.Println("Welcome to the GOL Notes App!")
	fmt.Println("Loading existing notes...")
	notes, err := Load("notes.json")
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

	gol_io.OutPutAndSave(notes, "notes.json")

	fmt.Println("Goodbye!")
}

func getNotes(notes *Notes) {
	for {
		fmt.Println("Please add a note...")
		title := gol_io.GetUserInput("Note Title:")
		content := gol_io.GetUserInput("Note Content:")

		note, err := New(title, content)
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
