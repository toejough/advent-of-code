package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAnswers(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		filename string
		solver   func(string) (string, error)
		result   string
	}{
		"day1-part1-example": {solver: solveDay1Part1, result: "24000"},
		"day1-part1-puzzle":  {solver: solveDay1Part1, result: "69528"},
		"day1-part2-example": {solver: solveDay1Part2, result: "45000"},
		"day1-part2-puzzle":  {solver: solveDay1Part2, result: "206152"},
		"day2-part1-example": {solver: solveDay2Part1, result: "15"},
		"day2-part1-puzzle":  {solver: solveDay2Part1, result: "13809"},
		"day2-part2-example": {solver: solveDay2Part2, result: "12"},
		"day2-part2-puzzle":  {solver: solveDay2Part2, result: "12316"},
	}

	for name, tc := range tests {
		testCase := tc
		testCaseName := name

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testNameParts := strings.Split(testCaseName, "-")
			filename := fmt.Sprintf("%s-input-%s.txt", testNameParts[0], testNameParts[2])

			text := mustReadFileText(filename)
			expectedOutput := testCase.result

			actualOutput, err := testCase.solver(text)
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

func FuzzDay1Part1(f *testing.F) {
	f.Add(mustReadFileText("day1-input-example.txt"))
	f.Add(mustReadFileText("day1-input-puzzle.txt"))
	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		_, _ = solveDay1Part1(s)

		// THEN the run is expected to return just fine.
	})
}

func FuzzDay1Part2(f *testing.F) {
	f.Add(mustReadFileText("day1-input-example.txt"))
	f.Add(mustReadFileText("day1-input-puzzle.txt"))
	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		_, _ = solveDay1Part2(s)

		// THEN the run is expected to return just fine.
	})
}

func FuzzDay2Part1(f *testing.F) {
	f.Add(mustReadFileText("day2-input-example.txt"))

	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		_, _ = solveDay2Part1(s)

		// THEN the run is expected to return just fine.
	})
}

func FuzzDay2Part2(f *testing.F) {
	f.Add(mustReadFileText("day2-input-example.txt"))

	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		_, _ = solveDay2Part2(s)

		// THEN the run is expected to return just fine.
	})
}
