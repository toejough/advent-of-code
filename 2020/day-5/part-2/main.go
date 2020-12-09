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

func skipEmpty(items []string) (nonEmpty []string) {
	for _, item := range items {
		if len(item) > 0 {
			nonEmpty = append(nonEmpty, item)
		}
	}

	return
}

func generateAll() (all []decoded) {
	for row := 0; row <= 127; row++ {
		for column := 0; column <= 7; column++ {
			all = append(all, decoded{seat: column, row: row})
		}
	}

	return all
}

func singleNotIn(i int, s []int) (result bool) {
	for _, item := range s {
		if i == item {
			return false
		}
	}

	return true
}

func difference(base []int, remove []int) (remaining []int) {
	for _, b := range base {
		if singleNotIn(b, remove) {
			remaining = append(remaining, b)
		}
	}

	return remaining
}

func add(items []int, value int) (added []int) {
	for _, i := range items {
		added = append(added, i+value)
	}

	return added
}

func notIn(items []int, container []int) (missing []int) {
	for _, i := range items {
		if singleNotIn(i, container) {
			missing = append(missing, i)
		}
	}

	return missing
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
	all := idsFrom(generateAll())
	missing := difference(all, ids)
	log.Printf("Missing ID's: %v\n", missing)
	options := add(notIn(add(missing, 1), missing), -1)
	options = add(notIn(add(options, -1), missing), 1)

	if len(options) != 1 { //nolint:gomd
		return "", errors.Errorf("expected exactly 1 option, but got %v", options)
	}

	return strconv.Itoa(options[0]), nil
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
