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

func solve(input string) (output string, err error) {
	lines := strings.Split(input, "\n")
	stripped := stripAll(lines)
	nonEmpty := skipEmpty(stripped)
	cumulative := 1

	cumulative *= countTrees(nonEmpty, slope{x: 1, y: 1})
	cumulative *= countTrees(nonEmpty, slope{x: 3, y: 1})
	cumulative *= countTrees(nonEmpty, slope{x: 5, y: 1})
	cumulative *= countTrees(nonEmpty, slope{x: 7, y: 1})
	cumulative *= countTrees(nonEmpty, slope{x: 1, y: 2})

	return strconv.Itoa(cumulative), nil
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
