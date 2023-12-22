package gol_todo

import (
	"errors"
	"fmt"

	gol_io "go-learning.com/learning/io"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid Note values")
	}

	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Save(fileName string) error {
	return gol_io.SaveJsonData(fileName, todo)
}

func Load(fileName string) (Todo, error) {
	todo, err := gol_io.LoadJsonData[Todo](fileName)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (todo Todo) Print() {
	fmt.Printf("Text: %v\n", todo.Text)
}
