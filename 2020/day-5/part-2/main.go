package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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
	if !d.byr.present || !isValidField(field{name: "byr", value: d.byr.value}) {
		return false
	}

	if !d.iyr.present || !isValidField(field{name: "iyr", value: d.iyr.value}) {
		return false
	}

	if !d.eyr.present || !isValidField(field{name: "eyr", value: d.eyr.value}) {
		return false
	}

	if !d.hgt.present || !isValidField(field{name: "hgt", value: d.hgt.value}) {
		return false
	}

	if !d.hcl.present || !isValidField(field{name: "hcl", value: d.hcl.value}) {
		return false
	}

	if !d.ecl.present || !isValidField(field{name: "ecl", value: d.ecl.value}) {
		return false
	}

	if !d.pid.present || !isValidField(field{name: "pid", value: d.pid.value}) {
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

type yearArgs struct {
	value string
	min   int
	max   int
}

func isValidYear(args yearArgs) (valid bool) {
	const yearLen = 4

	y, min, max := args.value, args.min, args.max

	if len(y) != yearLen {
		return false
	}

	asInt, err := strconv.Atoi(y)
	if err != nil {
		return false
	}

	if asInt < min {
		return false
	}

	if asInt > max {
		return false
	}

	return true
}

type field struct {
	name  string
	value string
}

var hgtRegex = regexp.MustCompile(`^(?P<number>\d+)(?P<unit>cm|in)$`) //nolint:gochecknoglobals
var hclRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)                   //nolint:gochecknoglobals
var pidRegex = regexp.MustCompile(`^[0-9]{9}$`)                       //nolint:gochecknoglobals

func isValidField(f field) (valid bool) {
	name, value := f.name, f.value

	const (
		minBirthYear      = 1920
		maxBirthYear      = 2002
		minIssuedYear     = 2010
		maxIssuedYear     = 2020
		minExpirationYear = 2020
		maxExpirationYear = 2030
	)

	switch name {
	case "byr":
		return isValidYear(yearArgs{value: value, min: minBirthYear, max: maxBirthYear})
	case "iyr":
		return isValidYear(yearArgs{value: value, min: minIssuedYear, max: maxIssuedYear})
	case "eyr":
		return isValidYear(yearArgs{value: value, min: minExpirationYear, max: maxExpirationYear})
	case "hgt":
		indices := hgtRegex.FindStringSubmatchIndex(value)
		if len(indices) == 0 {
			return false
		}

		number, err := strconv.Atoi(string(hgtRegex.ExpandString([]byte{}, "$number", value, indices)))
		if err != nil {
			panic(errors.Wrap(err, "Regex number match couldn't be converted to a number"))
		}

		unit := string(hgtRegex.ExpandString([]byte{}, "$unit", value, indices))
		switch unit {
		case "cm":
			if 150 <= number && number <= 193 {
				return true
			}
		case "in":
			if 59 <= number && number <= 76 {
				return true
			}
		}

		return false

	case "hcl":
		return hclRegex.MatchString(value)
	case "ecl":
		switch value {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		}

		return false
	case "pid":
		return pidRegex.MatchString(value)
	case "cid":
	}

	return false
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
