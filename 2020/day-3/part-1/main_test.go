package main

import (
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
	expectedOutput := "7"

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

// func TestSolution(t *testing.T) {
//    input, err := ioutil.ReadFile("input.txt")
//    if err != nil {
//        t.Fatalf("Error: %v", err)
//    }

//    expectedOutput := "638"

//    actualOutput, err := solve(string(input))
//    if err != nil {
//        t.Fatalf("Expected no errors, but got '%v'\n", err)
//    }

//    if expectedOutput != actualOutput {
//        t.Fatalf(
//            "Expected output was '%v', but we got '%v' instead\n",
//            expectedOutput,
//            actualOutput,
//        )
//    }
//}
