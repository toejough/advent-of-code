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

func (li *LineIterator) Next() (string, bool) {
	if len(li.Lines) <= li.Index {
		return "", false
	}

	line := li.Lines[li.Index]
	log.Print(line)
	li.Index++

	return line, true
}

type Fish struct {
	Timer int
}

func createFish(number int) Fish {
	return Fish{Timer: number}
}

func createFishSlice(numbers []int) []Fish {
	fs := []Fish{}
	for _, n := range numbers {
		fs = append(fs, createFish(n))
	}

	return fs
}

func main() {
	lineIterator := createLineIterator(os.Args[1])

	line, ok := lineIterator.Next()
	if !ok {
		log.Fatal("no input line")
	}

	numberStrings := strings.Split(line, ",")
	numbers := []int{}

	for _, ns := range numberStrings {
		numbers = append(numbers, atoi(ns))
	}

	fish := createFishSlice(numbers)

	const newFishTimer = 8

	for i := 0; i < 80; i++ {
		newFish := []Fish{}

		for j := range fish {
			f := &fish[j]
			if f.Timer == 0 {
				f.Timer = 6

				newFish = append(newFish, createFish(newFishTimer))
			} else {
				f.Timer--
			}
		}

		fish = append(fish, newFish...)
		log.Printf("Day %d, Count: %d", i, len(fish))
	}

	log.Printf("Final Count: %d", len(fish))
}
