package gol_io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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