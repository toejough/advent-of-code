package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read input
	filename := os.Args[1]
	text := mustReadFileText(filename)
	lines := strings.Split(text, "\n")

	// solve
	answer, err := solve(lines)
	if err != nil {
		panic(err)
	}

	// final
	log.Printf("Answer: %s", answer)
}

func mustReadFileText(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	text := string(bytes)

	return text
}

func solve(lines []string) (string, error) {
	// split into lists of calories
	stringLists := splitListByBlankLines(lines)

	lists, err := convListsOfStringsToListsOfInts(stringLists)
	if err != nil {
		return "", err
	}

	// sum the lists
	sums := sumLists(lists)

	// max
	max := max(sums)

	return fmt.Sprintf("%d", max), nil
}

func max(list []int) int {
	max := list[0]
	for _, value := range list[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

func sumLists(lists [][]int) []int {
	sums := []int{}

	for _, list := range lists {
		sums = append(sums, sum(list))
	}

	return sums
}

func sum(list []int) int {
	sum := 0
	for _, value := range list {
		sum += value
	}

	return sum
}

func splitListByBlankLines(lines []string) [][]string {
	lists := [][]string{}
	list := []string{}

	for _, line := range lines {
		if len(line) == 0 {
			lists = append(lists, list)
			list = []string{}
		} else {
			list = append(list, line)
		}
	}

	// final list
	lists = append(lists, list)

	return lists
}

func convListsOfStringsToListsOfInts(stringLists [][]string) ([][]int, error) {
	lists := [][]int{}

	for _, stringList := range stringLists {
		list := []int{}

		for _, s := range stringList {
			value, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("while converting string '%s' to int: %w", s, err)
			}

			list = append(list, value)
		}

		lists = append(lists, list)
	}

	return lists, nil
}
