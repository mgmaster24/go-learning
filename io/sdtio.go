package gol_io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CMDHandler struct {
}

func (cmd CMDHandler) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func (cmd CMDHandler) ReadLines() ([]float64, error) {
	fmt.Println("Please enter prices or q to quit")
	var prices []float64
	for {
		var price string
		ScanForInput("Add Price:", &price)
		if price == "q" {
			break
		}

		floatVal, err := strconv.ParseFloat(price, 0064)
		if err != nil {
			return nil, err
		}

		prices = append(prices, floatVal)
	}

	return prices, nil
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

func ScanForInput[T any](prompt string, val *T) {
	fmt.Print(prompt)
	fmt.Scan(val)
}
