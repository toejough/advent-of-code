package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestAnswers(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		result string
	}{
		"day1-part1-example": {result: "24000"},
		"day1-part1-puzzle":  {result: "69528"},
		"day1-part2-example": {result: "45000"},
		"day1-part2-puzzle":  {result: "206152"},
		"day2-part1-example": {result: "15"},
		"day2-part1-puzzle":  {result: "13809"},
		"day2-part2-example": {result: "12"},
		"day2-part2-puzzle":  {result: "12316"},
		"day3-part1-example": {result: "157"},
	}

	for name, tc := range tests {
		testCase := tc
		testCaseName := name

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			testNameParts := strings.Split(testCaseName, "-")
			day := testNameParts[0]
			part := testNameParts[1]
			kind := testNameParts[2]
			filename := fmt.Sprintf("%s-input-%s.txt", day, kind)

			text, err := readFileText(filename)
			if err != nil {
				log.Fatal(err)
			}

			expectedOutput := testCase.result

			actualOutput, err := solve(day, part, text)
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
	for _, inputType := range []string{"example", "puzzle"} {
		text, err := readFileText(fmt.Sprintf("day1-input-%s.txt", inputType))
		if err != nil {
			log.Fatal(err)
		}
		f.Add(text)
	}
	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		_, _ = solveDay1Part1(s)

		// THEN the run is expected to return just fine.
	})
}

func FuzzDay1Part2(f *testing.F) {
	for _, inputType := range []string{"example", "puzzle"} {
		text, err := readFileText(fmt.Sprintf("day1-input-%s.txt", inputType))
		if err != nil {
			log.Fatal(err)
		}
		f.Add(text)
	}
	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		_, _ = solveDay1Part2(s)

		// THEN the run is expected to return just fine.
	})
}

func FuzzDay2Part1(f *testing.F) {
	for _, inputType := range []string{"example", "puzzle"} {
		text, err := readFileText(fmt.Sprintf("day2-input-%s.txt", inputType))
		if err != nil {
			log.Fatal(err)
		}
		f.Add(text)
	}

	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		_, _ = solveDay2Part1(s)

		// THEN the run is expected to return just fine.
	})
}

func FuzzDay2Part2(f *testing.F) {
	for _, inputType := range []string{"example", "puzzle"} {
		text, err := readFileText(fmt.Sprintf("day2-input-%s.txt", inputType))
		if err != nil {
			log.Fatal(err)
		}
		f.Add(text)
	}

	f.Fuzz(func(_ *testing.T, s string) {
		// WHEN the program is called with the input
		_, _ = solveDay2Part2(s)

		// THEN the run is expected to return just fine.
	})
}