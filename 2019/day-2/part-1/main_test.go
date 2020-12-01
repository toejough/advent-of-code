package main

import (
	"strconv"
	"testing"
)

func TestExecIntCode(t *testing.T) {
	testData := []struct {
		input  string
		output string
	}{
		{
			input:  "1,0,0,0,99",
			output: "2,0,0,0,99",
		},
		{
			input:  "2,3,0,3,99",
			output: "2,3,0,6,99",
		},
		{
			input:  "2,4,4,5,99,0",
			output: "2,4,4,5,99,9801",
		},
		{
			input:  "1,1,1,4,99,5,6,0,99",
			output: "30,1,1,4,2,5,6,0,99",
		},
	}
	for i, td := range testData {
		td := td

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			expected := td.output
			actual := execIntCode(td.input)

			if expected != actual {
				t.Fatalf(
					"resulting intcode (%v) was not as expected (%v) for input (%v)",
					actual,
					expected,
					td.input,
				)
			}
		})
	}
}

func TestMain(t *testing.T) {
	t.Parallel()

	correctAnswer := "6568671"
	result := do()

	if result != correctAnswer {
		t.Fatalf("Actual result (%v) was incorrect (should have been %v)", result, correctAnswer)
	}
}
