package gol_io

import (
	"fmt"
)

type IoOps interface {
	Save(fileName string) error
	Print()
}

type IoHandler interface {
	WriteResult(data any) error
	ReadLines() ([]float64, error)
}

func OutPutAndSave(ioOperator IoOps, fileName string) {
	ioOperator.Print()
	fmt.Println("Saving...")
	err := ioOperator.Save(fileName)
	if err != nil {
		fmt.Println(err)
	}
}

func GetUserInput(prompt string) (val string) {
	fmt.Printf("%s", prompt)

	val, err := GetLongStringFromStdIn()
	if err != nil {
		fmt.Println(err)
	}

	return
}
