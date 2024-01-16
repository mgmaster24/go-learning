package gol_io

import (
	"fmt"
)

type IoOps interface {
	Save(fileName string) error
	Print()
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
