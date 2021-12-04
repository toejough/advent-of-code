package main

import (
	"bufio"
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

func atoi(line string) int {
	number, err := strconv.Atoi(line)
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

func (li *LineIterator) Next() (string, bool) {
	if len(li.Lines) <= li.Index {
		return "", false
	}

    line := li.Lines[li.Index]
    log.Printf(line)
    li.Index ++

    return line, true
}

func parseLine(line string) (string, int){
    parts := strings.Split(line, " ")
    return parts[0], atoi(parts[1])
}

func main() {
	lineIterator := createLineIterator(os.Args[1])

	vPos := 0
	hPos := 0
    aim := 0

	next, ok := lineIterator.Next()

	for ok {
		direction, magnitude := parseLine(next)
		switch direction {
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		case "forward":
			hPos += magnitude
            vPos += aim * magnitude
		}
		next, ok = lineIterator.Next()
	}

	log.Printf("vPos: %d", vPos)
	log.Printf("hPos: %d", hPos)
	log.Printf("aim: %d", aim)
	log.Printf("multiplied: %d", vPos*hPos)
}
