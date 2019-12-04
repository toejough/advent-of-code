package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func loadInputData() []int {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read input.txt: %w", err)
	}

	inputString := string(inputBytes)
	inputLines := strings.Split(inputString, "\n")
	nonEmptyInputLines := []string{}

	for _, s := range inputLines {
		if len(s) != 0 {
			nonEmptyInputLines = append(nonEmptyInputLines, s)
		}
	}

	inputNumbers := []int{}

	for _, s := range nonEmptyInputLines {
		iv, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("unable to convert \"%v\" to an integer: %w", s, err)
		}

		inputNumbers = append(inputNumbers, iv)
	}

	return inputNumbers
}

func calculateFuelForMass(v int) int {
	fv := float64(v)
	fv /= 3
	v = int(math.Floor(fv))
	v -= 2

	if v < 0 {
		return 0
	}

	return v + calculateFuelForMass(v)
}

func calculateFuelNeeds(inputNumbers []int) []int {
	results := []int{}

	for _, v := range inputNumbers {
		fuelNeed := calculateFuelForMass(v)
		results = append(results, fuelNeed)
	}

	return results
}

func sumInts(ints []int) int {
	sum := 0

	for _, v := range ints {
		sum += v
	}

	return sum
}

func do() int {
	inputNumbers := loadInputData()
	fuelNeeds := calculateFuelNeeds(inputNumbers)
	sum := sumInts(fuelNeeds)

	return sum
}

func main() {
	fmt.Printf("Fuel needs for all modules & their fuel, recursively: %v\n", do())
}
