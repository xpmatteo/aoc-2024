package day17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name           string
		input          Machine
		expectedA      int
		expectedB      int
		expectedC      int
		expectedOutput string
		skip           bool
	}{
		{
			name:           "2,2 -> set B to 2",
			input:          Machine{Program: []int{2, 2}},
			expectedB:      2,
			expectedOutput: "",
		},
		{
			name:      "2,3 -> set B to 3",
			input:     Machine{Program: []int{2, 3}},
			expectedB: 3,
		},
		{
			name:      "2,6 -> set B to 1",
			input:     Machine{C: 9, Program: []int{2, 6}},
			expectedB: 1,
			expectedC: 9,
		},
		{
			name: "0, 2: divide A by 4",
			input: Machine{
				A:       21,
				Program: []int{0, 2},
			},
			expectedA: 5,
		},
		{
			name: "output literal 0",
			input: Machine{
				Program: []int{5, 0},
			},
			expectedOutput: "0",
		},
		{
			name: "output 0,1,2",
			input: Machine{
				A:       10,
				Program: []int{5, 0, 5, 1, 5, 4},
			},
			expectedOutput: "0,1,2",
			expectedA:      10,
		},
		{
			name: "bxl",
			input: Machine{
				B:       3,
				Program: []int{1, 7},
			},
			expectedB: 4,
		},
		{
			name: "bxc",
			input: Machine{
				B:       1,
				C:       3,
				Program: []int{4, 0},
			},
			expectedB: 2,
			expectedC: 3,
		},
		{
			name: "3rd example",
			input: Machine{
				A:       2024,
				Program: []int{0, 1, 5, 4, 3, 0},
			},
			expectedA:      0,
			expectedOutput: "4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			name: "4th example",
			input: Machine{
				B:       29,
				Program: []int{1, 7},
			},
			expectedB: 26,
		},
		{
			name: "5th example",
			input: Machine{
				B:       2024,
				C:       43690,
				Program: []int{4, 0},
			},
			expectedB: 44354,
			expectedC: 43690,
		},
		{
			name: "sample",
			input: Machine{
				A:       729,
				Program: []int{0, 1, 5, 4, 3, 0},
			},
			expectedOutput: "4,6,3,5,6,3,5,2,1,0",
		},
		{
			name: "real pt 1",
			input: Machine{
				A:       50230824,
				Program: []int{2, 4, 1, 3, 7, 5, 0, 3, 1, 4, 4, 7, 5, 5, 3, 0},
			},
			expectedOutput: "2,1,4,7,6,0,3,1,4",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.skip {
				t.Skip()
			}
			output := test.input.Execute()
			if test.expectedA != 0 {
				assert.Equal(t, test.expectedA, test.input.A)
			}
			if test.expectedB != 0 {
				assert.Equal(t, test.expectedB, test.input.B)
			}
			if test.expectedC != 0 {
				assert.Equal(t, test.expectedC, test.input.C)
			}
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}
