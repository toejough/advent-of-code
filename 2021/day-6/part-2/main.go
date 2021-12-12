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

func main() {
	flag.Parse()
	lineIterator := createLineIterator(flag.Args()[0])

	line, ok := lineIterator.Next()
	if !ok {
		log.Fatal("no input line")
	}

	numberStrings := strings.Split(line, ",")
	numbers := []int{}

	for _, ns := range numberStrings {
		numbers = append(numbers, atoi(ns))
	}

	fish := createFishMap(numbers)
	log.Printf("fish: %v", fish)

	for day := 0; day < 256; day++ {
		newFish := map[int]int{}

		for timer, count := range fish {
			if timer == 0 {
				newFish[6] += count
				newFish[8] += count
			} else {
				newFish[timer-1] += count
			}
		}

		log.Printf("Day %d, Count: %d", day, sum(fish))
		fish = newFish
		log.Printf("fish: %v", fish)
	}

	log.Printf("Final Count: %d", sum(fish))
}
