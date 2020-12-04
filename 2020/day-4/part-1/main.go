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

func splitSections(lines []string) (sections [][]string) {
	s := []string{}

	for _, l := range lines {
		if len(l) > 0 {
			s = append(s, l)
		} else {
			sections = append(sections, s)
			s = []string{}
		}
	}

	return sections
}

type optString struct {
	present bool
	value   string
}

func (o *optString) set(v string) {
	o.present = true
	o.value = v
}

type doc struct {
	byr optString // (Birth Year)
	iyr optString // (Issue Year)
	eyr optString // (Expiration Year)
	hgt optString // (Height)
	hcl optString // (Hair Color)
	ecl optString // (Eye Color)
	pid optString // (Passport ID)
	cid optString // (Country ID)
}

func toSingleLine(lines []string) (line string) {
	if len(lines) == 0 {
		return line
	}

	line = lines[0]
	lines = lines[1:]

	if len(lines) == 0 {
		return line
	}

	for _, l := range lines {
		line += " " + l
	}

	return line
}

type pair struct {
	name  string
	value string
}

func toPair(s string) (p pair, err error) {
	parts := strings.Split(s, ":")

	const PAIR = 2

	if len(parts) != PAIR {
		return p, errors.Errorf("expected 'name:value' format, but got '%v'\n", s)
	}

	name := parts[0]
	value := parts[1]

	return pair{name: name, value: value}, nil
}

func toPairs(parts []string) (pairs []pair, err error) {
	for _, p := range parts {
		aPair, err := toPair(p)
		if err != nil {
			return nil, errors.Wrap(err, "unable to build pair")
		}

		pairs = append(pairs, aPair)
	}

	return pairs, nil
}

func toDoc(pairs []pair) (d doc) {
	for _, p := range pairs {
		switch p.name {
		case "byr":
			d.byr.set(p.value)
		case "iyr":
			d.iyr.set(p.value)
		case "eyr":
			d.eyr.set(p.value)
		case "hgt":
			d.hgt.set(p.value)
		case "hcl":
			d.hcl.set(p.value)
		case "ecl":
			d.ecl.set(p.value)
		case "pid":
			d.pid.set(p.value)
		case "cid":
			d.cid.set(p.value)
		}
	}

	return d
}

func skipEmpty(items []string) (nonEmpty []string) {
	for _, item := range items {
		if len(item) > 0 {
			nonEmpty = append(nonEmpty, item)
		}
	}

	return
}

func toDocs(sections [][]string) (docs []doc, err error) {
	for _, s := range sections {
		sl := toSingleLine(s)
		stripped := strings.TrimSpace(sl)
		parts := strings.Split(stripped, " ")
		nonEmpty := skipEmpty(parts)

		pairs, err := toPairs(nonEmpty)
		if err != nil {
			return nil, errors.Wrap(err, "unable to build pairs")
		}

		d := toDoc(pairs)
		docs = append(docs, d)
	}

	return docs, nil
}

func isValid(d doc) (valid bool) {
	if !d.byr.present {
		return false
	}

	if !d.iyr.present {
		return false
	}

	if !d.eyr.present {
		return false
	}

	if !d.hgt.present {
		return false
	}

	if !d.hcl.present {
		return false
	}

	if !d.ecl.present {
		return false
	}

	if !d.pid.present {
		return false
	}

	return true
}

func countValid(docs []doc) (numValid int) {
	for _, d := range docs {
		if isValid(d) {
			numValid++
		}
	}

	return numValid
}

func solve(input string) (output string, err error) {
	lines := strings.Split(input, "\n")
	stripped := stripAll(lines)
	sections := splitSections(stripped)

	docs, err := toDocs(sections)
	if err != nil {
		return "", errors.Wrap(err, "unable to build docs")
	}

	numValid := countValid(docs)

	return strconv.Itoa(numValid), nil
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	output, err := solve(string(input))
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Println(output)
}
