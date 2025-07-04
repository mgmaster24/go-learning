package clockwisematrixprint

import "testing"

func TestCwMatPrint(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected string
	}{
		{
			name: "3x4 matrix",
			input: [][]int{
				{2, 3, 4, 8},
				{5, 7, 9, 12},
				{1, 0, 6, 10},
			},
			expected: "2,3,4,8,12,10,6,0,1,5,7,9",
		},
		{
			name: "3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: "1,2,3,6,9,8,7,4,5",
		},
		{
			name: "2x4 matrix",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
			expected: "1,2,3,4,8,7,6,5",
		},
		{
			name: "1x1 matrix",
			input: [][]int{
				{42},
			},
			expected: "42",
		},
		{
			name: "1x4 matrix",
			input: [][]int{
				{1, 2, 3, 4},
			},
			expected: "1,2,3,4",
		},
		{
			name: "4x1 matrix",
			input: [][]int{
				{1}, {2}, {3}, {4},
			},
			expected: "1,2,3,4",
		},
		{
			name:     "empty matrix",
			input:    [][]int{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := clockwiseMatrixFlatten(tt.input)
			actualStr := getOutputStr(actual)
			if actualStr != tt.expected {
				t.Errorf("Expected: %q, Actual: %q", tt.expected, actualStr)
			}
		})
	}
}
