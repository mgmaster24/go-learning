package gol_io

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type FileManager struct {
	InputFileName string
	OutputFileName string
}

func NewFileManager(inputFile, outputFile string) FileManager{
	return FileManager{
		InputFileName: inputFile,
		OutputFileName: outputFile,
	}
}

func (fileManager FileManager) WriteResult(data any) {
	SaveJsonData(fileManager.OutputFileName, data)
}

func (fileManager FileManager) ReadLines() ([]float64, error) {
	vals, err := ReadFloatArrayFromFile(fileManager.InputFileName)
	return vals, err
}

func WriteFloatValToFile(fileName string, value float64) {
	WriteStringToFile(fileName, fmt.Sprint(value))
}

func WriteStringToFile(fileName string, value string) {
	os.WriteFile(fileName, []byte(value), 0644)
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

func ReadFloatArrayFromFile(file string) (vals []float64, err error) {
	f, err := os.Open(file)
	if (err != nil) {
		err = fmt.Errorf("Could not open file! - %w", err)
		f.Close()
		return
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		price, err := strconv.ParseFloat(line, 64)
		if (err != nil) {
			err = fmt.Errorf("Could not parse string as float64 - %w", err)
			break;
		}

		vals = append(vals, price)
	}

	err = scanner.Err()
	if (err != nil) {
		err = fmt.Errorf("Error reading file - %w", err)
		vals = nil
	}

	f.Close()

	return
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