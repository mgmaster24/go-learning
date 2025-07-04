package clockwisematrixprint

import (
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	input := [][]int{{2, 3, 4, 8}, {5, 7, 9, 12}, {1, 0, 6, 10}}
	fmt.Printf("Input: %v\n", input)
	out := clockwiseMatrixFlatten(input)
	outputStr := getOutputStr(out)
	fmt.Printf("Output: %s\n", outputStr)
}

func getOutputStr(output []int) string {
	outputVals := make([]string, len(output))
	for i := 0; i < len(output); i++ {
		outputVals[i] = strconv.Itoa(output[i])
	}

	return strings.Join(outputVals, ",")
}

func clockwiseMatrixFlatten(input [][]int) []int {
	if len(input) == 0 || len(input[0]) == 0 {
		return []int{}
	}
	top, left, bottom, right := 0, 0, len(input)-1, len(input[0])-1
	var result []int

	// while indices are within valid matrix range
	for top <= bottom && left <= right {
		// left to right
		for col := left; col <= right; col++ {
			result = append(result, input[top][col])
		}
		// move top down
		top++

		// Top to bottom right
		for row := top; row <= bottom; row++ {
			result = append(result, input[row][right])
		}
		// move right back
		right--

		// Print bottom row
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, input[bottom][col])
			}
			// move bottom up
			bottom--
		}

		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, input[row][left])
			}

			// move left in
			left++
		}
	}

	return result
}
