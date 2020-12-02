package main

import "testing"

func TestExample(t *testing.T) {
	input := `
        1721
        979
        366
        299
        675
        1456
    `
	expectedOutput := "241861950"

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
