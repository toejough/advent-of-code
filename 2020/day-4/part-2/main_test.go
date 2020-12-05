package main

import (
	"strings"
	"testing"
)

func TestValid(t *testing.T) {
	cases := `
        byr valid:   2002
        byr invalid: 2003

        hgt valid:   60in
        hgt valid:   190cm
        hgt invalid: 190in
        hgt invalid: 190

        hcl valid:   #123abc
        hcl invalid: #123abz
        hcl invalid: 123abc

        ecl valid:   brn
        ecl invalid: wat

        pid valid:   000000001
        pid invalid: 0123456789
    `

	lines := strings.Split(cases, "\n")
	stripped := stripAll(lines)
	nonEmpty := skipEmpty(stripped)
	f := func(r rune) bool {
		switch r {
		case ' ', ':':
			return true
		}

		return false
	}

	const numTestCaseParts = 3

	for _, testcase := range nonEmpty {
		parts := strings.FieldsFunc(testcase, f)
		if len(parts) != numTestCaseParts {
			t.Fatalf(
				"Found a malformed testcase ('%v'), which parsed into %v parts instead of 3\n",
				testcase,
				len(parts),
			)
		}

		fieldName := parts[0]
		expectedOutput := parts[1] == "valid"
		value := parts[2]

		t.Run(testcase, func(s *testing.T) {
			s.Parallel()
			actualOutput := isValidField(field{name: fieldName, value: value})
			if expectedOutput != actualOutput {
				s.Fatalf(
					"Expected output was '%v', but we got '%v' instead\n",
					expectedOutput,
					actualOutput,
				)
			}
		})
	}
}

func TestExampleValid(t *testing.T) {
	input := `
        pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
        hcl:#623a2f

        eyr:2029 ecl:blu cid:129 byr:1989
        iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

        hcl:#888785
        hgt:164cm byr:2001 iyr:2015 cid:88
        pid:545766238 ecl:hzl
        eyr:2022

        iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
    `
	expectedOutput := "4"

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

func TestExampleInvalid(t *testing.T) {
	input := `
        eyr:1972 cid:100
        hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

        iyr:2019
        hcl:#602927 eyr:1967 hgt:170cm
        ecl:grn pid:012533040 byr:1946

        hcl:dab227 iyr:2012
        ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

        hgt:59cm ecl:zzz
        eyr:2038 hcl:74454a iyr:2023
        pid:3556412378 byr:2007
    `
	expectedOutput := "0"

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
// input, err := ioutil.ReadFile("input.txt")
// if err != nil {
// t.Fatalf("Error: %v", err)
//}

// expectedOutput := "206"

// actualOutput, err := solve(string(input))
// if err != nil {
// t.Fatalf("Expected no errors, but got '%v'\n", err)
//}

// if expectedOutput != actualOutput {
// t.Fatalf(
//"Expected output was '%v', but we got '%v' instead\n",
// expectedOutput,
// actualOutput,
//)
//}
//}
