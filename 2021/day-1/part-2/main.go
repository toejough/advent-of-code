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

func main() {
	scanner := bufio.NewScanner(openFile(os.Args[1]))
	ok, first := readInt(scanner)
	if !ok {
		log.Fatalln("No first line found to read...")
	}

	ok, second := readInt(scanner)
	if !ok {
		log.Fatalln("No second line found to read...")
	}

	ok, third := readInt(scanner)
	if !ok {
		log.Fatalln("No third line found to read...")
	}

	lastDepthSum := first + second + third
	log.Printf("DepthSum: %d", lastDepthSum)

	ok, depth := readInt(scanner)
	numIncreases := 0

	for ok {
		first, second, third = second, third, depth
		depthSum := first + second + third
		log.Printf("DepthSum: %d", depthSum)
		if depthSum > lastDepthSum {
			numIncreases += 1
			log.Println("increased!")
		}
		lastDepthSum = depthSum
		ok, depth = readInt(scanner)
	}

	log.Printf("NumIncreases: %d", numIncreases)
}
