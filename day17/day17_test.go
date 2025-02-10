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
				Program: []int{5, 0, 5, 1, 5, 2},
			},
			expectedOutput: "0,1,2",
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.skip {
				t.Skip()
			}
			output := test.input.Execute()
			assert.Equal(t, test.expectedA, test.input.A)
			assert.Equal(t, test.expectedB, test.input.B)
			assert.Equal(t, test.expectedC, test.input.C)
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}
