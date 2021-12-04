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
	li.Index++

	return line, true
}

func main() {
	lineIterator := createLineIterator(os.Args[1])

	next, ok := lineIterator.Next()
	lines := []string{}

	for ok {
		lines = append(lines, next)
		next, ok = lineIterator.Next()
	}

    c02Lines := make([]string, len(lines))
    copy(c02Lines, lines)

    index := 0

	for len(lines) > 1 {
		count := 0

		for _, l := range lines {
			if l[index] == '1' {
				count++
			}
		}
        log.Printf("count of 1's: %d of %d", count, len(lines))

		mostFrequent := '0'

		if float64(count) >= float64(len(lines))/2 {
			mostFrequent = '1'
		}
        log.Printf("most frequent: %c", mostFrequent)

		oldLines := lines

		lines = []string{}

		for _, l := range oldLines {
			if ([]rune(l))[index] == mostFrequent {
				lines = append(lines, l)
                log.Println(l)
			}
		}
        index++
	}


	o2Rating, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		panic(err)
	}

    lines = c02Lines

    index = 0

	for len(lines) > 1 {
		count := 0

		for _, l := range lines {
			if l[index] == '1' {
				count++
			}
		}
        log.Printf("count of 1's: %d of %d", count, len(lines))

		leastFrequent := '1'

		if float64(count) >= float64(len(lines))/2 {
			leastFrequent = '0'
		}
        log.Printf("least frequent: %c", leastFrequent)

		oldLines := lines

		lines = []string{}

		for _, l := range oldLines {
			if ([]rune(l))[index] == leastFrequent {
				lines = append(lines, l)
                log.Println(l)
			}
		}
        index++
	}

	c02Rating, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		panic(err)
	}

	log.Printf("O2Rating: %d", o2Rating)
	log.Printf("C02Rating: %d", c02Rating)
	log.Printf("Life support: %d", o2Rating*c02Rating)
}
