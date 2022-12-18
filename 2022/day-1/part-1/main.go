package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read input
	day := os.Args[1]
	part := os.Args[2]

	filename := fmt.Sprintf("%s-input-puzzle.txt", day)
	text := mustReadFileText(filename)

	// solve
	var (
		answer string
		err    error
	)

	switch day {
	case "day1":
		switch part {
		case "part1":
			answer, err = solveDay1Part1(text)
			if err != nil {
				panic(err)
			}
		case "part2":
			answer, err = solveDay1Part2(text)
			if err != nil {
				panic(err)
			}
		default:
			log.Fatalf("No solver for %s %s", day, part)
		}
	default:
		log.Fatalf("No solvers for %s ", day)
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

func splitNoEmpty(s string, sep string) []string {
	list := strings.Split(s, sep)
	noEmpty := []string{}

	for _, text := range list {
		if len(text) == 0 {
			continue
		}

		noEmpty = append(noEmpty, text)
	}

	return noEmpty
}

func solveDay1Part1(text string) (string, error) {
	// split into lists of calories
	hunks := splitNoEmpty(text, "\n\n")
	stringLists := splitHunks(hunks)

	lists, err := convListsOfStringsToListsOfInts(stringLists)
	if err != nil {
		return "", err
	}

	// sum the lists
	sums := sumLists(lists)

	// max
	max, err := max(sums)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", max), nil
}

func solveDay1Part2(text string) (string, error) {
	// split into lists of calories
	hunks := splitNoEmpty(text, "\n\n")
	stringLists := splitHunks(hunks)

	lists, err := convListsOfStringsToListsOfInts(stringLists)
	if err != nil {
		return "", err
	}

	// sum the lists
	sums := sumLists(lists)

	// sort
	sort.Ints(sums)

	// top 3
	top3 := sums[len(sums)-3:]

	// sum
	total := sum(top3)

	return fmt.Sprintf("%d", total), nil
}

var ErrNoMaxPossible = fmt.Errorf("no max possible: input list was empty")

func max(list []int) (int, error) {
	if len(list) == 0 {
		return 0, ErrNoMaxPossible
	}

	max := list[0]
	for _, value := range list[1:] {
		if value > max {
			max = value
		}
	}

	return max, nil
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

func splitHunks(hunks []string) [][]string {
	lists := [][]string{}

	for _, hunk := range hunks {
		list := splitNoEmpty(hunk, "\n")
		lists = append(lists, list)
	}

	return lists
}

func convListsOfStringsToListsOfInts(stringLists [][]string) ([][]int, error) {
	lists := [][]int{}

	for _, stringList := range stringLists {
		list, err := convStringsToInts(stringList)
		if err != nil {
			return nil, fmt.Errorf("while converting string list '%v' to int: %w", list, err)
		}

		lists = append(lists, list)
	}

	return lists, nil
}

func convStringsToInts(stringList []string) ([]int, error) {
	list := []int{}

	for _, s := range stringList {
		value, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("while converting string '%s' to int: %w", s, err)
		}

		list = append(list, value)
	}

	return list, nil
}
