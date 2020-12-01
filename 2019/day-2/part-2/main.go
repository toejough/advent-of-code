package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func execIntCode(ic string) string {
	stringValues := strings.Split(ic, ",")
	intValues := []int{}

	for _, sv := range stringValues {
		iv, err := strconv.Atoi(sv)
		if err != nil {
			log.Fatalf("unable to convert string value (%v) to int: %w", sv, err)
		}

		intValues = append(intValues, iv)
	}

	commandIndex := 0
	for commandIndex < len(intValues) {
		command := intValues[commandIndex+0]
		if command == 99 {
			break
		}

		argIndex1 := intValues[commandIndex+1]
		argIndex2 := intValues[commandIndex+2]
		outputIndex := intValues[commandIndex+3]

		switch command {
		case 1:
			intValues[outputIndex] = intValues[argIndex1] + intValues[argIndex2]
		case 2:
			intValues[outputIndex] = intValues[argIndex1] * intValues[argIndex2]
		default:
			log.Fatalf(
				"unable to run intcode operation for unknown operation type (%v)",
				command,
			)
		}

		commandIndex += 4
	}

	stringValues = []string{}

	for _, iv := range intValues {
		stringValues = append(stringValues, strconv.Itoa(iv))
	}

	return strings.Join(stringValues, ",")
}

func loadInputData() string {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read input.txt: %w", err)
	}

	inputString := strings.TrimSpace(string(inputBytes))

	return inputString
}

func updateString(ic string, noun, verb int) string {
	stringValues := strings.Split(ic, ",")
	intValues := []int{}

	for _, sv := range stringValues {
		iv, err := strconv.Atoi(sv)
		if err != nil {
			log.Fatalf("unable to convert string value (%v) to int: %w", sv, err)
		}

		intValues = append(intValues, iv)
	}

	intValues[1] = noun
	intValues[2] = verb

	stringValues = []string{}

	for _, iv := range intValues {
		stringValues = append(stringValues, strconv.Itoa(iv))
	}

	return strings.Join(stringValues, ",")
}

func firstInt(ic string) int {
	stringValues := strings.Split(ic, ",")
	intValues := []int{}

	for _, sv := range stringValues {
		iv, err := strconv.Atoi(sv)
		if err != nil {
			log.Fatalf("unable to convert string value (%v) to int: %w", sv, err)
		}

		intValues = append(intValues, iv)
	}

	return intValues[0]
}

func do() int {
	var noun, verb int
OuterLoop:
	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			inputString := loadInputData()
			updatedString := updateString(inputString, noun, verb)
			outputString := execIntCode(updatedString)
			intCodeResult := firstInt(outputString)
			if intCodeResult == 19690720 {
				break OuterLoop
			}
		}
	}

	return 100*noun + verb
}

func main() {
	fmt.Printf("The correct IntCode input is: %v\n", do())
}
