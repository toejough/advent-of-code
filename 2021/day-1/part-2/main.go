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

func readInt(scanner *bufio.Scanner) (bool, int) {
	isALine := scanner.Scan()
	if !isALine {
		return false, 0
	}

	line := scanner.Text()
	log.Printf("line: %s", line)

	number := atoi(line)

	return true, number
}

func atoi(line string) int {
	number, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return number
}

type LineIterator struct {
	LineCh chan string
}

func createLineIterator(path string) LineIterator {
	file := openFile(os.Args[1])

	scanner := bufio.NewScanner(file)
	lineCh := make(chan string)

	go func() {
		for scanner.Scan() {
			lineCh <- scanner.Text()
		}
		close(lineCh)
		file.Close()
	}()

	return LineIterator{LineCh: lineCh}
}

func (li LineIterator) NextAsInt() (int, bool) {
	line, ok := <-li.LineCh
    if !ok {
        return 0, ok
    }

	number := atoi(line)
	return number, ok
}

func getThreeLines(li LineIterator) (int, int, int) {
	first, ok := li.NextAsInt()
	if !ok {
		log.Fatalln("No first line found to read...")
	}

	second, ok := li.NextAsInt()
	if !ok {
		log.Fatalln("No second line found to read...")
	}

	third, ok := li.NextAsInt()
	if !ok {
		log.Fatalln("No third line found to read...")
	}

    return first, second, third
}


func main() {
	lineIterator := createLineIterator(os.Args[1])

    first, second, third := getThreeLines(lineIterator)

	lastDepthSum := first + second + third
	log.Printf("DepthSum: %d", lastDepthSum)

	depth, ok := lineIterator.NextAsInt()
	numIncreases := 0

	for ok {
		first, second, third = second, third, depth
		depthSum := first + second + third

		log.Printf("DepthSum: %d", depthSum)

		if depthSum > lastDepthSum {
			numIncreases++

			log.Println("increased!")
		}

		lastDepthSum = depthSum
		depth, ok = lineIterator.NextAsInt()
	}

	log.Printf("NumIncreases: %d", numIncreases)
}
