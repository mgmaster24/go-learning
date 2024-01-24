package gol_io

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type FileHandler struct {
	InputFileName  string
	OutputFileName string
}

func NewFileHandler(inputFile, outputFile string) FileHandler {
	return FileHandler{
		InputFileName:  inputFile,
		OutputFileName: outputFile,
	}
}

func (fileHandler FileHandler) WriteResult(data any) error {
	return SaveJsonData(fileHandler.OutputFileName, data)
}

func (fileHandler FileHandler) ReadLines() ([]float64, error) {
	vals, err := ReadFloatArrayFromFile(fileHandler.InputFileName)
	return vals, err
}

func WriteFloatValToFile(fileName string, value float64) error {
	return WriteStringToFile(fileName, fmt.Sprint(value))
}

func WriteStringToFile(fileName string, value string) error {
	return os.WriteFile(fileName, []byte(value), 0644)
}

func GetFloatValueFromFile(fileName string) (float64, error) {
	fileContents, err := os.ReadFile(fileName)
	var value float64
	if err != nil {
		return 0, fmt.Errorf("Unable to load the account balance from file!\n%w", err)
	} else {
		valueAsStr := string(fileContents)
		value, err = strconv.ParseFloat(valueAsStr, 64)
		if err != nil {
			return 0, fmt.Errorf("Unable to parse value saved in balance.txt\n%w", err)
		}
	}

	return value, nil
}

func ReadFloatArrayFromFile(file string) ([]float64, error) {
	f, err := os.Open(file)
	if err != nil {
		err = fmt.Errorf("Could not open file! - %w", err)
		return nil, err
	}

	var vals []float64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			err = fmt.Errorf("Could not parse value (%v) as float64 - %w", line, err)
			return nil, err
		}

		vals = append(vals, price)
	}

	err = scanner.Err()
	if err != nil {
		err = fmt.Errorf("Error reading file - %w", err)
		return nil, err
	}

	defer f.Close()

	return vals, nil
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

func SaveJsonData(fileName string, val interface{}) error {
	json, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, json, 0644)
	return err
}
