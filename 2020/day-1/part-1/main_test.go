package main

import (
	"io/ioutil"
	"testing"
)

func TestExample(t *testing.T) {
	input := `
        1721
        979
        366
        299
        675
        1456
    `
	expectedOutput := "514579"

	actualOutput, err := solve(input)
	if err != nil {
		t.Fatalf("Expected no errors, but got '%v'\n", err)
	}

	if expectedOutput != actualOutput {
		t.Fatalf(
			"Expected output was '%v', but we got '%v' instead\n",
			expectedOutput,
			actualOutput,
		)
	}
}

func TestSolution(t *testing.T) {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expectedOutput := "842016"

	actualOutput, err := solve(string(input))
	if err != nil {
		t.Fatalf("Expected no errors, but got '%v'\n", err)
	}

	if expectedOutput != actualOutput {
		t.Fatalf(
			"Expected output was '%v', but we got '%v' instead\n",
			expectedOutput,
			actualOutput,
		)
	}
}
