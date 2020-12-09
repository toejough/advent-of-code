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

type decoded struct {
	row  int
	seat int
}

func (d decoded) id() int {
	return d.row*8 + d.seat
}

var seatRegex = regexp.MustCompile(`^(?P<row>(?:F|B){7})(?P<seat>(?:R|L){3})$`) // nolint:gochecknoglobals
func decodeRow(spec string) (row int) {
	min, max := 0, 127

	for _, v := range spec {
		if v == 'F' {
			max = (max-min+1)/2 + min - 1 //nolint:gomnd
			row = min
		} else {
			min = (max-min+1)/2 + min //nolint:gomnd
			row = max
		}
	}

	return row
}

func decodeSeat(spec string) (seat int) {
	min, max := 0, 7

	for _, v := range spec {
		if v == 'L' {
			max = (max-min+1)/2 + min - 1 //nolint:gomnd
			seat = min
		} else {
			min = (max-min+1)/2 + min //nolint:gomnd
			seat = max
		}
	}

	return seat
}

func decode(s string) (d decoded, err error) {
	log.Printf("string to decode: %v\n", s)

	indices := seatRegex.FindStringSubmatchIndex(s)
	if len(indices) == 0 {
		return d, errors.Errorf("unable to decode `%v`\n", s)
	}

	rowSpec := string(seatRegex.ExpandString([]byte{}, "$row", s, indices))

	log.Printf("rowSpec: %v\n", rowSpec)

	d.row = decodeRow(rowSpec)

	log.Printf("row: %v\n", d.row)

	seatSpec := string(seatRegex.ExpandString([]byte{}, "$seat", s, indices))

	log.Printf("seatSpec: %v\n", seatSpec)

	d.seat = decodeSeat(seatSpec)

	log.Printf("seat: %v\n", d.seat)

	return d, nil
}

func decodeAll(lines []string) (allDecoded []decoded, err error) {
	for _, l := range lines {
		oneDecoded, err := decode(l)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to decode `%v`", l)
		}

		allDecoded = append(allDecoded, oneDecoded)
	}

	return allDecoded, nil
}

func idsFrom(d []decoded) (ids []int) {
	for _, thisD := range d {
		ids = append(ids, thisD.id())
	}

	return ids
}

type maybeInt struct {
	value   int
	err     error
	isValue bool
}

func (m *maybeInt) setError(err error) {
	m.err = err
	m.isValue = false
}

func (m *maybeInt) setValue(v int) {
	m.value = v
	m.isValue = true
}

func maxInt(ints []int) (max maybeInt) {
	if len(ints) == 0 {
		max.setError(errors.New("no ints to get the max from"))
		return max
	}

	maxValue := ints[0]
	for _, i := range ints[1:] {
		if maxValue < i {
			maxValue = i
		}
	}

	max.setValue(maxValue)

	return max
}

func skipEmpty(items []string) (nonEmpty []string) {
	for _, item := range items {
		if len(item) > 0 {
			nonEmpty = append(nonEmpty, item)
		}
	}

	return
}

func solve(input string) (output string, err error) {
	lines := strings.Split(input, "\n")
	stripped := stripAll(lines)
	nonEmpty := skipEmpty(stripped)

	decoded, err := decodeAll(nonEmpty)
	if err != nil {
		return "", errors.Wrap(err, "unable to decode all")
	}

	ids := idsFrom(decoded)

	maxID := maxInt(ids)
	if !maxID.isValue {
		return "", errors.Wrap(maxID.err, "unable to find max int")
	}

	return strconv.Itoa(maxID.value), nil
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
