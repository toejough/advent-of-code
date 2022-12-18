package main

import (
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		filename string
		result   string
	}{
		"example": {filename: "example-input.txt", result: "24000"},
		"puzzle":  {filename: "puzzle-input.txt", result: "69528"},
	}

	for name, tc := range tests {
		testCase := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			text := mustReadFileText(testCase.filename)
			lines := strings.Split(text, "\n")
			expectedOutput := testCase.result

			actualOutput, err := solve(lines)
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
		})
	}
}

func FuzzSolve(f *testing.F) {
	f.Add(mustReadFileText("example-input.txt"))
	f.Add(mustReadFileText("puzzle-input.txt"))
	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		lines := strings.Split(s, "\n")
		_, _ = solve(lines)

		// THEN the run is expected to return just fine.
	})
}
