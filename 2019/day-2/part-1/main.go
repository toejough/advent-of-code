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

func updateString(ic string) string {
	stringValues := strings.Split(ic, ",")
	intValues := []int{}

	for _, sv := range stringValues {
		iv, err := strconv.Atoi(sv)
		if err != nil {
			log.Fatalf("unable to convert string value (%v) to int: %w", sv, err)
		}

		intValues = append(intValues, iv)
	}

	intValues[1] = 12
	intValues[2] = 2

	stringValues = []string{}

	for _, iv := range intValues {
		stringValues = append(stringValues, strconv.Itoa(iv))
	}

	return strings.Join(stringValues, ",")
}

func firstString(ic string) string {
	stringValues := strings.Split(ic, ",")
	intValues := []int{}

	for _, sv := range stringValues {
		iv, err := strconv.Atoi(sv)
		if err != nil {
			log.Fatalf("unable to convert string value (%v) to int: %w", sv, err)
		}

		intValues = append(intValues, iv)
	}

	return strconv.Itoa(intValues[0])
}

func do() string {
	inputString := loadInputData()
	updatedString := updateString(inputString)
	outputString := execIntCode(updatedString)

	return firstString(outputString)
}

func main() {
	fmt.Printf("The correct IntCode start is: %v\n", do())
}
