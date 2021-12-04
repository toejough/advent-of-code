package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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


func main() {
	lineIterator := createLineIterator(os.Args[1])

	next, ok := lineIterator.Next()
    counts := make([]int, len(next))
    numLines := 0

	for ok {
        numLines++

        for i, v := range(next) {
            if v == '1' {
                counts[i]++
            }
        }

		next, ok = lineIterator.Next()
	}

    gammaStr := ""
    epsilonStr := ""
    for _, v:= range(counts) {
        if v > numLines / 2 {
            gammaStr += "1"
            epsilonStr += "0"
        } else {
            gammaStr += "0"
            epsilonStr += "1"
        }
    }

    gamma, err := strconv.ParseInt(gammaStr, 2, 64)
    if err != nil {
        panic(err)
    }

    epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
    if err != nil {
        panic(err)
    }

	log.Printf("gammaStr: %s", gammaStr)
	log.Printf("gamma: %d", gamma)
	log.Printf("epsilonStr: %s", epsilonStr)
	log.Printf("epsilon: %d", epsilon)
	log.Printf("power: %d", gamma*epsilon)
}
