package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func stripAll(items []string) (allStripped []string) {
	for _, item := range items {
		stripped := strings.TrimSpace(item)
		allStripped = append(allStripped, stripped)
	}
	return
}

func skipEmpty(items []string) (nonEmpty []string) {
	for _, item := range items {
		if len(item) > 0 {
			nonEmpty = append(nonEmpty, item)
		}
	}
	return
}

func allToInts(items []string) (allInts []int, err error) {
	for _, item := range items {
		thisInt, err := strconv.Atoi(item)
		if err != nil {
			return nil, errors.Wrapf(err, "Converting '%v' from string to int", item)
		}
		allInts = append(allInts, thisInt)
	}
	return allInts, nil
}

func solve(input string) (output string, err error) {
	lines := strings.Split(input, "\n")
	stripped := stripAll(lines)
	nonEmpty := skipEmpty(stripped)
	intItems, err := allToInts(nonEmpty)
	if err != nil {
		return "", errors.Wrap(err, "Converting input []string to []int")
	}
	iItems := intItems[:len(intItems)-1]
	for i, iItem := range iItems {
		jItems := intItems[i+1:]
		for _, jItem := range jItems {
			if iItem+jItem == 2020 {
				output = strconv.Itoa(iItem * jItem)
				return output, nil
			}
		}
	}
	err = errors.New("Never found two values that summed to 2020")
	return "", err
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	output, err := solve(string(input))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(output)
}
