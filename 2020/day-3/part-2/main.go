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

func countTrees(treeMap []string, slopeX, slopeY int) (count int) {
	x := 0
	y := 0
	trees := 0
	for y < len(treeMap) {
		// log.Printf("line: '%v', x: '%v', y: '%v'\n", nonEmpty[x], x, y)
		if treeMap[y][x] == '#' {
			trees++
		}
		x += slopeX
		x %= len(treeMap[0])
		y += slopeY
	}
	return trees
}

func solve(input string) (output string, err error) {
	lines := strings.Split(input, "\n")
	stripped := stripAll(lines)
	nonEmpty := skipEmpty(stripped)
	cumulative := 1

	cumulative *= countTrees(nonEmpty, 1, 1)
	cumulative *= countTrees(nonEmpty, 3, 1)
	cumulative *= countTrees(nonEmpty, 5, 1)
	cumulative *= countTrees(nonEmpty, 7, 1)
	cumulative *= countTrees(nonEmpty, 1, 2)

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
