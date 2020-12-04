package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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

func countTrees(treeMap []string, s slope) (count int) {
	x := 0
	y := 0
	trees := 0

	for y < len(treeMap) {
		if treeMap[y][x] == '#' {
			trees++
		}

		x += s.x
		x %= len(treeMap[0])
		y += s.y
	}

	return trees
}

type slope struct {
	x int
	y int
}

func solve(input string) (output string) {
	lines := strings.Split(input, "\n")
	stripped := stripAll(lines)
	nonEmpty := skipEmpty(stripped)
	cumulative := 1

	slopes := []slope{
		{x: 1, y: 1}, //nolint:gomnd
		{x: 3, y: 1}, //nolint:gomnd
		{x: 5, y: 1}, //nolint:gomnd
		{x: 7, y: 1}, //nolint:gomnd
		{x: 1, y: 2}, //nolint:gomnd
	}
	for _, s := range slopes {
		cumulative *= countTrees(nonEmpty, s)
	}

	return strconv.Itoa(cumulative)
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	output := solve(string(input))
	fmt.Println(output)
}
