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

func solve(input string) (output string, err error) {
	lines := strings.Split(input, "\n")
	stripped := stripAll(lines)
	nonEmpty := skipEmpty(stripped)
	cumulative := 1

	{
		x := 0
		y := 0
		trees := 0
		for y < len(nonEmpty) {
			// log.Printf("line: '%v', x: '%v', y: '%v'\n", nonEmpty[x], x, y)
			if nonEmpty[y][x] == '#' {
				trees++
			}
			x += 1
			x %= len(nonEmpty[0])
			y += 1
		}
		cumulative *= trees
	}
	{
		x := 0
		y := 0
		trees := 0
		for y < len(nonEmpty) {
			// log.Printf("line: '%v', x: '%v', y: '%v'\n", nonEmpty[x], x, y)
			if nonEmpty[y][x] == '#' {
				trees++
			}
			x += 3
			x %= len(nonEmpty[0])
			y += 1
		}
		cumulative *= trees
	}
	{
		x := 0
		y := 0
		trees := 0
		for y < len(nonEmpty) {
			// log.Printf("line: '%v', x: '%v', y: '%v'\n", nonEmpty[x], x, y)
			if nonEmpty[y][x] == '#' {
				trees++
			}
			x += 5
			x %= len(nonEmpty[0])
			y += 1
		}
		cumulative *= trees
	}
	{
		x := 0
		y := 0
		trees := 0
		for y < len(nonEmpty) {
			// log.Printf("line: '%v', x: '%v', y: '%v'\n", nonEmpty[x], x, y)
			if nonEmpty[y][x] == '#' {
				trees++
			}
			x += 7
			x %= len(nonEmpty[0])
			y += 1
		}
		cumulative *= trees
	}
	{
		x := 0
		y := 0
		trees := 0
		for y < len(nonEmpty) {
			// log.Printf("line: '%v', x: '%v', y: '%v'\n", nonEmpty[x], x, y)
			if nonEmpty[y][x] == '#' {
				trees++
			}
			x += 1
			x %= len(nonEmpty[0])
			y += 2
		}
		cumulative *= trees
	}
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
