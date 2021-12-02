package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
    log.Printf("File: %s", file.Name())
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
    line := scanner.Text()
    log.Printf("line: %s", line)

	lastDepth, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}


	numIncreases := 0

	for scanner.Scan() {
        line = scanner.Text()
        log.Printf("line: %s", line)
		depth, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if depth > lastDepth {
			numIncreases += 1
            log.Println("increased!")
		}
		lastDepth = depth
	}
	log.Printf("NumIncreases: %d", numIncreases)
}
