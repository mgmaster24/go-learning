package gol_io

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

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
