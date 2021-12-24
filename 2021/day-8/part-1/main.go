package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	log.Printf("File: %s", file.Name())

	return file
}

func atoi(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return number
}

type LineIterator struct {
	Lines []string
	Index int
}

func createLineIterator(path string) *LineIterator {
	file := openFile(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	li := &LineIterator{}

	for scanner.Scan() {
		li.Lines = append(li.Lines, scanner.Text())
	}

	return li
}

func readLines(path string) []string {
	file := openFile(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
func (li *LineIterator) Next() (string, bool) {
	if len(li.Lines) <= li.Index {
		return "", false
	}

	line := li.Lines[li.Index]
	log.Print(line)
	li.Index++

	return line, true
}

func createFishMap(numbers []int) map[int]int {
	fish := map[int]int{}
	for _, n := range numbers {
		fish[n]++
	}

	return fish
}

func sum(fish map[int]int) int {
	total := 0
	for _, v := range fish {
		total += v
	}

	return total
}

/*
Decoding Numbers

0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg

With the above as a guide:
number: segments: num segments
0: abcefg: 5
1: cf: 2
2: acdeg: 5
3: acdfg: 5
4: bcdf: 4
5: abdfg: 5
6: abdefg: 6
7: acf: 3
8: abcdefg: 7
9: abcdfg: 6

Unique numbers:
1: cf: 2
4: bcdf: 4
7: acf: 3
8: abcdefg: 7
*/

type display struct {
	patterns []string
	output   []string
}

func parseDisplay(line string) display {
	parts := strings.Split(line, "|")
	return display{
		patterns: strings.Fields(parts[0]),
		output:   strings.Fields(parts[1]),
	}
}

func parseDisplays(lines []string) []display {
	displays := make([]display, len(lines))
	for i, l := range lines {
		displays[i] = parseDisplay(l)
	}
	return displays
}

func extractOutputs(displays []display) [][]string {
	outputs := make([][]string, len(displays))
	for i, d := range displays {
		outputs[i] = d.output
	}
	return outputs
}

func flattenOutputs(outputs [][]string) []string {
	flattened := []string{}
	for _, o := range outputs {
		for _, each := range o {
			flattened = append(flattened, each)
		}
	}
	return flattened
}

func filterToUnique(outputs []string) []string {
	unique := []string{}
	for _, o := range outputs {
		switch len(o) {
		case 2, 4, 3, 7:
			unique = append(unique, o)
		}
	}
	return unique
}

func main() {
	flag.Parse()
	args := flag.Args()
	fileToRead := args[0]
	lines := readLines(fileToRead)
	displays := parseDisplays(lines)
	outputs := extractOutputs(displays)
	numbers := flattenOutputs(outputs)
	uniqueNumbers := filterToUnique(numbers)
	count := len(uniqueNumbers)

	log.Printf("Final Count: %d", count)
}
