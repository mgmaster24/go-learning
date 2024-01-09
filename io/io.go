package gol_io

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func GetLongStringFromStdIn() (val string, err error) {
	cmdReader := bufio.NewReader(os.Stdin)
	val, err = cmdReader.ReadString('\n')
	if err != nil {
		return "", err
	}

	val = strings.TrimSuffix(val, "\n")

	// Windows specific - remove carriage return
	val = strings.TrimSuffix(val, "\r")

	return val, nil

}

func WriteFloatValToFile(fileName string, value float64) {
	os.WriteFile(fileName, []byte(fmt.Sprint(value)), 0644)
}

func GetFloatValueFromFile(fileName string) (float64, error) {
	fileContents, err := os.ReadFile(fileName)
	var value float64
	if err != nil {
		return 0, errors.New("Unable to load the account balance from file!")
	} else {
		valueAsStr := string(fileContents)
		value, err = strconv.ParseFloat(valueAsStr, 64)
		if err != nil {
			return 0, errors.New("Unable to parse value saved in balance.txt")
		}
	}

	return value, nil
}

func LoadJsonData[T any](fileName string) (T, error) {
	var val T
	fileByteContent, err := os.ReadFile(fileName)
	if err != nil {
		return val, err
	}

	err = json.Unmarshal(fileByteContent, &val)
	if err != nil {
		return val, err
	}

	return val, nil
}

func SaveJsonData(fileName string, val any) error {
	json, err := json.Marshal(val)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, json, 0644)
	return nil
}

func GetUserInput(prompt string) (val string) {
	fmt.Printf("%s ", prompt)

	val, err := GetLongStringFromStdIn()
	if err != nil {
		fmt.Println(err)
	}

	return
}
