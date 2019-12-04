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

func main() {
	fmt.Printf("Fuel needs for all modules: %v\n", do())
}

func do() int {
	inputNumbers := loadInputData()
	results := []int{}

	for _, v := range inputNumbers {
		fv := float64(v)
		fv /= 3
		v = int(math.Floor(fv))
		v -= 2
		results = append(results, v)
	}

	sum := 0

	for _, v := range results {
		sum += v
	}

	return sum
}
