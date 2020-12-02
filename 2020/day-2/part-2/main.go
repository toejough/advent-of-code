package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func stripAll(items []string) (allStripped []string) {
	for _, item := range items {
		stripped := strings.TrimSpace(item)
		allStripped = append(allStripped, stripped)
	}
	return
}

func skipEmpty(items []string) (nonEmpty []string) {
	for _, item := range items {
		if len(item) > 0 {
			nonEmpty = append(nonEmpty, item)
		}
	}
	return
}

type pwSpec struct {
	pos1     int
	pos2     int
	letter   rune
	password string
}

func allToStructs(items []string) (allStructs []pwSpec, err error) {
	for _, item := range items {
		parts := strings.Split(item, " ")
		numParts := len(parts)
		if numParts != 3 {
			return nil, errors.Errorf("Expected 3 parts, but got %v: %v", numParts, parts)
		}

		limitsSpec, letterSpec, password := parts[0], parts[1], parts[2]

		limits := strings.Split(limitsSpec, "-")
		numLimits := len(limits)
		if numLimits != 2 {
			return nil, errors.Errorf("Expected 2 limits, but got %v: %v", numLimits, limits)
		}

		minLimitStr := limits[0]
		minLimit, err := strconv.Atoi(minLimitStr)
		if err != nil {
			return nil, errors.Wrapf(err, "Converting minLimit '%v' from string to int", minLimitStr)
		}

		maxLimitStr := limits[1]
		maxLimit, err := strconv.Atoi(maxLimitStr)
		if err != nil {
			return nil, errors.Wrapf(err, "Converting maxLimit '%v' from string to int", maxLimitStr)
		}

		numRunes := len(letterSpec)
		if numRunes != 2 {
			return nil, errors.Errorf("Expected 2 runes, but got %v: %v", numRunes, letterSpec)
		}
		letter := rune(letterSpec[0])
		allStructs = append(
			allStructs,
			pwSpec{pos1: minLimit, pos2: maxLimit, letter: letter, password: password},
		)
	}
	return allStructs, nil
}

func isValid(spec pwSpec) (valid bool) {
	inFirst := len(spec.password) >= spec.pos1 && rune(spec.password[spec.pos1-1]) == spec.letter
	inSecond := len(spec.password) >= spec.pos2 && rune(spec.password[spec.pos2-1]) == spec.letter
	if inFirst && inSecond {
		return false
	}
	if inFirst || inSecond {
		return true
	}
	return false
}

func skipInvalid(pwSpecs []pwSpec) (allValid []pwSpec) {
	for _, spec := range pwSpecs {
		if isValid(spec) {
			allValid = append(allValid, spec)
		}
	}
	return
}

func solve(input string) (output string, err error) {
	lines := strings.Split(input, "\n")
	stripped := stripAll(lines)
	nonEmpty := skipEmpty(stripped)
	structItems, err := allToStructs(nonEmpty)
	if err != nil {
		return "", errors.Wrap(err, "Converting input []string to []pwSpec")
	}

	validStructs := skipInvalid(structItems)
	return strconv.Itoa(len(validStructs)), nil
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	output, err := solve(string(input))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(output)
}
