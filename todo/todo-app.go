package gol_todo

import (
	"fmt"

	gol_io "go-learning.com/learning/io"
)

func Run() {
	todoText := gol_io.GetUserInput("Todo text: ")
	todo, err := New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	gol_io.OutPutAndSave(todo, "todo.json")
}
