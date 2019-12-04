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
	fmt.Println("Load input data")

	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("unable to read input.txt: %w", err)
	}

	log.Println(inputBytes)

	inputString := string(inputBytes)
	log.Println(inputString)

	inputLines := strings.Split(inputString, "\n")
	log.Println(inputLines)

	nonEmptyInputLines := []string{}

	for _, s := range inputLines {
		if len(s) != 0 {
			nonEmptyInputLines = append(nonEmptyInputLines, s)
		}
	}

	log.Println(nonEmptyInputLines)

	inputNumbers := []int{}

	for _, s := range nonEmptyInputLines {
		iv, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("unable to convert \"%v\" to an integer: %w", s, err)
		}

		inputNumbers = append(inputNumbers, iv)
	}

	log.Println(inputNumbers)

	return inputNumbers
}

func main() {
	inputNumbers := loadInputData()

	fmt.Println("For each input...")

	results := []int{}

	for _, v := range inputNumbers {
		fmt.Println("  Divide by 3")

		fv := float64(v)
		fv /= 3

		fmt.Println("  Round down")

		v = int(math.Floor(fv))

		fmt.Println("  Subtract 2")

		v -= 2

		log.Println(v)

		results = append(results, v)
	}

	log.Println(results)

	fmt.Println("Sum the individual results")

	sum := 0

	for _, v := range results {
		sum += v
	}

	log.Println(sum)
}
