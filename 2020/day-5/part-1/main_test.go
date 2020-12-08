package main

import (
	"io/ioutil"
	"testing"
)

func TestExample(t *testing.T) {
	input := `
        FBFBBFFRLR
        BFFFBBFRRR
        FFFBBBFRRR
        BBFFBBFRLL
    `
	expectedOutput := "820"

	actualOutput, err := solve(input)
	if err != nil {
		t.Fatalf("expected no error on solve, but got '%v'\n", err)
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

	expectedOutput := "980"

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
