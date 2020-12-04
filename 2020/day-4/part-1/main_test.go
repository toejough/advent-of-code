package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	input := `
        ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
        byr:1937 iyr:2017 cid:147 hgt:183cm

        iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
        hcl:#cfa07d byr:1929

        hcl:#ae17e1 iyr:2013
        eyr:2024
        ecl:brn pid:760753108 byr:1931
        hgt:179cm

        hcl:#cfa07d eyr:2025 pid:166559648
        iyr:2011 ecl:brn hgt:59in
    `
	expectedOutput := "2"

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

// func TestSolution(t *testing.T) {
//    input, err := ioutil.ReadFile("input.txt")
//    if err != nil {
//        t.Fatalf("Error: %v", err)
//    }

//    expectedOutput := "151"

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