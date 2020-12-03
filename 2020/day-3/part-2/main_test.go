package main

import (
	"io/ioutil"
	"testing"
)

func TestExample(t *testing.T) {
	input := `
        ..##.......
        #...#...#..
        .#....#..#.
        ..#.#...#.#
        .#...##..#.
        ..#.##.....
        .#.#.#....#
        .#........#
        #.##...#...
        #...##....#
        .#..#...#.#
    `
	expectedOutput := "336"

	actualOutput := solve(input)

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

	expectedOutput := "7540141059"

	actualOutput := solve(string(input))

	if expectedOutput != actualOutput {
		t.Fatalf(
			"Expected output was '%v', but we got '%v' instead\n",
			expectedOutput,
			actualOutput,
		)
	}
}
