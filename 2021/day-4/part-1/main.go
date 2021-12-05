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
	li.Index++

	return line, true
}

func parseNumbers(li *LineIterator) []int {
    line, ok := li.Next()
    if !ok {
        log.Fatal("no numbers line found")
    }
    numStrings := strings.Split(line, ",")
    nums := []int{}

    for _, s := range(numStrings) {
        nums = append(nums, atoi(s))
    }

    return nums
}

func parseBoards(li *LineIterator) []board {

}

func main() {
	lineIterator := createLineIterator(os.Args[1])

    numbers := parseNumbers(lineIterator)
    boards := parseBoards(lineIterator)

    for _, n :=range(numbers) {
        for _, b:= range(boards) {
            b = markNumber(n, b)
            won := checkBoard(b)
            if won {
                sum := sumUnmarked(b)
                log.Printf("Final Number: %d", n)
                log.Printf("Sum: %d", sum)
                log.Printf("Score: %d", sum * n)
                return
            }
        }
    }

	log.Fatal("Got to the end without any board winning")
}
